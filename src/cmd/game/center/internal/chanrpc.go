package internal

import (
	"frame/conf"

	"github.com/eric2918/leaf/cluster"
	"github.com/eric2918/leaf/gate"
	"github.com/eric2918/leaf/log"
)

var (
	clientCount     = 0
	accountAgentMap = map[string]gate.Agent{}
	playerAgentMap  = map[string]gate.Agent{}
)

func handleRpc(id interface{}, f interface{}) {
	cluster.SetRoute(id, ChanRPC)
	skeleton.RegisterChanRPC(id, f)
}

func init() {
	skeleton.RegisterChanRPC("NewAgent", NewAgent)
	skeleton.RegisterChanRPC("CloseAgent", CloseAgent)

	skeleton.RegisterChanRPC("AccountOnline", AccountOnline)
	skeleton.RegisterChanRPC("AccountOffline", AccountOffline)
	skeleton.RegisterChanRPC("PlayerOnline", PlayerOnline)
	skeleton.RegisterChanRPC("PlayerOffline", PlayerOffline)

	handleRpc("GetGameInfo", GetGameInfo)

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

func GetGameInfo(args []interface{}) ([]interface{}, error) {
	return []interface{}{0, conf.Server.MaxConnNum, conf.Server.TCPAddr, conf.Server.Region}, nil
}

func AccountOnline(args []interface{}) (interface{}, error) {
	accountId := args[0].(string)
	agent := args[1].(gate.Agent)
	if oldAgent, ok := accountAgentMap[accountId]; ok {
		delete(accountAgentMap, accountId)
		oldAgent.Destroy()
		return false, nil
	} else {
		accountAgentMap[accountId] = agent

		clientCount += 1
		cluster.Go("gateway", "UpdateGameInfo", conf.Server.Region, conf.Server.ServerName, clientCount)

		log.Debug("%v account is online", accountId)
		return true, nil
	}
}

func AccountOffline(args []interface{}) {
	accountId := args[0].(string)
	agent := args[1].(gate.Agent)
	oldAgent, ok := accountAgentMap[accountId]
	if ok && agent == oldAgent {
		delete(accountAgentMap, accountId)

		clientCount -= 1
		cluster.Go("gateway", "UpdateGameInfo", conf.Server.Region, conf.Server.ServerName, clientCount)

		log.Debug("%v account is offline", accountId)
	}
}

func PlayerOnline(args []interface{}) {
	playerId := args[0].(string)
	agent := args[1].(gate.Agent)
	playerAgentMap[playerId] = agent
	log.Debug("%v player is online", playerId)
}

func PlayerOffline(args []interface{}) {
	playerId := args[0].(string)
	agent := args[1].(gate.Agent)
	oldAgent, ok := playerAgentMap[playerId]
	if ok && agent == oldAgent {
		delete(playerAgentMap, playerId)
		log.Debug("%v player is offline", playerId)
	}
}
