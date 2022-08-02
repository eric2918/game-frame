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

	//msg.Processor.SetHandler(&pb.GS2C_GetSkills{}, handleSkills)
	//msg.Processor.SetHandler(&pb.GS2C_GetRoles{}, handleRoles)
	msg.Processor.SetHandler(&pb.GS2C_GetUserTeams{}, handleGetUserTeams)
	msg.Processor.SetHandler(&pb.GS2C_GetUserRoles{}, handleGetUserRoles)
	msg.Processor.SetHandler(&pb.GS2C_EditUserTeam{}, handleEditUserTeam)
	msg.Processor.SetHandler(&pb.GS2C_Heartbeat{}, handleHeartbeat)

	msg.Processor.SetHandler(&pb.GS2C_UpdateUserProps{}, handleUpdateUserProps)
	msg.Processor.SetHandler(&pb.GS2C_GetUserProps{}, handleGetUserProps)
	msg.Processor.SetHandler(&pb.GS2C_RoleUpgrade{}, handleRoleUpgrade)
	msg.Processor.SetHandler(&pb.GS2C_RoleAdvanced{}, handleRoleAdvanced)
	msg.Processor.SetHandler(&pb.GS2C_GetUserStageMission{}, handleGetUserStageMission)
	msg.Processor.SetHandler(&pb.GS2C_UpdateUserStageMission{}, handleUpdateUserStageMission)
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

func handleSkills(args []interface{}) {
	rsp := args[0].(*pb.GS2C_GetSkills)
	log.Release("%v get skill success", rsp)
}

func handleRoles(args []interface{}) {
	rsp := args[0].(*pb.GS2C_GetRoles)
	log.Release("%v get roles success", rsp)
}

func handleGetUserTeams(args []interface{}) {
	rsp := args[0].(*pb.GS2C_GetUserTeams)
	log.Release("%v get team  success", rsp)
}

func handleGetUserRoles(args []interface{}) {
	rsp := args[0].(*pb.GS2C_GetUserRoles)
	log.Release("%v get user role  success", rsp)

}

func handleEditUserTeam(args []interface{}) {
	rsp := args[0].(*pb.GS2C_EditUserTeam)
	log.Release("%v edit team  success", rsp)

}

func handleHeartbeat(args []interface{}) {
	rsp := args[0].(*pb.GS2C_Heartbeat)
	log.Release("%v heart beat success", rsp)
}

func handleUpdateUserProps(args []interface{}) {
	rsp := args[0].(*pb.GS2C_UpdateUserProps)
	log.Release("%v get user props success", rsp)
}

func handleGetUserProps(args []interface{}) {
	rsp := args[0].(*pb.GS2C_GetUserProps)
	if rsp.Code == 200 {
		log.Release("%v get user props success", rsp)
	} else {
		log.Release("get user props fail: %s", code.Text(rsp.Code))
	}
}

func handleRoleUpgrade(args []interface{}) {
	rsp := args[0].(*pb.GS2C_RoleUpgrade)
	if rsp.Code == 200 {
		log.Release("%v role upgrade success", rsp)
	} else {
		log.Release("role upgrade fail: %s", code.Text(rsp.Code))
	}
}

func handleRoleAdvanced(args []interface{}) {
	rsp := args[0].(*pb.GS2C_RoleAdvanced)
	if rsp.Code == 200 {
		log.Release("%v role advanced success", rsp)
	} else {
		log.Release("role advanced fail: %s", code.Text(rsp.Code))
	}
}

func handleGetUserStageMission(args []interface{}) {
	rsp := args[0].(*pb.GS2C_GetUserStageMission)
	if rsp.Code == 200 {
		log.Release("%v get user stage mission success", rsp)
	} else {
		log.Release("get user stage mission fail: %s", code.Text(rsp.Code))
	}
}

func handleUpdateUserStageMission(args []interface{}) {
	rsp := args[0].(*pb.GS2C_UpdateUserStageMission)
	if rsp.Code == 200 {
		log.Release("%v update user stage mission success", rsp)
	} else {
		log.Release("update user stage mission fail: %s", code.Text(rsp.Code))
	}
}
