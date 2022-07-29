package internal

import (
	"fmt"
	"frame/conf"
	"frame/pb"
	"frame/pkg/code"
	"frame/pkg/jwt"
	"frame/pkg/md5"
	"frame/pkg/mongo"
	"frame/pkg/rand"
	"frame/pkg/redis"
	"frame/pkg/snowflake"
	"reflect"
	"time"

	jwt2 "github.com/dgrijalva/jwt-go"
	"github.com/eric2918/leaf/cluster"
	"github.com/eric2918/leaf/gate"
	"github.com/eric2918/leaf/log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func handle(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handle(&pb.C2AS_Login{}, handleLogin)
}

func handleLogin(args []interface{}) {
	req := args[0].(*pb.C2AS_Login)
	agent := args[1].(gate.Agent)

	sendMsg := &pb.AS2C_Login{
		Code: 200,
	}
	sendErrFunc := func(codeErr int64) {
		sendMsg.Code = codeErr
		log.Error(code.Text(codeErr))
		agent.WriteMsg(sendMsg)
	}

	if req.Username == "" || req.Password == "" {
		sendErrFunc(code.InvalidUsernameOrPassword)
		return
	}

	var account pb.Account
	err := mongo.Collection(mongo.ACCOUNT_DB, mongo.ACCOUNTS_COLLECTION).Find(bson.M{"username": req.Username}).One(&account)
	if err == mgo.ErrNotFound {
		salt := rand.RandomString(6, 3)
		password := md5.Encode(req.Password + salt)

		account = pb.Account{
			AccountId: snowflake.GenStr(),
			Username:  req.Username,
			Password:  password,
			Salt:      salt,
			CreateAt:  time.Now().Unix(),
			Status:    0,
			Region:    req.Region,
		}

		err = mongo.Collection(mongo.ACCOUNT_DB, mongo.ACCOUNTS_COLLECTION).Insert(&account)
	}

	if err != nil {
		sendErrFunc(code.ServerInternalError)
		return
	} else if account.Password != md5.Encode(req.Password+account.Salt) {
		sendErrFunc(code.InvalidUsernameOrPassword)
		return
	}

	results, err := cluster.CallN("gateway", "GetBestGameInfo", req.Region, account.AccountId)
	if err != nil {
		sendErrFunc(code.NoGameServerAlloc)
		return
	}
	//gameName := results[0].(string)
	gameAddr := results[1].(string)

	// 生成token
	j := jwt.NewJWT()
	claims := jwt.CustomClaims{
		AccountId: account.AccountId,
		Username:  account.Username,
		StandardClaims: jwt2.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(conf.Server.JwtTimeout) * time.Second).Unix(),
		},
	}

	token, err := j.CreateToken(claims)
	if err != nil {
		sendErrFunc(code.ServerInternalError)
		return
	}

	// 保存到redis
	key := fmt.Sprintf("token_%s", claims.AccountId)
	if err = redis.Client.Set(key, token, time.Duration(conf.Server.JwtTimeout)*time.Second).Err(); err != nil {
		sendErrFunc(code.InvalidUsernameOrPassword)
		return
	}

	sendMsg.Data = &pb.AS2C_LoginData{
		AccountId: claims.AccountId,
		GameAddr:  gameAddr,
		Token:     token,
	}
	agent.WriteMsg(sendMsg)
}
