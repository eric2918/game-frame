package internal

import (
	"frame/cmd/game/game/player"
	"frame/cmd/game/game/static"
	"frame/pb"
	"reflect"
	"time"

	"github.com/eric2918/leaf/gate"
)

func handle(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handle(&pb.C2GS_Heartbeat{}, handleHeartbeat)
}

func handleSkills(args []interface{}) {
	agent := args[1].(gate.Agent)

	sendMsg := &pb.GS2C_GetSkills{
		Code: 200,
		Data: nil,
	}
	sendMsg.Data = static.Data.RoleSkill

	agent.WriteMsg(sendMsg)
}

func handleRoles(args []interface{}) {
	agent := args[1].(gate.Agent)

	sendMsg := &pb.GS2C_GetRoles{
		Code: 200,
		Data: nil,
	}

	sendMsg.Data = static.Data.RoleBasic

	agent.WriteMsg(sendMsg)
}

func handleHeartbeat(args []interface{}) {
	agent := args[1].(gate.Agent)

	timestamp := time.Now().Unix()

	sendMsg := &pb.GS2C_Heartbeat{
		Timestamp: timestamp,
	}

	// 记录最后心跳时间，定时或推出时更新数据库
	playerData := agent.UserData().(*player.Player)
	playerData.Player.LastHeartbeatTime = timestamp

	// 设置缓存
	agent.SetUserData(playerData)

	agent.WriteMsg(sendMsg)
}
