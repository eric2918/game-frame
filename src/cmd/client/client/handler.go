package client

import (
	"fmt"
	"frame/msg"
	"frame/pb"
	"frame/pkg/code"

	"github.com/eric2918/leaf/log"
)

func init() {
	msg.Processor.SetHandler(&pb.AS2C_Login{}, handleLogin)
	msg.Processor.SetHandler(&pb.GS2C_CheckLogin{}, handleCheckLogin)
	msg.Processor.SetHandler(&pb.GS2C_CreatePlayer{}, handleCreatePlayer)
}

func handleLogin(args []interface{}) {
	resp := args[0].(*pb.AS2C_Login)
	fmt.Println("handleLogin", resp)
	if resp.Code != 200 {
		Close()
		log.Error("login is error: %v, please input login cmd", code.Text(resp.Code))
		return
	}

	userData.AccountId = resp.Data.AccountId
	Start(resp.Data.GameAddr)

	sendMsg := &pb.C2GS_CheckLogin{Token: resp.Data.Token}
	Client.WriteMsg(sendMsg)
}

func handleCheckLogin(args []interface{}) {
	resp := args[0].(*pb.GS2C_CheckLogin)
	if resp.Code != 200 {
		Close()
		log.Error("check login is error: %v, please input login cmd", code.Text(resp.Code))
		return
	}

	if resp.Data != nil && resp.Data.AccountId != "" {
		userData.AccountId = resp.Data.AccountId
		userData.Nickname = resp.Data.Nickname

		log.Release("%v(%v) login user success", userData.Nickname, userData.AccountId)
	} else {
		sendMsg := &pb.C2GS_CreatePlayer{Nickname: userData.Nickname}
		Client.WriteMsg(sendMsg)
	}
}

func handleCreatePlayer(args []interface{}) {
	resp := args[0].(*pb.GS2C_CreatePlayer)
	userData.PlayerId = resp.Data.PlayerId

	log.Release("%v(%v) login and create player success", userData.Nickname, userData.PlayerId)
}
