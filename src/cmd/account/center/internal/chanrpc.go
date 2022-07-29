package internal

import (
	"errors"
	"fmt"
	"frame/pkg/jwt"
	"frame/pkg/redis"

	"github.com/eric2918/leaf/gate"

	"github.com/eric2918/leaf/cluster"
)

func handleRpc(id interface{}, f interface{}) {
	cluster.SetRoute(id, ChanRPC)
	skeleton.RegisterChanRPC(id, f)
}

func init() {
	skeleton.RegisterChanRPC("NewAgent", NewAgent)
	skeleton.RegisterChanRPC("CloseAgent", CloseAgent)

	handleRpc("CheckToken", CheckToken)
}

func NewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	_ = a
}

func CloseAgent(args []interface{}) error {
	a := args[0].(gate.Agent)
	_ = a
	return nil
}

func CheckToken(args []interface{}) (res []interface{}, err error) {
	tokenStr := args[0].(string)

	j := jwt.NewJWT()
	claims, err := j.ParserToken(tokenStr)
	if err != nil {
		return nil, err
	}

	key := fmt.Sprintf("token_%s", claims.AccountId)
	token, err := redis.Client.Get(key).Result()
	if err != nil {
		return nil, err
	}
	if token != tokenStr {
		return nil, errors.New("invalid token")
	}

	return []interface{}{claims.AccountId, claims.Username}, nil
}
