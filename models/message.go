package models

import (
	"bytes"
	"log"
	"strconv"
	"time"

	"github.com/egorsmth/go_chat/shared"
)

type Message struct {
	ID         *int       `json:"id"`
	User       User       `json:"author"`
	AuthorID   *int       `json:"author_id,string"`
	ChatRoomID *int       `json:"chat_room_id,string"`
	Message    *string    `json:"text"`
	Date       *time.Time `json:"created,time"`
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
	rows, err := shared.Db.Query("select username, user_id, chat_room_id, message, created_at, avatar from user_profile_message "+
		"join auth_user on auth_user.id=user_id "+
		"left join user_profile_profile on user_profile_profile.user_id=user_id "+
		"where chat_room_id=$1 "+
		"order by created_at", ID)
	if err != nil {
		log.Println("error while get messages for chat ", ID)
		return nil, err
	}
	messages := []Message{}
	for rows.Next() {
		usr := User{}
		msg := Message{}
		err := rows.Scan(&usr.Username, &msg.AuthorID, &msg.ChatRoomID, &msg.Message, &msg.Date, &usr.Avatar)
		if err != nil {
			log.Println("error while scan message for chat ", ID)
			return nil, err
		}
		msg.User = usr
		messages = append(messages, msg)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return &messages, nil
}

func GetMessages(roomsIds *[]int) (*map[string]*[]Message, error) {
	buf := bytes.NewBufferString("select user_profile_message.id, username, user_profile_message.user_id, chat_room_id, message, created_at, avatar " +
		"from user_profile_message " +
		"join auth_user on auth_user.id=user_profile_message.user_id " +
		"left join user_profile_profile on user_profile_profile.user_id=user_profile_message.user_id " +
		"where chat_room_id in (")
	for i, v := range *roomsIds {
		if i > 0 {
			buf.WriteString(",")
		}
		buf.WriteString(strconv.Itoa(v))
	}
	buf.WriteString(")")
	rows, err := shared.Db.Query(buf.String())
	if err != nil {
		return nil, err
	}

	messages := make(map[string]*[]Message)
	for rows.Next() {
		usr := User{}
		msg := Message{}
		err = rows.Scan(&msg.ID, &usr.Username, &msg.AuthorID, &msg.ChatRoomID, &msg.Message, &msg.Date, &usr.Avatar)
		if err != nil {
			log.Println("error while scan messages for chats")
			return nil, err
		}
		log.Println(usr)
		msg.User = usr
		intID := strconv.Itoa(*msg.ChatRoomID)
		if _, exist := messages[intID]; !exist {
			messages[intID] = &[]Message{}
		}

		*messages[intID] = append(*messages[intID], msg)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return &messages, nil
}
