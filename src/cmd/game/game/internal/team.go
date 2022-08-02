package internal

import (
	"frame/cmd/game/game/player"
	"frame/pb"
	"frame/pkg/code"
	"frame/pkg/mongo"

	"github.com/eric2918/leaf/gate"
	"github.com/eric2918/leaf/log"
	"gopkg.in/mgo.v2/bson"
)

func init() {
	handle(&pb.C2GS_GetUserTeams{}, handleGetUserTeams)
	handle(&pb.C2GS_EditUserTeam{}, handleEditUserTeam)
}

func handleGetUserTeams(args []interface{}) {
	agent := args[1].(gate.Agent)

	playerData := agent.UserData().(*player.Player)

	sendMsg := &pb.GS2C_GetUserTeams{
		Code: 200,
		Data: nil,
	}

	sendMsg.Data = playerData.Player.Teams

	agent.WriteMsg(sendMsg)
}

func handleEditUserTeam(args []interface{}) {
	req := args[0].(*pb.C2GS_EditUserTeam)

	agent := args[1].(gate.Agent)
	playerData := agent.UserData().(*player.Player)

	sendMsg := &pb.GS2C_EditUserTeam{
		Code: 200,
		Data: nil,
	}

	teamId := req.TeamId
	//add := false
	if teamId == "" {
		sendMsg.Code = code.ParamBindError
		agent.WriteMsg(sendMsg)
		return
		//teamId = snowflake.GenStr()
		//add = true
	}

	if len(req.UserRoleIds) > 6 {
		req.UserRoleIds = req.UserRoleIds[:6]
	}

	// 判断队员是否合法（后期根据业务需求添加）

	team := pb.UserTeam{
		TeamId:      teamId,
		TeamName:    req.TeamName,
		UserRoleIds: req.UserRoleIds,
	}
	//if add {
	//	playerData.Player.Teams = append(playerData.Player.Teams, &team)
	//} else {
	for i, userTeam := range playerData.Player.Teams {
		if userTeam.TeamId == teamId {
			playerData.Player.Teams[i] = &team
		}
	}
	//}

	update := bson.M{"$set": bson.M{"teams": playerData.Player.Teams}}
	selector := bson.M{"playerId": playerData.Player.PlayerId}
	err := mongo.Collection(mongo.GAME_DB, mongo.PLAYERS_COLLECTION).Update(selector, update)
	if err != nil {
		log.Error("edit user group error : %v \n", err)
		sendMsg.Code = code.ServerInternalError
	}

	// 更新缓存
	agent.SetUserData(playerData)

	sendMsg.Data = playerData.Player.Teams

	agent.WriteMsg(sendMsg)
}
