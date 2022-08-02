package msg

import (
	"frame/pb"

	"github.com/eric2918/leaf/network/protobuf"
)

var (
	Processor = protobuf.NewProcessor()
)

func init() {
	Processor.Register(C2AS_LOGIN, &pb.C2AS_Login{})
	Processor.Register(AS2C_LOGIN, &pb.AS2C_Login{})

	Processor.Register(C2GS_CHECK_LOGIN, &pb.C2GS_CheckLogin{})
	Processor.Register(GS2C_CHECK_LOGIN, &pb.GS2C_CheckLogin{})
	Processor.Register(C2GS_CREATE_PLAYER, &pb.C2GS_CreatePlayer{})
	Processor.Register(GS2C_CREATE_PLAYER, &pb.GS2C_CreatePlayer{})

	Processor.Register(C2GS_GET_USER_TEAMS, &pb.C2GS_GetUserTeams{})
	Processor.Register(GS2C_GET_USER_TEAMS, &pb.GS2C_GetUserTeams{})
	Processor.Register(C2GS_GET_USER_ROLES, &pb.C2GS_GetUserRoles{})
	Processor.Register(GS2C_GET_USER_ROLES, &pb.GS2C_GetUserRoles{})
	Processor.Register(C2GS_EDIT_USER_TEAM, &pb.C2GS_EditUserTeam{})
	Processor.Register(GS2C_EDIT_USER_TEAM, &pb.GS2C_EditUserTeam{})

	Processor.Register(GS2C_UPDATE_USER_PROPS, &pb.GS2C_UpdateUserProps{})
	Processor.Register(C2GS_GET_USER_PROPS, &pb.C2GS_GetUserProps{})
	Processor.Register(GS2C_GET_USER_PROPS, &pb.GS2C_GetUserProps{})
	Processor.Register(C2GS_ROLE_UPGRADE, &pb.C2GS_RoleUpgrade{})
	Processor.Register(GS2C_ROLE_UPGRADE, &pb.GS2C_RoleUpgrade{})
	Processor.Register(C2GS_ROLE_ADVANCED, &pb.C2GS_RoleAdvanced{})
	Processor.Register(GS2C_ROLE_ADVANCED, &pb.GS2C_RoleAdvanced{})

	Processor.Register(C2GS_GET_USER_STAGE_MISSION, &pb.C2GS_GetUserStageMission{})
	Processor.Register(GS2C_GET_USER_STAGE_MISSION, &pb.GS2C_GetUserStageMission{})
	Processor.Register(C2GS_UPDATE_USER_STAGE_MISSION, &pb.C2GS_UpdateUserStageMission{})
	Processor.Register(GS2C_UPDATE_USER_STAGE_MISSION, &pb.GS2C_UpdateUserStageMission{})

	Processor.Register(C2GS_HEARTBEAT, &pb.C2GS_Heartbeat{})
	Processor.Register(GS2C_HEARTBEAT, &pb.GS2C_Heartbeat{})

}
