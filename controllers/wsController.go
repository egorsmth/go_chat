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

type RequestMessage struct {
	Action string          `json:"action"`
	Data   json.RawMessage `json:"data"`
}

type ReadRecievedData struct {
	MessageIds []int `json:"messageIds"`
	RoomId     int   `json:"roomId,string"`
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

	if _, ok := shared.WsChatRooms[chatRoomID]; ok {
		shared.WsChatRooms[chatRoomID].Rooms[ws] = true
	} else {
		rooms := shared.WsChatRoom{&sync.Mutex{}, make(map[*websocket.Conn]bool)}
		rooms.Mu = &sync.Mutex{}
		shared.WsChatRooms[chatRoomID] = rooms
		shared.WsChatRooms[chatRoomID].Rooms[ws] = true
	}
	for {
		_, req, err := ws.ReadMessage()
		if err != nil {
			log.Println("err read message", err)
			delete(shared.WsChatRooms[chatRoomID].Rooms, ws)
			break
		}

		requestMessage := RequestMessage{}
		err = json.Unmarshal(req, &requestMessage)
		if err != nil {
			log.Println("err while READing json from ws", err)
			delete(shared.WsChatRooms[chatRoomID].Rooms, ws)
			break
		}
		err = dispatchAction(&requestMessage, chatRoomID, user, ws)
		if err != nil {
			break
		}
	}

	log.Printf("user id: %v disconected from chat room %v", *user.ID, chatRoomID)
}

func dispatchAction(request *RequestMessage, chatRoomID int, user *models.User, ws *websocket.Conn) error {
	if request.Action == "send" {
		msg := models.Message{}
		err := json.Unmarshal(request.Data, &msg)
		if err != nil {
			log.Println("err while unmarshal send message from ws", err)
			return err
		}
		saved, err := msg.SaveMessage()
		if err != nil {
			log.Println("err while saving message", err)
			return err
		}
		saved.User = *user
		savedJSON, err := json.Marshal(saved)
		rsp := &ResponseMessage{"success", "MESSEGE_RECIEVED", string(savedJSON)}
		sendResponse(chatRoomID, rsp)
	} else if request.Action == "read" {
		recieved := ReadRecievedData{}
		err := json.Unmarshal(request.Data, &recieved)
		if err != nil {
			log.Println("err while unmarshal read recieved data from ws", err)
			return err
		}

		err = models.ReadMessages(recieved.MessageIds)
		if err != nil {
			log.Println("err while read messages", err)
			return err
		}
		rsp := &ResponseMessage{"success", "MESSEGE_READED", string(request.Data)}
		sendResponse(chatRoomID, rsp)
	}
	return nil
}

func sendResponse(chatRoomID int, rsp *ResponseMessage) {
	shared.WsChatRooms[chatRoomID].Mu.Lock()
	defer shared.WsChatRooms[chatRoomID].Mu.Unlock()
	for conn := range shared.WsChatRooms[chatRoomID].Rooms {
		err := conn.WriteJSON(rsp)
		if err != nil {
			log.Println("err while WRITEing json to ws", err)
			conn.Close()
			delete(shared.WsChatRooms[chatRoomID].Rooms, conn)
		}
	}
}
