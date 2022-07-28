package conf

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/eric2918/leaf/conf"
)

var Server struct {
	LogLevel string
	LogPath  string

	MachineID int64
	StartTime string

	Region     string
	ServerName string
	WSAddr     string
	CertFile   string
	KeyFile    string
	TCPAddr    string
	MaxConnNum int
	Language   string

	ListenAddr      string
	ConnAddrs       map[string]string
	PendingWriteNum int

	ConsolePort   int
	ConsolePrompt string
	ProfilePath   string

	MongodbAddr       string
	MongodbSessionNum int

	RedisAddr     string
	RedisPassword string
	RedisDb       int

	JwtSecret  string
	JwtTimeout int

	RoomModuleCount int
	LoginAddr       string

	MaxTeamsCounts int
}

func init() {
	argsLen := len(os.Args)
	if argsLen < 2 {
		log.Fatalf("os args of len(%v) less than 2", argsLen)
	}

	filePath := os.Args[1]

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("%v", err)
	}
	err = json.Unmarshal(data, &Server)
	if err != nil {
		log.Fatalf("%v", err)
	}
}

func Init() {
	conf.LogLevel = Server.LogLevel
	conf.LogPath = Server.LogPath
	conf.LogFlag = LogFlag
	conf.ConsolePort = Server.ConsolePort
	conf.ProfilePath = Server.ProfilePath
	conf.ServerName = Server.ServerName
	conf.ListenAddr = Server.ListenAddr
	conf.ConnAddrs = Server.ConnAddrs
	conf.PendingWriteNum = Server.PendingWriteNum
	conf.HeartBeatInterval = HeartBeatInterval
}
