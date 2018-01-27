package controllers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/egorsmth/go_chat/middleware"
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
	user, err := middleware.GetUserFromSession(r)
	if err != nil {
		http.Redirect(w, r, "/", 301)
		return
	}

	t := template.New("chat")                           // Create a template.
	t = template.Must(t.ParseFiles("public/chat.html")) // Parse template file.
	if err != nil {
		log.Println("parse file err:", err)
	}

	cr := chatResponse{}
	appData := appData{}

	chatRooms, chatRoomsID, err := models.GetChatRooms(user)
	if err != nil {
		log.Println("err while getting initial chat rooms:", err)
	}

	messages, err := models.GetMessages(chatRoomsID)
	if err != nil {
		log.Println("err while getting initial messages:", err)
	}

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
