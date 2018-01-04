package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/websocket"

	"github.com/egorsmth/go_chat/models"
	"github.com/egorsmth/go_chat/shared"
)

type ChatRoomResponse struct {
	User     models.User
	Messages []models.Message
	RoomID   int
}

func ChatRoom(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	ID := params.Get("id")
	chatRoomID, err := strconv.Atoi(ID)
	if err != nil {
		log.Fatal("Unable to get id frow chat room url:", err)
	}

	sid, err := r.Cookie("sessionid")
	if err != nil {
		log.Println(err)
		// todo redirect to login if no session
	}
	user, err := models.GetUserFromSession(sid.Value)
	if err != nil {
		log.Println("In ChatRoom GetUserFromSession error:", err)
		// todo redirect to login if no session
	}

	t := template.New("chat_room")                           // Create a template.
	t = template.Must(t.ParseFiles("public/chat_room.html")) // Parse template file.

	cr := ChatRoomResponse{*user, nil, chatRoomID}
	err = t.Execute(w, cr)
	if err != nil {
		log.Println("template Execute err:", err)
	}
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	ID := params.Get("id")
	chatRoomID, err := strconv.Atoi(ID)
	if err != nil {
		log.Fatal("Unable to get id frow ws url:", err)
	}

	sid, err := r.Cookie("sessionid")
	if err != nil {
		log.Println(err)
		// todo redirect to login if no session
	}
	user, err := models.GetUserFromSession(sid.Value)
	if err != nil {
		log.Println("In ws handling GetUserFromSession error:", err)
		// todo redirect to login if no session
	}

	ws, err := shared.WsUpgrader.Upgrade(w, r, nil)
	log.Printf("user id: %v connected by ws to room %v", user.ID, chatRoomID)
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
	log.Printf("user id: %v disconected from chat room %v", user.ID, chatRoomID)
}

type ResponseMessage struct {
	Username string    `json:"username"`
	Message  string    `json:"message"`
	Created  time.Time `json:"created"`
}

func handleMessages(clients map[*websocket.Conn]bool, broadcast chan models.Message, user *models.User) {
	for {
		msg := <-broadcast
		log.Printf("user id: %v send messege %v", user.ID, msg)
		rsp := ResponseMessage{user.Username, msg.Message, msg.Date}
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