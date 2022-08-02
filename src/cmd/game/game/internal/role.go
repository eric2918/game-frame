package internal

import (
	"fmt"
	"frame/cmd/game/game/player"
	"frame/cmd/game/game/static"
	"frame/pb"
	"frame/pkg/code"
	"frame/pkg/mongo"
	"frame/pkg/tools"
	"strings"

	"github.com/eric2918/leaf/gate"
	"github.com/spf13/cast"
	"gopkg.in/mgo.v2/bson"
)

func init() {
	handle(&pb.C2GS_GetUserRoles{}, handleGetUserRoles)
	handle(&pb.C2GS_RoleUpgrade{}, handleRoleUpgrade)
	handle(&pb.C2GS_RoleAdvanced{}, handleRoleAdvanced)
}

func handleGetUserRoles(args []interface{}) {
	agent := args[1].(gate.Agent)

	playerData := agent.UserData().(*player.Player)

	sendMsg := &pb.GS2C_GetUserRoles{
		Code: 200,
		Data: nil,
	}

	sendMsg.Data = playerData.Player.Roles

	agent.WriteMsg(sendMsg)
}

func handleRoleUpgrade(args []interface{}) {
	req := args[0].(*pb.C2GS_RoleUpgrade)
	agent := args[1].(gate.Agent)

	sendMsg := &pb.GS2C_RoleUpgrade{
		Code: 200,
	}

	playerData := agent.UserData().(*player.Player)

	// 获取玩家角色信息
	var roleIndex int
	var ok bool
	if roleIndex, ok = playerData.RolesIndex[req.UserRoleId]; !ok {
		sendMsg.Code = code.UserRoleIdError
		agent.WriteMsg(sendMsg)
		return
	}

	// 加锁，防并发
	key := fmt.Sprintf("role_upgrade_%s", playerData.Player.PlayerId)
	if errCode := tools.Locker(key); errCode != 200 {
		sendMsg.Code = errCode
		agent.WriteMsg(sendMsg)
		return
	}
	defer tools.UnLocker(key)

	// 玩家角色
	playerRoles := playerData.Player.Roles

	// 已升级到当前阶段最高等级
	if playerRoles[roleIndex].MaxLevel < playerRoles[roleIndex].Level+1 {
		sendMsg.Code = code.RoleMaxLevel
		agent.WriteMsg(sendMsg)
		return
	}

	// 升级消耗金币
	var usedGold int64 = 0
	var needGold int64 = 0
	var usedExp int64 = 0

	updateProps := make(map[int64]int64)

	satisfy := true
	selectErr := false
	playerProps := playerData.Player.Props
	for _, used := range req.Used {
		// 验证素材是否充足
		i := playerData.PropsIndex[used.ItemId]
		if playerProps[i].Counts < used.Counts {
			satisfy = false
			break
		}
		// 足够扣除道具
		playerProps[i].Counts -= used.Counts

		updateProps[used.ItemId] = playerProps[i].Counts

		// 计算消耗是否足够
		propIndex := static.Data.PropIndex[used.ItemId]
		prop := static.Data.PropBasic[propIndex]
		switch prop.Type {
		case 3:
			usedGold += used.Counts
		case 8:
			usedItems := strings.Split(prop.UsedCondition, ",")
			if len(usedItems) >= 2 && usedItems[0] == "10" {
				needGold += cast.ToInt64(usedItems[1]) * used.Counts
			}
			usedExp += prop.HeroExp * used.Counts
		default:
			selectErr = true
			break
		}
	}

	if !satisfy {
		sendMsg.Code = code.PropNotEnough
		agent.WriteMsg(sendMsg)
		return
	}

	if usedGold < needGold || selectErr {
		sendMsg.Code = code.SelectPropError
		agent.WriteMsg(sendMsg)
		return
	}

	// 计算增加经验、等级
	playerRoles[roleIndex].Exp += usedExp
	for {
		levelInfo, _ := static.Data.LevelBasic[playerRoles[roleIndex].Level+1]
		if playerRoles[roleIndex].Exp < levelInfo.LevelExp || playerRoles[roleIndex].MaxLevel < playerRoles[roleIndex].Level {
			break
		}
		playerRoles[roleIndex].Level += 1
		if playerRoles[roleIndex].Level == playerRoles[roleIndex].MaxLevel {
			playerRoles[roleIndex].Exp = levelInfo.LevelExp
		}
	}
	// 返回当前等级和经验
	sendMsg.Level = playerRoles[roleIndex].Level
	sendMsg.Exp = playerRoles[roleIndex].Exp

	// 更新数据
	selector := bson.M{"playerId": playerData.Player.PlayerId}
	update := bson.M{"$set": bson.M{"roles": playerRoles, "props": playerProps}}
	if err := mongo.Collection(mongo.GAME_DB, mongo.PLAYERS_COLLECTION).Update(selector, update); err != nil {
		fmt.Println("error:", err.Error())
	}

	// 设置缓存
	playerData.Player.Roles = playerRoles
	playerData.Player.Props = playerProps
	agent.SetUserData(playerData)

	// 给客户端发送道具更细消息
	if len(updateProps) > 0 {
		var updateUserProps []*pb.UserProp
		for itemId, counts := range updateProps {
			updateUserProps = append(updateUserProps, &pb.UserProp{
				ItemId: itemId,
				Counts: counts,
			})
		}
		agent.WriteMsg(&pb.GS2C_UpdateUserProps{
			Data: updateUserProps,
		})
	}

	agent.WriteMsg(sendMsg)
}

