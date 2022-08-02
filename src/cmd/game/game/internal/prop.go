package internal

import (
	"frame/cmd/game/game/player"
	"frame/pb"

	"github.com/eric2918/leaf/gate"
)

func init() {
	handle(&pb.C2GS_GetUserProps{}, handleGetUserProps)
}

func handleGetUserProps(args []interface{}) {
	agent := args[1].(gate.Agent)

	playerData := agent.UserData().(*player.Player)

	sendMsg := &pb.GS2C_GetUserProps{
		Code: 200,
		Data: nil,
	}

	sendMsg.Data = playerData.Player.Props

	agent.WriteMsg(sendMsg)
}
