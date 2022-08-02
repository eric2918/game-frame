package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"frame/conf"
	"frame/pb"
	"frame/pkg/tools"

	"github.com/spf13/cast"
)

var (
	userData = UserData{}
)

type UserData struct {
	AccountId string
	PlayerId  string
	Username  string
	Nickname  string
}

func init() {
	skeleton.RegisterCommand("login", "login account: input accountName, password", login)

	//skeleton.RegisterCommand("getSkills", "get all skill", getAllSkills)
	//skeleton.RegisterCommand("getRoles", "get all role", getAllRoles)
	skeleton.RegisterCommand("getUserTeams", "get user team", getUserTeams)
	skeleton.RegisterCommand("editUserTeams", "edit user team", editUserTeam)
	skeleton.RegisterCommand("getUserRoles", "edit user role", getUserRoles)
	skeleton.RegisterCommand("heartbeat", "heart beat", heartbeat)

	skeleton.RegisterCommand("getUserProps", "get user props", getUserProps)
	skeleton.RegisterCommand("roleUpgrade", "role upgrade", roleUpgrade)
	skeleton.RegisterCommand("roleAdvanced", "role advanced", roleAdvanced)
	skeleton.RegisterCommand("getUserStageMission", "get user stage mission", getUserStageMission)
	skeleton.RegisterCommand("updateUserStageMission", "update user stage mission", updateUserStageMission)

}

func login(args []interface{}) (ret interface{}, err error) {
	ret = ""
	if len(args) < 2 {
		err = errors.New("args len is less than 2")
		return
	}

	username := args[0].(string)
	password := args[1].(string)
	region := args[2].(string)
	userData.Username = username

	Start(conf.Server.LoginAddr)
	msg := &pb.C2AS_Login{
		Username: username,
		Password: password,
		Region:   region,
	}
	Client.WriteMsg(msg)
	return
}

func getAllSkills(args []interface{}) (ret interface{}, err error) {
	if Client == nil {
		err = errors.New("net is disconnect, please input login cmd")
		return
	}

	msg := &pb.C2GS_GetSkills{}
	Client.WriteMsg(msg)
	return
}

func getAllRoles(args []interface{}) (ret interface{}, err error) {
	if Client == nil {
		err = errors.New("net is disconnect, please input login cmd")
		return
	}

	msg := &pb.C2GS_GetRoles{}
	Client.WriteMsg(msg)
	return
}

func getUserTeams(args []interface{}) (ret interface{}, err error) {
	fmt.Println("getUserTeams")
	if Client == nil {
		err = errors.New("net is disconnect, please input login cmd")
		return
	}

	msg := &pb.C2GS_GetUserTeams{}
	Client.WriteMsg(msg)
	return
}

func getUserRoles(args []interface{}) (ret interface{}, err error) {
	if Client == nil {
		err = errors.New("net is disconnect, please input login cmd")
		return
	}

	msg := &pb.C2GS_GetUserRoles{}
	Client.WriteMsg(msg)
	return
}

func editUserTeam(args []interface{}) (ret interface{}, err error) {
	if len(args) < 3 {
		err = errors.New("args len is less than 2")
		return
	}

	teamId := args[0].(string)
	teamName := args[1].(string)
	userRoleIds := args[2].(string)

	if Client == nil {
		err = errors.New("net is disconnect, please input edit team cmd")
		return
	}

	msg := &pb.C2GS_EditUserTeam{
		TeamName:    teamName,
		UserRoleIds: tools.StringToInt64(userRoleIds),
	}

	if teamId != "nil" {
		msg.TeamId = teamId
	}

	Client.WriteMsg(msg)
	return
}

func heartbeat(args []interface{}) (ret interface{}, err error) {
	if Client == nil {
		err = errors.New("net is disconnect, please input login cmd")
		return
	}

	msg := &pb.C2GS_Heartbeat{}
	Client.WriteMsg(msg)
	return
}

func getUserProps(args []interface{}) (ret interface{}, err error) {
	if Client == nil {
		err = errors.New("net is disconnect, please input login cmd")
		return
	}

	msg := &pb.C2GS_GetUserProps{}
	Client.WriteMsg(msg)
	return
}

func roleUpgrade(args []interface{}) (ret interface{}, err error) {
	if len(args) < 2 {
		err = errors.New("args len is less than 2")
		return
	}

	roleId := args[0].(string)
	used := args[1].(string)
	// roleUpgrade 28681574325161984 [{"ItemId":200023,"Counts":1},{"ItemId":10,"Counts":20}]
	// roleUpgrade 26583464614563840 [{"ItemId":200023,"Counts":10},{"ItemId":10,"Counts":200}]
	// roleUpgrade 26583464614563840 [{"ItemId":200023,"Counts":25},{"ItemId":10,"Counts":500}]
	// roleAdvanced 26583464614563840
	// roleAdvanced 26583464614563840 1000025

	var Used []*pb.C2GS_RoleUpgradeUsed
	if err = json.Unmarshal([]byte(used), &Used); err != nil {
		return nil, err
	}

	if Client == nil {
		err = errors.New("net is disconnect, please input upgrade cmd")
		return
	}

	msg := &pb.C2GS_RoleUpgrade{
		UserRoleId: roleId,
		Used:       Used,
	}

	Client.WriteMsg(msg)
	return
}

func roleAdvanced(args []interface{}) (ret interface{}, err error) {
	if len(args) < 1 {
		err = errors.New("args len is less than 2")
		return
	}

	if Client == nil {
		err = errors.New("net is disconnect, please input advanced cmd")
		return
	}

	userRoleId := args[0].(string)
	msg := &pb.C2GS_RoleAdvanced{
		UserRoleId: userRoleId,
	}
	if len(args) >= 2 {
		careerId := args[1].(string)
		if careerId != "" {
			msg.CareerId = cast.ToInt64(careerId)
		}
	}

	Client.WriteMsg(msg)
	return
}

func getUserStageMission(args []interface{}) (ret interface{}, err error) {
	if Client == nil {
		err = errors.New("net is disconnect, please input advanced cmd")
		return
	}

	msg := &pb.C2GS_GetUserStageMission{}
	Client.WriteMsg(msg)
	return
}

func updateUserStageMission(args []interface{}) (ret interface{}, err error) {
	if len(args) < 1 {
		err = errors.New("args len is less than 2")
		return
	}

	if Client == nil {
		err = errors.New("net is disconnect, please input advanced cmd")
		return
	}

	missionID := args[0].(string)
	msg := &pb.C2GS_UpdateUserStageMission{
		MissionID: cast.ToInt64(missionID),
	}
	Client.WriteMsg(msg)
	return
}
