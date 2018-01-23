package middleware

import (
	"log"
	"net/http"

	"github.com/egorsmth/go_chat/models"
)

func GetUserFromSession(r *http.Request) (*models.User, error) {
	sid, err := r.Cookie("sessionid")
	if err != nil {
		log.Println("sessionid not found in cookies", err)
		return nil, err
	}
	user, err := models.GetUserFromSession(sid.Value)
	if err != nil {
		log.Println("Cant get user from session:", err)
		return nil, err
	}
	return user, nil
}
