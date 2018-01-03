package shared

import (
	"github.com/gorilla/websocket"
)

var WsChatRooms = make(map[int]*[]*websocket.Conn)
var WsUpgrader = websocket.Upgrader{}
