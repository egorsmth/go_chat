package models

import (
	"bytes"
	"strconv"

	"../shared"
)

type ChatRoom struct {
	ID            int
	LastMessageID int
	Type          string
	Status        string
}

func GetChatRooms(user *User) (*[]ChatRoom, error) {
	rows, err := shared.Db.Query("select chat_room_id from user_profile_friends where user_id=$1", user.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	roomsIds := []int{}
	for rows.Next() {
		var ChatRoomID int
		err := rows.Scan(&ChatRoomID)
		// log.Print("chat room id", ChatRoomID)
		if err != nil {
			return nil, err
		}
		roomsIds = append(roomsIds, ChatRoomID)
	}

	rooms := []ChatRoom{}

	err = selectRooms(roomsIds, &rooms)
	if err != nil {
		return nil, err
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return &rooms, nil
}

func selectRooms(roomsIds []int, rooms *[]ChatRoom) error {
	buf := bytes.NewBufferString("select * from user_profile_chatroom where id in (")
	for i, v := range roomsIds {
		if i > 0 {
			buf.WriteString(",")
		}
		buf.WriteString(strconv.Itoa(v))
	}
	buf.WriteString(")")
	rows, err := shared.Db.Query(buf.String())
	if err != nil {
		return err
	}

	for rows.Next() {
		chr := ChatRoom{}
		err = rows.Scan(&chr.ID, &chr.LastMessageID, &chr.Type, &chr.Status)
		if err != nil {
			return err
		}
		*rooms = append(*rooms, chr)
	}

	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}
