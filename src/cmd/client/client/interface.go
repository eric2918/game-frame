package client

import (
	"errors"
	"frame/conf"
	"frame/pb"
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
