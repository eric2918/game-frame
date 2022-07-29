package code

var ZhCnText = map[int64]string{
	ServerInternalError: "服务内部错误",
	TooManyRequests:     "请求过多",
	ParamBindError:      "参数信息错误",
	SignatureError:      "签名信息错误",
	ResubmitError:       "请勿重复提交",
	SendEmailError:      "发送邮件失败",

	InvalidUsernameOrPassword: "无效的用户名或密码",
	UsernameExist:             "用户名已存在",
	UserNotExist:              "用户不存在",
	AccountDisabled:           "帐户已禁用",
	NoGameServerAlloc:         "没有可分配的游戏服务器",

	UserRoleIdError: "玩家角色ID有误",
	SelectPropError: "选择道具有误",
	RoleMaxLevel:    "当前已是最大等级",
	PropNotEnough:   "道具数量不足",
	NotAdvanced:     "升级到最大等级之后才可进阶",
	MaxAdvanced:     "已进阶到当前职业最高阶",
	NotCareer:       "无可转职的职业",
	NotTransfer:     "选择职业不可转职",
}
