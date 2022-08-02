package main

import (
	"encoding/json"
	"fmt"
	"frame/pb"
	"frame/pkg/mongo"
	"io/ioutil"
	"os"
	"strings"

	"github.com/eric2918/leaf/log"
)

func main() {
	// 初始化mongo
	mongo.Init()
	InitConfig()
}

func InitConfig() {
	var files []string
	var err error
	files, err = GetAllFile("../bin/static")
	if err != nil {
		log.Fatal("err: %s", err.Error())
		return
	}

	for _, file := range files {
		slice := strings.Split(file, "/")
		filename := slice[len(slice)-1]
		coll := strings.Split(filename, ".")[0]

		data, err := ioutil.ReadFile(file)
		if err != nil {
			log.Error("err: %s", err.Error())
			continue
		}

		if len(data) == 0 {
			fmt.Println(coll, "data is nil")
			continue
		}

		configs := conf{
			data:       data,
			collection: coll,
		}

		switch coll {
		case "RoleBasic":
			var docs []*pb.RoleBasic
			err = json.Unmarshal(data, &docs)
			if err != nil {
				log.Fatal("err:", err)
			}
			for _, doc := range docs {
				configs.docs = append(configs.docs, doc)
			}
		case "RoleAttrBasic":
			var docs []*pb.RoleAttrBasic
			err = json.Unmarshal(data, &docs)
			if err != nil {
				log.Fatal("err:", err)
			}
			for _, doc := range docs {
				configs.docs = append(configs.docs, doc)
			}
		case "RoleSkill":
			var docs []*pb.RoleSkill
			err = json.Unmarshal(data, &docs)
			if err != nil {
				log.Fatal("err:", err)
			}
			for _, doc := range docs {
				configs.docs = append(configs.docs, doc)
			}
		case "RoleSkin":
			var docs []*pb.RoleSkin
			err = json.Unmarshal(data, &docs)
			if err != nil {
				log.Fatal("err:", err)
			}
			for _, doc := range docs {
				configs.docs = append(configs.docs, doc)
			}
		case "PropBasic":
			var docs []*pb.PropBasic
			err = json.Unmarshal(data, &docs)
			if err != nil {
				log.Fatal("err:", err)
			}
			for _, doc := range docs {
				configs.docs = append(configs.docs, doc)
			}
		case "RoleLvExp":
			var docs []*pb.RoleLvExp
			err = json.Unmarshal(data, &docs)
			if err != nil {
				log.Fatal("err:", err)
			}
			for _, doc := range docs {
				configs.docs = append(configs.docs, doc)
			}
		case "Career":
			var docs []*pb.Career
			err = json.Unmarshal(data, &docs)
			if err != nil {
				log.Fatal("err:", err)
			}
			for _, doc := range docs {
				configs.docs = append(configs.docs, doc)
			}
		case "CareerAdvancement":
			var docs []*pb.CareerAdvancement
			err = json.Unmarshal(data, &docs)
			if err != nil {
				log.Fatal("err:", err)
			}
			for _, doc := range docs {
				configs.docs = append(configs.docs, doc)
			}
		case "StageMission":
			var docs []*pb.StageMission
			err = json.Unmarshal(data, &docs)
			if err != nil {
				log.Fatal("err:", err)
			}
			for _, doc := range docs {
				configs.docs = append(configs.docs, doc)
			}
		default:
			continue
		}
		configs.Init()
	}
}

type conf struct {
	data       []byte
	docs       []interface{}
	collection string
}

func (c *conf) Init() (err error) {
	if err = mongo.Collection(mongo.STATIC_DB, c.collection).DropCollection(); err != nil && err.Error() != "ns not found" {
		log.Error("drop %s collection error: %s", c.collection, err.Error())
	}

	if err = mongo.Collection(mongo.STATIC_DB, c.collection).Insert(c.docs...); err != nil {
		log.Error("init role config error: %s", err.Error())
		return err
	}
	return nil
}

// 忽略文件
var filter = map[string]bool{
	".DS_Store": true,
	".WeDrive":  true,
}

// 递归获取指定目录下的所有文件名
func GetAllFile(pathname string) (res []string, err error) {
	var fis []os.FileInfo
	fis, err = ioutil.ReadDir(pathname)
	if err != nil {
		log.Error("读取文件目录(%s)失败:%s \n", pathname, err.Error())
		return res, err
	}

	// 所有文件/文件夹
	for _, fi := range fis {
		fileName := fi.Name()
		if _, ok := filter[fileName]; ok {
			continue
		}

		fullName := pathname + "/" + fi.Name()
		// 是文件夹则递归进入获取;是文件，则压入数组
		if fi.IsDir() {
			var temp []string
			temp, err = GetAllFile(fullName)
			if err != nil {
				log.Error("读取文件目录(%s)失败:%s \n", fullName, err.Error())
				return res, err
			}
			res = append(res, temp...)
		} else {
			res = append(res, fullName)
		}
	}

	return res, nil
}
