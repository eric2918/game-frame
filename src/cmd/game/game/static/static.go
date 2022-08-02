package static

import (
	"frame/pb"
	"frame/pkg/mongo"

	"github.com/eric2918/leaf/log"
)

type Config struct {
	RoleSkill         []*pb.RoleSkill
	RoleBasic         []*pb.RoleBasic
	RoleAttrBasic     []*pb.RoleAttrBasic
	Career            map[int64]*pb.Career
	RoleCareerIds     map[int64][]int64
	CareerAdvancement map[int64]*pb.CareerAdvancement
	LevelBasic        map[int64]*pb.RoleLvExp
	PropBasic         []*pb.PropBasic
	PropIndex         map[int64]int
	RoleSkins         map[int64][]int64
	StageMission      map[int64]*pb.StageMission
}

var Data Config

func Init() {
	if err := mongo.Collection(mongo.STATIC_DB, "RoleSkill").Find(nil).Sort("skillId").All(&Data.RoleSkill); err != nil {
		log.Error("get skills error: %#v \n", err.Error())
	}

	if err := mongo.Collection(mongo.STATIC_DB, "RoleBasic").Find(nil).Sort("roleId").All(&Data.RoleBasic); err != nil {
		log.Error("get roles error: %#v \n", err.Error())
	}
	Data.RoleCareerIds = make(map[int64][]int64)
	for _, basic := range Data.RoleBasic {
		Data.RoleCareerIds[basic.RoleId] = basic.CareerIds
	}

	if err := mongo.Collection(mongo.STATIC_DB, "RoleAttrBasic").Find(nil).Sort("id").All(&Data.RoleAttrBasic); err != nil {
		log.Error("get roles error: %#v \n", err.Error())
	}

	if err := mongo.Collection(mongo.STATIC_DB, "PropBasic").Find(nil).Sort("itemId").All(&Data.PropBasic); err != nil {
		log.Error("get roles error: %#v \n", err.Error())
	}
	Data.PropIndex = make(map[int64]int)
	for i, prop := range Data.PropBasic {
		Data.PropIndex[prop.ItemId] = i
	}

	var levelBasic []*pb.RoleLvExp
	if err := mongo.Collection(mongo.STATIC_DB, "RoleLvExp").Find(nil).Sort("level").All(&levelBasic); err != nil {
		log.Error("get roles error: %#v \n", err.Error())
	}
	Data.LevelBasic = make(map[int64]*pb.RoleLvExp)
	for _, basic := range levelBasic {
		Data.LevelBasic[basic.Level] = basic
	}

	var skins []*pb.RoleSkin
	if err := mongo.Collection(mongo.STATIC_DB, "RoleSkin").Find(nil).Sort("skinId").All(&skins); err != nil {
		log.Error("get skins error: %#v \n", err.Error())
	}
	Data.RoleSkins = make(map[int64][]int64)
	for _, skin := range skins {
		Data.RoleSkins[skin.RoleId] = append(Data.RoleSkins[skin.RoleId], skin.SkinId)
	}

	var Career []*pb.Career
	if err := mongo.Collection(mongo.STATIC_DB, "Career").Find(nil).Sort("careerID").All(&Career); err != nil {
		log.Error("get roles error: %#v \n", err.Error())
	}
	Data.Career = make(map[int64]*pb.Career)
	for _, career := range Career {
		Data.Career[career.CareerID] = career
	}

	var advancements []*pb.CareerAdvancement
	if err := mongo.Collection(mongo.STATIC_DB, "CareerAdvancement").Find(nil).Sort("advancementID").All(&advancements); err != nil {
		log.Error("get roles error: %#v \n", err.Error())
	}
	Data.CareerAdvancement = make(map[int64]*pb.CareerAdvancement)
	for _, advancement := range advancements {
		Data.CareerAdvancement[advancement.AdvancementID] = advancement
	}

	var stageMissions []*pb.StageMission
	if err := mongo.Collection(mongo.STATIC_DB, "StageMission").Find(nil).Sort("missionId").All(&stageMissions); err != nil {
		log.Error("get stage mission error: %#v \n", err.Error())
	}
	Data.StageMission = make(map[int64]*pb.StageMission)
	for _, stageMission := range stageMissions {
		Data.StageMission[stageMission.MissionID] = stageMission
	}
}
