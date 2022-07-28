package internal

import (
	"errors"
	"math"

	"github.com/eric2918/leaf/log"

	"github.com/eric2918/leaf/cluster"
)

func handleRpc(id interface{}, f interface{}) {
	cluster.SetRoute(id, ChanRPC)
	skeleton.RegisterChanRPC(id, f)
}

func init() {
	cluster.AgentChanRPC = ChanRPC
	skeleton.RegisterChanRPC("NewServerAgent", NewServerAgent)
	skeleton.RegisterChanRPC("CloseServerAgent", CloseServerAgent)

	handleRpc("GetBestFrontInfo", GetBestGameInfo)
	handleRpc("UpdateFrontInfo", UpdateGameInfo)
	handleRpc("UpdateChatInfo", UpdateChatInfo)
}

func NewServerAgent(args []interface{}) {
	serverName := args[0].(string)
	agent := args[1].(*cluster.Agent)

	switch serverName[:4] {
	case "game":
		res, err := agent.CallN("GetGameInfo")
		if err != nil {
			log.Error("GetGameInfo is error: %v", err)
			return
		}

		gameInfo := &GameInfo{
			serverName:     serverName,
			clientCount:    res[0].(int),
			maxClientCount: res[1].(int),
			clientAddr:     res[2].(string),
		}
		region := res[3].(string)
		if gameInfoMap[region] == nil {
			gameInfoMap[region] = make(map[string]*GameInfo)
		}
		gameInfoMap[region][serverName] = gameInfo
		gameRegion[serverName] = region
	case "chat":
		res, err := agent.CallN("GetChatInfo")
		if err != nil {
			log.Error("GetChatInfo is error: %v", err)
			return
		}

		chatInfo := &ChatInfo{
			serverName:  serverName,
			clientCount: res[0].(int),
			clientAddr:  res[1].(string),
		}
		region := res[2].(string)

		if chatInfoMap[region] == nil {
			chatInfoMap[region] = make(map[string]*ChatInfo)
		}
		chatInfoMap[region][serverName] = chatInfo
		chatRegion[serverName] = region
	}
}

func CloseServerAgent(args []interface{}) {
	serverName := args[0].(string)
	switch serverName[:4] {
	case "game":
		region, ok := gameRegion[serverName]
		if !ok {
			return
		}
		delete(gameInfoMap[region], serverName)
		delete(gameRegion, serverName)
	case "chat":
		region, ok := chatRegion[serverName]
		if !ok {
			return
		}
		delete(chatInfoMap[region], serverName)
		delete(chatRegion, serverName)
	}
}

func GetBestGameInfo(args []interface{}) ([]interface{}, error) {
	region := args[0].(string)
	accountId := args[1].(string)

	var ok bool
	var gameInfo *GameInfo
	if gameInfo, ok = accountGameMap[accountId]; !ok {
		minClientCount := math.MaxInt32
		for _, _gameInfo := range gameInfoMap[region] {
			if _gameInfo.clientCount < minClientCount && _gameInfo.clientCount < _gameInfo.maxClientCount {
				gameInfo = _gameInfo
			}
		}
	}

	if gameInfo == nil {
		return []interface{}{}, errors.New("no game server to alloc")
	} else {
		accountGameMap[accountId] = gameInfo
		log.Debug("%v account ask front info", accountId)
		return []interface{}{gameInfo.serverName, gameInfo.clientAddr}, nil
	}
}

func UpdateGameInfo(args []interface{}) {
	region := args[0].(string)
	serverName := args[1].(string)
	clientCount := args[2].(int)
	frontInfo, ok := gameInfoMap[region][serverName]
	if ok {
		frontInfo.clientCount = clientCount
		log.Debug("%v server of client count is %v", serverName, clientCount)
	}
}

func UpdateChatInfo(args []interface{}) {
	region := args[0].(string)
	serverName := args[1].(string)
	clientCount := args[2].(int)
	chatInfo, ok := chatInfoMap[region][serverName]
	if ok {
		chatInfo.clientCount = clientCount
		log.Debug("%v server of client count is %v", serverName, clientCount)
	}
}
