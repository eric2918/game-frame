package gate

import (
	"frame/cmd/account/login"
	"frame/msg"
	"frame/pb"
)

func init() {
	msg.Processor.SetRouter(&pb.C2AS_Login{}, login.ChanRPC)
}
