package controllers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/egorsmth/go_chat/models"
)

type chatResponse struct {
	User      models.User
	ChatRooms []models.ChatRoom
}

func Chat(w http.ResponseWriter, r *http.Request) {
	sid, err := r.Cookie("sessionid")
	if err != nil {
		log.Println("sessionid not found in cookies", err)
		http.Redirect(w, r, "/", 301)
	}
	_, err = models.GetUserFromSession(sid.Value)
	if err != nil {
		log.Println("Cant get user from session:", err)
		http.Redirect(w, r, "/", 301)
	}

	t := template.New("chat")                           // Create a template.
	t = template.Must(t.ParseFiles("public/chat.html")) // Parse template file.
	if err != nil {
		log.Println("parse file err:", err)
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Println("template Execute err:", err)
	}
}
