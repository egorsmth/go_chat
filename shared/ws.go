package shared

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var WsChatRooms = make(map[int]map[*websocket.Conn]bool)
var WsUpgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
