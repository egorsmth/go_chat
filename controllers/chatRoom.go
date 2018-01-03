package controllers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	"../models"
	"../shared"
)

type ChatRoomResponse struct {
	User     models.User
	Messages []models.Message
}

func ChatRoom(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	chatRoomID := params.Get("id")

	t := template.New("chat_room")                           // Create a template.
	t = template.Must(t.ParseFiles("public/chat_room.html")) // Parse template file.
	if err != nil {
		log.Println("parse file err:", err)
	}

	cr = ChatRoomResponse{}
	err = t.Execute(w, cr)
	if err != nil {
		log.Println("template Execute err:", err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	chatRoomID := params.Get("id")

	ws, err := shared.WsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	if val, ok := shared.WsChatRooms[chatRoomID]; ok {
		val = append(val, ws)
	} else {
		shared.WsChatRooms[chatRoomID] = []*websocket.Conn{ws}
	}

	broadcast := make(chan Message)
	go handleMessages(shared.WsChatRooms[chatRoomID], broadcast)

	for {
		msg := Message{}
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Fatal(err)
			delete(shared.WsChatRooms[chatRoomID], ws)
			break
		}
		broadcast <- msg
	}
}

func handleMessages(*clients, broadcast) {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Fatal(err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