func handleRoleAdvanced(args []interface{}) {
	req := args[0].(*pb.C2GS_RoleAdvanced)
	agent := args[1].(gate.Agent)

	sendMsg := &pb.GS2C_RoleAdvanced{
		Code: 200,
	}

	playerData := agent.UserData().(*player.Player)

	// 获取玩家角色信息
	var roleIndex int
	var ok bool
	if roleIndex, ok = playerData.RolesIndex[req.UserRoleId]; !ok {
		sendMsg.Code = code.UserRoleIdError
		agent.WriteMsg(sendMsg)
		return
	}

	// 加锁，防并发
	key := fmt.Sprintf("role_advanced_%s", playerData.Player.PlayerId)
	if errCode := tools.Locker(key); errCode != 200 {
		sendMsg.Code = errCode
		agent.WriteMsg(sendMsg)
		return
	}
	defer tools.UnLocker(key)

	// 玩家角色
	playerRole := playerData.Player.Roles

	// 转职当前职业，不做处理
	if playerRole[roleIndex].CareerId == req.CareerId {
		agent.WriteMsg(sendMsg)
		return
	}

	// 转职回转职过的职业，不消耗材料，回到之前的阶段
	filter := false
	cost := true
	var advancementIndex int64 = 0
	if req.CareerId > 0 {
		// 转职

		i, ok := playerRole[roleIndex].UnlockCareer[cast.ToString(req.CareerId)]
		if ok {
			playerRole[roleIndex].CareerId = req.CareerId

			// 职业已解锁过，不消耗转职材料
			cost = false

			// 计算阶段ID索引
			advancementIndex = i
		} else {
			// 职业未解锁 ，判断是否是第一次转职，第一次转职不消耗魔法石
			if len(playerRole[roleIndex].UnlockCareer) < 3 {
				filter = true
			}

			careerIds, ok1 := static.Data.RoleCareerIds[playerRole[roleIndex].RoleId]
			if !ok1 {
				// 无可转职的职业
				sendMsg.Code = code.NotCareer
				agent.WriteMsg(sendMsg)
				return
			}

			res := player.UnlockCareer(playerRole[roleIndex].CareerId, careerIds)
			if _, ok2 := res[req.CareerId]; !ok2 {
				// 当前职业不可转职
				sendMsg.Code = code.NotTransfer
				agent.WriteMsg(sendMsg)
				return
			}

			playerRole[roleIndex].CareerId = req.CareerId

			// 默认阶段ID索引
			advancementIndex = 0
		}
	} else {
		if playerRole[roleIndex].Level < playerRole[roleIndex].MaxLevel {
			// 升级到最大等级之后才可进阶
			sendMsg.Code = code.NotAdvanced
			agent.WriteMsg(sendMsg)
			return
		}
		// 进阶，计算当前阶段ID索引
		i, ok := playerRole[roleIndex].UnlockCareer[cast.ToString(playerRole[roleIndex].CareerId)]
		if !ok {
			sendMsg.Code = code.ServerInternalError
			agent.WriteMsg(sendMsg)
			return
		}
		if 2 < i+1 {
			sendMsg.Code = code.MaxAdvanced
			agent.WriteMsg(sendMsg)
			return
		}

		// 当前阶段ID索引+1
		advancementIndex = i + 1
	}

	// 角色场景信息
	career := static.Data.Career[playerRole[roleIndex].CareerId]

	// 角色下个阶段ID
	playerRole[roleIndex].AdvancementId = career.Advancement[advancementIndex]

	// 添加/更新解锁阶段索引
	playerRole[roleIndex].UnlockCareer[cast.ToString(playerRole[roleIndex].CareerId)] = advancementIndex

	// 角色下个阶段信息
	advancement, ok := static.Data.CareerAdvancement[playerRole[roleIndex].AdvancementId]
	if !ok {
		//	阶段信息有误
		sendMsg.Code = code.ServerInternalError
		agent.WriteMsg(sendMsg)
		return
	}

	// 角色最大等级
	if playerRole[roleIndex].MaxLevel < advancement.SetRoleLvCap {
		playerRole[roleIndex].MaxLevel = advancement.SetRoleLvCap
	}

	updateProps := make(map[int64]int64)

	// 解析消耗材料
	satisfy := true
	playerProps := playerData.Player.Props
	if cost {
		// 进阶消耗材料
		advanceCosts := advancement.AdvanceCost
		for _, advanceCost := range advanceCosts {
			if filter && advanceCost.PropID == 15 {
				continue
			}
			// 判断玩家道具是否充足
			i := playerData.PropsIndex[advanceCost.PropID]
			if playerProps[i].Counts < advanceCost.PropNum {
				satisfy = false
				break
			}
			playerProps[i].Counts -= advanceCost.PropNum

			updateProps[advanceCost.PropID] = playerProps[i].Counts
		}
	}
	if !satisfy {
		sendMsg.Code = code.PropNotEnough
		agent.WriteMsg(sendMsg)
		return
	}

	// 进阶/转职后返回当前职业、阶段、最大等级
	sendMsg.CareerId = playerRole[roleIndex].CareerId
	sendMsg.AdvancementId = playerRole[roleIndex].AdvancementId
	sendMsg.MaxLevel = playerRole[roleIndex].MaxLevel

	// 更新数据
	selector := bson.M{"playerId": playerData.Player.PlayerId}
	update := bson.M{"$set": bson.M{"roles": playerRole, "props": playerProps}}
	if err := mongo.Collection(mongo.GAME_DB, mongo.PLAYERS_COLLECTION).Update(selector, update); err != nil {
		fmt.Println("error:", err.Error())
	}

	// 设置缓存
	playerData.Player.Roles = playerRole
	playerData.Player.Props = playerProps
	agent.SetUserData(playerData)

	// 给客户端发送道具更细消息
	if len(updateProps) > 0 {
		var updateUserProps []*pb.UserProp
		for itemId, counts := range updateProps {
			updateUserProps = append(updateUserProps, &pb.UserProp{
				ItemId: itemId,
				Counts: counts,
			})
		}
		agent.WriteMsg(&pb.GS2C_UpdateUserProps{
			Data: updateUserProps,
		})
	}

	agent.WriteMsg(sendMsg)
}
