package player

import (
	"frame/cmd/game/game/stage"
	"frame/cmd/game/game/static"
	"frame/conf"
	"frame/pb"
	"frame/pkg/mongo"
	"frame/pkg/snowflake"
	"time"

	"github.com/eric2918/leaf/log"
	"github.com/spf13/cast"
	"gopkg.in/mgo.v2/bson"
)

type Player struct {
	Player     *pb.Player
	PropsIndex map[int64]int
	RolesIndex map[string]int
}

func New(player *pb.Player) *Player {
	p := &Player{
		Player: player,
	}
	p.Index()
	return p
}

func (p *Player) Index() {
	p.RolesIndex = make(map[string]int)
	for i, role := range p.Player.Roles {
		p.RolesIndex[role.UserRoleId] = i
	}

	p.PropsIndex = make(map[int64]int)
	for i, prop := range p.Player.Props {
		p.PropsIndex[prop.ItemId] = i
	}
}

func (p *Player) Init(accountId string, nickname string) {
	var userRoles []*pb.UserRole
	var userRoleIds []string
	for _, role := range static.Data.RoleBasic {
		if role.IsForbidden {
			continue
		}

		var skinId int64 = 0
		if len(static.Data.RoleSkins[role.RoleId]) > 0 {
			skinId = static.Data.RoleSkins[role.RoleId][0]
		}

		// 获取角色职业Id
		var careerId int64 = 0
		if len(role.CareerIds) > 0 {
			careerId = role.CareerIds[0]
		}

		// 获取角色进阶ID
		career, ok := static.Data.Career[careerId]
		if !ok {
			continue
		}
		advancementId := career.Advancement[0]

		advancement, ok := static.Data.CareerAdvancement[advancementId]
		if !ok {
			continue
		}

		// 获取角色当前阶段最大等级
		maxLevel := advancement.SetRoleLvCap

		userRole := &pb.UserRole{
			UserRoleId:    snowflake.GenStr(),
			RoleId:        role.RoleId,
			Level:         1,
			SkinId:        skinId,
			CareerId:      careerId,
			AdvancementId: advancementId,
			MaxLevel:      maxLevel,
			Exp:           0,
		}
		userRole.UnlockCareer = make(map[string]int64)
		userRole.UnlockCareer[cast.ToString(careerId)] = 0

		userRoles = append(userRoles, userRole)
		userRoleIds = append(userRoleIds, userRole.UserRoleId)
	}

	var userTeams []*pb.UserTeam
	for i := 1; i <= conf.Server.MaxTeamsCounts; i++ {
		userTeams = append(userTeams, &pb.UserTeam{
			TeamId: snowflake.GenStr(),
		})
	}

	var userProps []*pb.UserProp
	for _, basic := range static.Data.PropBasic {
		var counts int64 = 10000
		switch basic.Type {
		case 1:
			counts = 3000
		case 2:
			counts = 200
		case 3:
			counts = 5000
		case 4:
			counts = 1000
		case 7:
			counts = 10
		}
		userProps = append(userProps, &pb.UserProp{
			ItemId: basic.ItemId,
			Counts: counts,
		})
	}

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
		Roles:             userRoles,
		Teams:             userTeams,
		Props:             userProps,
		CreateAt:          timestamp,
		LastHeartbeatTime: timestamp,
	}

	// 解锁关卡
	s := stage.New(p.Player.StageMission, static.Data.StageMission)
	s.Unlock([]string{"10001010"})
	p.Player.StageMission = s.Mission

	// 设置索引
	p.Index()
}

// UnlockCareer 可解锁职业
func UnlockCareer(currentCareerId int64, careerIds []int64) map[int64]bool {
	index := 0
	for i, id := range careerIds {
		if id == currentCareerId {
			index = i
		}
	}

	resMap := make(map[int64]bool)
	switch index {
	case 0, 2, 4, 5:
		resMap = map[int64]bool{
			careerIds[1]: true,
			careerIds[3]: true,
			careerIds[5]: true,
		}
	case 1:
		resMap = map[int64]bool{
			careerIds[2]: true,
			careerIds[3]: true,
			careerIds[5]: true,
		}
	case 3:
		resMap = map[int64]bool{
			careerIds[4]: true,
			careerIds[3]: true,
			careerIds[5]: true,
		}
	case 6:
		resMap = map[int64]bool{
			careerIds[7]: true,
			careerIds[3]: true,
			careerIds[5]: true,
		}
	}
	return resMap
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
