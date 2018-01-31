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
	Status     *string    `json:"status"`
	Date       *time.Time `json:"created,time"`
}

func (msg Message) SaveMessage() (*Message, error) {
	tx, err := shared.Db.Begin()
	if err != nil {
		log.Println("Can't start transaction while saving message", err)
		return nil, err
	}
	var lastMessageID int
	err = shared.Db.QueryRow("insert into user_profile_message (user_id, chat_room_id, message, status, created_at) values ($1, $2, $3, $4, $5) RETURNING id", msg.AuthorID, msg.ChatRoomID, msg.Message, msg.Status, msg.Date).Scan(&lastMessageID)
	if err != nil {
		tx.Rollback()
		log.Println("error while saving message", err)
		return nil, err
	}

	if err != nil {
		tx.Rollback()
		log.Println("error getting id of messgae while saving message", err)
		return nil, err
	}
	_, err = shared.Db.Exec("update user_profile_chatroom set last_message_id=$1 where id=$2", lastMessageID, msg.ChatRoomID)
	if err != nil {
		tx.Rollback()
		log.Println("error updating chat room last message while saving message", err)
		return nil, err
	}
	tx.Commit()
	msg.ID = &lastMessageID

	return &msg, nil
}

func GetMessagesByChatRoomID(ID int) (*[]Message, error) {
	rows, err := shared.Db.Query("select username, user_id, chat_room_id, message, status, created_at, avatar from user_profile_message "+
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
		err := rows.Scan(&usr.Username, &msg.AuthorID, &msg.ChatRoomID, &msg.Message, &msg.Status, &msg.Date, &usr.Avatar)
		if err != nil {
			log.Println("error while scan message for chat ", ID)
			return nil, err
		}
		usr.ID = msg.AuthorID
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
	buf := bytes.NewBufferString("select user_profile_message.id, username, user_profile_message.user_id, chat_room_id, message, status, created_at, avatar " +
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
	buf.WriteString(" order by created_at")
	rows, err := shared.Db.Query(buf.String())
	if err != nil {
		return nil, err
	}

	messages := make(map[string]*[]Message)
	for rows.Next() {
		usr := User{}
		msg := Message{}
		err = rows.Scan(&msg.ID, &usr.Username, &msg.AuthorID, &msg.ChatRoomID, &msg.Message, &msg.Status, &msg.Date, &usr.Avatar)
		if err != nil {
			log.Println("error while scan messages for chats")
			return nil, err
		}
		usr.ID = msg.AuthorID
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

func ReadMessages(ids []int) error {
	buf := bytes.NewBufferString("update user_profile_message set status='read' where id in (")
	for i, v := range ids {
		if i > 0 {
			buf.WriteString(",")
		}
		buf.WriteString(strconv.Itoa(v))
	}
	buf.WriteString(")")
	_, err := shared.Db.Exec(buf.String())
	return err
}
