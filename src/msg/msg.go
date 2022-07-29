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
}
