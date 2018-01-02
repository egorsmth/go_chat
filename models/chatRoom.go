package models

import (
	"../shared"
)

type ChatRoom struct {
	Id int
}

type UserToChatRoom struct {
	UserId     int
	ChatRoomId int
}

func GetChatRooms(user *User) (*[]UserToChatRoom, error) {
	rows, err := shared.Db.Query("select * from user_to_chat_room where user_id=$1", user.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rooms := []UserToChatRoom{}
	for rows.Next() {
		utcr := UserToChatRoom{}
		err := rows.Scan(utcr.UserId, utcr.ChatRoomId)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, utcr)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return &rooms, nil
}
