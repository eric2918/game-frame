package internal

import (
	"fmt"
)

func init() {
	skeleton.RegisterCommand("getGameInfo", "return all game info", getGameInfo)
	skeleton.RegisterCommand("getChatInfo", "return all chat info", getChatInfo)
	skeleton.RegisterCommand("getBestGameInfo", "get best game info", getBestGameInfo)
}

func getGameInfo(args []interface{}) (ret interface{}, err error) {
	ret = fmt.Sprintf("%s", gameInfoMap)
	return
}

func getChatInfo(args []interface{}) (ret interface{}, err error) {
	ret = fmt.Sprintf("%s", chatInfoMap)
	return
}

func getBestGameInfo(args []interface{}) (ret interface{}, err error) {
	res, err := GetBestGameInfo([]interface{}{"zh-en", "123132"})
	ret = fmt.Sprintf("%s", res)
	return
}
