package msg

const (
	C2AS_LOGIN = 30101
	AS2C_LOGIN = 30102

	C2GS_CHECK_LOGIN = 40101
	GS2C_CHECK_LOGIN = 40102

	C2GS_CREATE_PLAYER = 40201
	GS2C_CREATE_PLAYER = 40202

	C2GS_GET_USER_TEAMS = 40301
	GS2C_GET_USER_TEAMS = 40302
	C2GS_GET_USER_ROLES = 40303
	GS2C_GET_USER_ROLES = 40304
	C2GS_EDIT_USER_TEAM = 40305
	GS2C_EDIT_USER_TEAM = 40306

	GS2C_UPDATE_USER_PROPS = 40401
	C2GS_GET_USER_PROPS    = 40402
	GS2C_GET_USER_PROPS    = 40403

	C2GS_ROLE_UPGRADE  = 40501
	GS2C_ROLE_UPGRADE  = 40502
	C2GS_ROLE_ADVANCED = 40503
	GS2C_ROLE_ADVANCED = 40504

	C2GS_GET_USER_STAGE_MISSION    = 40601
	GS2C_GET_USER_STAGE_MISSION    = 40602
	C2GS_UPDATE_USER_STAGE_MISSION = 40603
	GS2C_UPDATE_USER_STAGE_MISSION = 40604

	//C2GS_SEND_MSG = 40101
	//GS2C_SEND_MSG = 40102
	//GS2C_MSG_LIST = 40103

	C2GS_HEARTBEAT  = 10001 // 心跳请求
	GS2C_HEARTBEAT  = 10002 // 心跳返回
	GS2C_KICKPLAYER = 10003 // 踢玩家下线
)
