package shared

import (
	"github.com/gorilla/websocket"
)

var WsChatRooms = make(map[int]map[*websocket.Conn]bool)
var WsUpgrader = websocket.Upgrader{}
