package internal

var (
	gameInfoMap    = map[string]map[string]*GameInfo{}
	gameRegion     = map[string]string{}
	chatInfoMap    = map[string]map[string]*ChatInfo{}
	chatRegion     = map[string]string{}
	accountGameMap = map[string]*GameInfo{}
)

type GameInfo struct {
	serverName     string
	clientCount    int
	maxClientCount int
	clientAddr     string
}

type ChatInfo struct {
	serverName  string
	clientCount int
	clientAddr  string
}
