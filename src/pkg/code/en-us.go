package code

var EnUsText = map[int64]string{
	ServerInternalError: "Service internal error",
	TooManyRequests:     "Too many requests",
	ParamBindError:      "Parameter error",
	SignatureError:      "Signature error",
	ResubmitError:       "Please do not submit repeatedly",
	SendEmailError:      "Failed to send mail",

	InvalidUsernameOrPassword: "invalid username or password",
	UsernameExist:             "Username already exists",
	UserNotExist:              "User not exist",
	AccountDisabled:           "Account is disabled",
	NoGameServerAlloc:         "No game server to alloc",

	UserRoleIdError: "Player role id is wrong",
	SelectPropError: "Select prop error",
	RoleMaxLevel:    "Role maximum level",
	PropNotEnough:   "Prop not enough",
	NotAdvanced:     "You can only advance after upgrading to the maximum level",
	MaxAdvanced:     "Advanced to the highest level of current profession",
	NotCareer:       "No transferable occupation",
	NotTransfer:     "Choose an occupation and cannot be transferred",
}
