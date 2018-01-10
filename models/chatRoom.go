package models

import (
	"bytes"
	"strconv"

	"github.com/egorsmth/go_chat/shared"
)

type ChatRoom struct {
	ID            int     `json:"id,string"`
	LastMessage   Message `json:"last_message,omitempty"`
	LastMessageID int     `json:"last_message_id,omitempty"`
	Status        string  `json:"status,string"`
}

func GetChatRooms(user *User) (*[]ChatRoom, error) {
	rows, err := shared.Db.Query("select chat_room_id from user_profile_usertopairchatroom where user_id=$1", user.ID)
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
	if len(roomsIds) == 0 {
		var cr []ChatRoom
		return &cr, nil
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
		err = rows.Scan(&chr.ID, &chr.LastMessageID, &chr.Status)
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
