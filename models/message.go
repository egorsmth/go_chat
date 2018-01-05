package models

import (
	"log"
	"time"

	"github.com/egorsmth/go_chat/shared"
)

type Message struct {
	User       User
	AuthorID   int       `json:"user_id,string"`
	ChatRoomID int       `json:"chat_room_id,string"`
	Message    string    `json:"message"`
	Date       time.Time `json:"created,time"`
}

func (msg Message) SaveMessage() error {
	log.Println("msg", msg)
	_, err := shared.Db.Exec("insert into user_profile_message (user_id, chat_room_id, message, created_at) values ($1, $2, $3, $4)", msg.AuthorID, msg.ChatRoomID, msg.Message, msg.Date)
	if err != nil {
		log.Println("error while saving message", err)
		return err
	}
	return nil
}

func GetMessagesByChatRoomID(ID int) (*[]Message, error) {
	rows, err := shared.Db.Query("select username, user_id, chat_room_id, message, created_at from user_profile_message join auth_user on auth_user.id=user_id where chat_room_id=$1", ID)
	if err != nil {
		log.Println("error while get messages for chat ", ID)
		return nil, err
	}
	messages := []Message{}
	for rows.Next() {
		usr := User{}
		msg := Message{}
		err := rows.Scan(&usr.Username, &msg.AuthorID, &msg.ChatRoomID, &msg.Message, &msg.Date)
		msg.User = usr
		if err != nil {
			log.Println("error while scan message for chat ", ID)
			return nil, err
		}
		messages = append(messages, msg)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return &messages, nil
}
