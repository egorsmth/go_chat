package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/egorsmth/go_chat/middleware"
	"github.com/egorsmth/go_chat/models"
	"github.com/egorsmth/go_chat/shared"
	"github.com/gorilla/websocket"
)

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	user, err := middleware.GetUserFromSession(r)
	if err != nil {
		http.Redirect(w, r, "/", 301)
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

	if chatRoom, ok := shared.WsChatRooms[chatRoomID]; ok {
		chatRoom[ws] = true
	} else {
		shared.WsChatRooms[chatRoomID] = make(map[*websocket.Conn]bool)
		shared.WsChatRooms[chatRoomID][ws] = true
	}

	broadcast := make(chan models.Message)
	go handleMessages(shared.WsChatRooms[chatRoomID], broadcast, user)

	for {
		msg := models.Message{}
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Println("err while READing json from ws", err)
			delete(shared.WsChatRooms[chatRoomID], ws)
			break
		}
		broadcast <- msg
	}
	log.Printf("user id: %v disconected from chat room %v", *user.ID, chatRoomID)
}

type ResponseMessage struct {
	Status string `json:"status"`
	Type   string `json:"type"`
	Data   string `json:"data"`
}

func handleMessages(clients map[*websocket.Conn]bool, broadcast chan models.Message, user *models.User) {
	for {
		msg := <-broadcast
		log.Printf("user id: %v send messege %v", *user.ID, msg)
		saved, err := msg.SaveMessage()
		if err != nil {
			log.Println("err while saving message", err)
			break
		}
		saved.User = *user
		savedJSON, err := json.Marshal(saved)
		rsp := ResponseMessage{"success", "MESSEGE_RECIEVED", string(savedJSON)}
		for client := range clients {
			err := client.WriteJSON(rsp)
			if err != nil {
				log.Println("err while WRITEing json to ws", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
