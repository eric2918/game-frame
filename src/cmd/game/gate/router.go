package gate

import (
	"frame/cmd/game/game"
	"frame/msg"
	"frame/pb"
)

func init() {
	msg.Processor.SetRouter(&pb.C2GS_CheckLogin{}, game.ChanRPC)
	msg.Processor.SetRouter(&pb.C2GS_CreatePlayer{}, game.ChanRPC)

	msg.Processor.SetRouter(&pb.C2GS_GetUserTeams{}, game.ChanRPC)
	msg.Processor.SetRouter(&pb.C2GS_GetUserRoles{}, game.ChanRPC)
	msg.Processor.SetRouter(&pb.C2GS_EditUserTeam{}, game.ChanRPC)
	msg.Processor.SetRouter(&pb.C2GS_Heartbeat{}, game.ChanRPC)
	msg.Processor.SetRouter(&pb.C2GS_GetUserProps{}, game.ChanRPC)
	msg.Processor.SetRouter(&pb.C2GS_RoleUpgrade{}, game.ChanRPC)
	msg.Processor.SetRouter(&pb.C2GS_RoleAdvanced{}, game.ChanRPC)
	msg.Processor.SetRouter(&pb.C2GS_GetUserStageMission{}, game.ChanRPC)
	msg.Processor.SetRouter(&pb.C2GS_UpdateUserStageMission{}, game.ChanRPC)
}
