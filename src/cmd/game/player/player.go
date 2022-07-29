package player

import (
	"frame/pb"
	"frame/pkg/mongo"
	"frame/pkg/snowflake"
	"time"

	"github.com/eric2918/leaf/log"
	"gopkg.in/mgo.v2/bson"
)

type Player struct {
	Player *pb.Player
}

func New(player *pb.Player) *Player {
	p := &Player{
		Player: player,
	}
	return p
}

func (p *Player) Init(accountId string, nickname string) {
	// 初始化用户编队
	timestamp := time.Now().Unix()
	p.Player = &pb.Player{
		PlayerId:          snowflake.GenStr(),
		AccountId:         accountId,
		Nickname:          nickname,
		Avatar:            "",
		Level:             1,
		Exp:               0,
		WorldLevel:        1,
		CreateAt:          timestamp,
		LastHeartbeatTime: timestamp,
	}
}

func (p *Player) UpdateLastHeartbeat() (err error) {
	// 更新最后心跳时间
	selector := bson.M{"playerId": p.Player.PlayerId}
	update := bson.M{"$set": bson.M{"lastHeartbeatTime": p.Player.LastHeartbeatTime}}
	if err = mongo.Collection(mongo.GAME_DB, mongo.PLAYERS_COLLECTION).Update(selector, update); err != nil {
		log.Error("update lastHeartbeatTime error:%v", err)
		return
	}
	return
}
