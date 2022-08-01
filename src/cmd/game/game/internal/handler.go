package internal

import (
	"fmt"
	"frame/cmd/game/center"
	"frame/cmd/game/player"
	"frame/pb"
	"frame/pkg/code"
	"frame/pkg/mongo"
	"reflect"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/eric2918/leaf/cluster"
	"github.com/eric2918/leaf/gate"
	"github.com/eric2918/leaf/log"
)

func handle(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handle(&pb.C2GS_CheckLogin{}, handleCheckLogin)
	handle(&pb.C2GS_CreatePlayer{}, handleCreatePlayer)
}

func handleCheckLogin(args []interface{}) {
	req := args[0].(*pb.C2GS_CheckLogin)
	agent := args[1].(gate.Agent)

	sendMsg := &pb.GS2C_CheckLogin{
		Code: 200,
	}
	sendErrFunc := func(codeErr int64) {
		sendMsg.Code = codeErr
		log.Error(code.Text(codeErr))
		agent.WriteMsg(sendMsg)
	}

	res, err := cluster.CallN("account", "CheckToken", req.Token)
	fmt.Println("handleCheckLogin", res, err)
	if err != nil {
		sendErrFunc(code.InvalidUsernameOrPassword)
		return
	}
	accountId := res[0].(string)
	username := res[1].(string)

	for {
		ok, err := center.ChanRPC.Call1("AccountOnline", accountId, agent)
		if err != nil {
			sendErrFunc(code.InvalidUsernameOrPassword)
			return
		}

		if ok.(bool) {
			break
		} else {
			time.Sleep(time.Second)
		}
	}

	var playerInfo pb.Player
	err = mongo.Collection(mongo.GAME_DB, mongo.PLAYERS_COLLECTION).Find(bson.M{"accountId": accountId}).One(&playerInfo)
	if err == nil {
		sendMsg.Code = 200
		sendMsg.Data = &pb.GS2C_CheckLoginData{
			AccountId: playerInfo.AccountId,
			Nickname:  playerInfo.Nickname,
		}

		agent.SetUserData(player.New(&playerInfo))
		center.ChanRPC.Go("PlayerOnline", playerInfo.PlayerId, agent)
	} else if err == mgo.ErrNotFound {
		agent.SetUserData(map[string]interface{}{
			"AccountId": accountId,
			"Username":  username,
		})
	} else {
		sendErrFunc(code.ServerInternalError)
		return
	}
	agent.WriteMsg(sendMsg)
}

func handleCreatePlayer(args []interface{}) {
	req := args[0].(*pb.C2GS_CreatePlayer)
	agent := args[1].(gate.Agent)

	sendMsg := &pb.GS2C_CreatePlayer{}
	account, ok := agent.UserData().(map[string]interface{})
	accountId := account["AccountId"].(string)
	if !ok {
		sendMsg.Code = code.ServerInternalError
		agent.WriteMsg(sendMsg)
		return
	}

	if req.Nickname == "" {
		req.Nickname = account["Username"].(string)
	}

	p := player.New(&pb.Player{})
	p.Init(accountId, req.Nickname)

	if err := mongo.Collection(mongo.GAME_DB, mongo.PLAYERS_COLLECTION).Insert(&p.Player); err != nil {
		sendMsg.Code = code.ServerInternalError
		agent.WriteMsg(sendMsg)
		return
	} else {
		sendMsg.Code = 200
		sendMsg.Data = &pb.GS2C_CreatePlayerData{
			PlayerId: p.Player.PlayerId,
		}

		agent.SetUserData(player.New(p.Player))
		center.ChanRPC.Go("PlayerOnline", p.Player.PlayerId, agent)
	}

	agent.WriteMsg(sendMsg)

}
