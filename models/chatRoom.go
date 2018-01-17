package models

import (
	"bytes"
	"strconv"

	"github.com/egorsmth/go_chat/shared"
)

type ChatRoom struct {
	ID            int      `json:"id,string"`
	LastMessage   *Message `json:"lastMessage,omitempty"`
	LastMessageID int      `json:"lastMessageId,omitempty"`
	Status        string   `json:"status,string"`
}

func GetChatRooms(user *User) (*[]ChatRoom, *[]int, error) {
	rows, err := shared.Db.Query("select chat_room_id from user_profile_usertopairchatroom where user_id=$1", user.ID)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	roomsIds := []int{}
	for rows.Next() {
		var ChatRoomID int
		err := rows.Scan(&ChatRoomID)
		// log.Print("chat room id", ChatRoomID)
		if err != nil {
			return nil, nil, err
		}
		roomsIds = append(roomsIds, ChatRoomID)
	}
	if len(roomsIds) == 0 {
		var cr []ChatRoom
		var chatRoomsIds []int
		return &cr, &chatRoomsIds, nil
	}

	rooms, err := selectRooms(roomsIds)
	if err != nil {
		return nil, nil, err
	}

	err = rows.Err()
	if err != nil {
		return nil, nil, err
	}
	return rooms, &roomsIds, nil
}

func selectRooms(roomsIds []int) (*[]ChatRoom, error) {
	buf := bytes.NewBufferString("select id, last_message_id, type from user_profile_chatroom " +
		"where id in (")
	for i, v := range roomsIds {
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

	rooms := []ChatRoom{}
	for rows.Next() {
		chr := ChatRoom{}
		err = rows.Scan(&chr.ID, &chr.LastMessageID, &chr.Status)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, chr)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return &rooms, nil
}
