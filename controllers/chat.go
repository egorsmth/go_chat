package controllers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/egorsmth/go_chat/models"
)

type appData struct {
	ChatRooms *[]models.ChatRoom            `json:"chatRooms,omitonempty"`
	Messages  *map[string]*[]models.Message `json:"messages"`
}

type chatResponse struct {
	User    string
	AppData string
}

func Chat(w http.ResponseWriter, r *http.Request) {
	sid, err := r.Cookie("sessionid")
	if err != nil {
		log.Println("sessionid not found in cookies", err)
		http.Redirect(w, r, "/", 301)
	}
	user, err := models.GetUserFromSession(sid.Value)
	if err != nil {
		log.Println("Cant get user from session:", err)
		http.Redirect(w, r, "/", 301)
	}

	t := template.New("chat")                           // Create a template.
	t = template.Must(t.ParseFiles("public/chat.html")) // Parse template file.
	if err != nil {
		log.Println("parse file err:", err)
	}

	cr := chatResponse{}
	appData := appData{}

	chatRooms, chatRoomsids, err := models.GetChatRooms(user)
	if err != nil {
		log.Println("err while getting initial chat rooms:", err)
	}

	messages, err := models.GetMessages(*chatRoomsids)
	if err != nil {
		log.Println("err while getting initial messages:", err)
	}

	chatRooms = linkChatRoomsWithLastMessage(*chatRooms, *messages)
	appData.ChatRooms = chatRooms
	appData.Messages = messages

	jsonAppData, err := json.Marshal(appData)
	if err != nil {
		log.Println("err while marshal initial app data:", err)
	}
	cr.AppData = string(jsonAppData)

	jsonUser, err := json.Marshal(user)
	if err != nil {
		log.Println("err while marshal user:", err)
	}
	cr.User = string(jsonUser)

	err = t.Execute(w, cr)
	if err != nil {
		log.Println("template Execute err:", err)
	}
}

func linkChatRoomsWithLastMessage(chatRooms []models.ChatRoom, messages map[string]*[]models.Message) *[]models.ChatRoom {
	newChatRooms := []models.ChatRoom{}
	for _, chatRoom := range chatRooms {
		chatRoomID := strconv.Itoa(chatRoom.ID)
		chatRoomMessages := messages[chatRoomID]
		if chatRoomMessages != nil {
			msg := (*chatRoomMessages)[len(*chatRoomMessages)-1]
			chatRoom.LastMessage = &msg
		}
		newChatRooms = append(newChatRooms, chatRoom)
	}
	return &newChatRooms
}
