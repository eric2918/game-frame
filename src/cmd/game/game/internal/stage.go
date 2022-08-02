package internal

import (
	"fmt"
	"frame/cmd/game/game/player"
	"frame/cmd/game/game/stage"
	"frame/cmd/game/game/static"
	"frame/pb"
	"frame/pkg/code"
	"frame/pkg/mongo"
	"frame/pkg/tools"

	"github.com/eric2918/leaf/gate"
	"github.com/spf13/cast"
	"gopkg.in/mgo.v2/bson"
)

func init() {
	handle(&pb.C2GS_GetUserStageMission{}, handleGetUserStageMission)
	handle(&pb.C2GS_UpdateUserStageMission{}, handleUpdateUserStageMission)
}

func handleGetUserStageMission(args []interface{}) {
	agent := args[1].(gate.Agent)

	sendMsg := &pb.GS2C_GetUserStageMission{
		Code: 200,
		Data: nil,
	}

	playerData := agent.UserData().(*player.Player)
	for _, v := range playerData.Player.StageMission {
		sendMsg.Data = append(sendMsg.Data, v)
	}

	agent.WriteMsg(sendMsg)
}

func handleUpdateUserStageMission(args []interface{}) {
	req := args[0].(*pb.C2GS_UpdateUserStageMission)
	agent := args[1].(gate.Agent)

	sendMsg := &pb.GS2C_UpdateUserStageMission{
		Code: 200,
		Data: nil,
	}
	missionId := req.GetMissionID()

	playerData := agent.UserData().(*player.Player)

	// 获取解锁关卡索引
	missionID := cast.ToString(req.MissionID)
	stageMission, ok := playerData.Player.StageMission[missionID]
	if !ok {
		fmt.Println("该关卡未解锁")
		sendMsg.Code = code.ServerInternalError
		agent.WriteMsg(sendMsg)
		return
	}
	if stageMission.StarLevel >= 3 {
		fmt.Println("该关卡已通关")
		sendMsg.Code = code.ServerInternalError
		agent.WriteMsg(sendMsg)
		return
	}

	// 测试默认升1级(后续添加升星条件)
	stageMission.StarLevel += 1

	// 解锁下个关卡
	nextMission := static.Data.StageMission[missionId].MissionNext
	if nextMission != "" && stageMission.StarLevel > 0 {
		s := stage.New(playerData.Player.StageMission, static.Data.StageMission)
		ids := tools.StringToSlice(nextMission)
		s.Unlock(ids)
		playerData.Player.StageMission = s.Mission
	}

	// 更新数据库
	selector := bson.M{"playerId": playerData.Player.PlayerId}
	update := bson.M{"$set": bson.M{"stageMission": playerData.Player.StageMission}}
	if err := mongo.Collection(mongo.GAME_DB, mongo.PLAYERS_COLLECTION).Update(selector, update); err != nil {
		fmt.Println("error:", err.Error())
		sendMsg.Code = code.ServerInternalError
		agent.WriteMsg(sendMsg)
		return
	}

	// 设置缓存
	agent.SetUserData(playerData)

	// 升星之后，解锁下关卡
	agent.WriteMsg(sendMsg)
}
