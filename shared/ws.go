package shared

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type WsChatRoom struct {
	Mu    *sync.Mutex
	Rooms map[*websocket.Conn]bool
}

var WsChatRooms = make(map[int]WsChatRoom)
var WsUpgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
