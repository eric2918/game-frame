package gate

import (
	"frame/cmd/game/game"
	"frame/msg"
	"frame/pb"
)

func init() {
	msg.Processor.SetRouter(&pb.C2GS_CheckLogin{}, game.ChanRPC)
	msg.Processor.SetRouter(&pb.C2GS_CreatePlayer{}, game.ChanRPC)
}
