package code

import "frame/conf"

const (
	// 10101
	ServerInternalError = 10100
	TooManyRequests     = 10101
	ParamBindError      = 10102
	SignatureError      = 10103
	ResubmitError       = 10104
	SendEmailError      = 10105

	InvalidUsernameOrPassword = 10200
	UsernameExist             = 10201
	UserNotExist              = 10202
	AccountDisabled           = 10203
	NoGameServerAlloc         = 20204

	UserRoleIdError = 30101
	SelectPropError = 30102
	RoleMaxLevel    = 30103
	PropNotEnough   = 30104
	NotAdvanced     = 30105
	MaxAdvanced     = 30106
	NotCareer       = 30107
	NotTransfer     = 30108
)

func Text(code int64) string {
	lang := conf.Server.Language
	if lang == "zh-cn" {
		return ZhCnText[code]
	} else {
		return EnUsText[code]
	}
	return ""
}
