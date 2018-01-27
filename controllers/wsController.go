package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/egorsmth/go_chat/middleware"
	"github.com/egorsmth/go_chat/models"
	"github.com/egorsmth/go_chat/shared"
	"github.com/gorilla/websocket"
)

type ResponseMessage struct {
	Status string `json:"status"`
	Type   string `json:"type"`
	Data   string `json:"data"`
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	user, err := middleware.GetUserFromSession(r)
	if err != nil {
		http.Redirect(w, r, "/", 301)
		return
	}

	params := r.URL.Query()
	ID := params.Get("id")
	chatRoomID, err := strconv.Atoi(ID)
	if err != nil {
		log.Fatal("Unable to get id frow ws url:", err)
	}

	ws, err := shared.WsUpgrader.Upgrade(w, r, nil)
	log.Printf("user id: %v connected by ws to room %v", *user.ID, chatRoomID)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	var chatRoom shared.WsChatRoom
	if chatRoom, ok := shared.WsChatRooms[chatRoomID]; ok {
		chatRoom.Rooms[ws] = true
	} else {
		rooms := shared.WsChatRoom{&sync.Mutex{}, make(map[*websocket.Conn]bool)}
		shared.WsChatRooms[chatRoomID] = rooms
		shared.WsChatRooms[chatRoomID].Rooms[ws] = true
		chatRoom = shared.WsChatRooms[chatRoomID]
	}

	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("err read message", err)
			delete(chatRoom.Rooms, ws)
			break
		}
		msg := models.Message{}
		err = json.Unmarshal(message, &msg)
		if err != nil {
			log.Println("err while READing json from ws", err)
			delete(chatRoom.Rooms, ws)
			break
		}
		saved, err := msg.SaveMessage()
		if err != nil {
			log.Println("err while saving message", err)
			break
		}
		saved.User = *user
		savedJSON, err := json.Marshal(saved)
		rsp := ResponseMessage{"success", "MESSEGE_RECIEVED", string(savedJSON)}
		chatRoom.Mu.Lock()
		for conn := range chatRoom.Rooms {
			err := conn.WriteJSON(rsp)
			if err != nil {
				log.Println("err while WRITEing json to ws", err)
				conn.Close()
				delete(chatRoom.Rooms, conn)
				break
			}
		}
		chatRoom.Mu.Unlock()
	}

	log.Printf("user id: %v disconected from chat room %v", *user.ID, chatRoomID)
}
