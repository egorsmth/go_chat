package models

import (
	"bytes"
	"log"
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
	buf := bytes.NewBufferString("select chatroom.id as chat_room_id, last_message_id, type, " + // user_profile_chatroom rows
		"user_profile_message.id as message_id, user_profile_message.user_id, chat_room_id, message, created_at, " + // user_profile_message rows
		"auth_user.id as uid, username, avatar " + // auth_user and user_profile_profile rows
		"from user_profile_chatroom as chatroom" +
		"left join user_profile_message on last_message_id=user_profile_message.id " +
		"left join auth_user on user_profile_message.user_id=auth_user.id " +
		"left join user_profile_profile on user_profile_profile.user_id=user_profile_message.user_id " +
		"where chatroom.id in (")
	for i, v := range roomsIds {
		if i > 0 {
			buf.WriteString(",")
		}
		buf.WriteString(strconv.Itoa(v))
	}
	buf.WriteString(")")
	rows, err := shared.Db.Query(buf.String())
	if err != nil {
		log.Println("err while selecting chat rooms")
		return nil, err
	}

	rooms := []ChatRoom{}
	for rows.Next() {
		chr := ChatRoom{}
		msg := Message{}
		usr := User{}

		err = rows.Scan(&chr.ID, &chr.LastMessageID, &chr.Status, &msg.ID, &msg.AuthorID, &msg.ChatRoomID, &msg.Message, &msg.Date, &usr.ID, &usr.Username, &usr.Avatar)
		if err != nil {
			log.Println("err while scaning chat rooms")
			return nil, err
		}
		if msg.ID != nil {
			msg.User = &usr
			chr.LastMessage = &msg
		}

		rooms = append(rooms, chr)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return &rooms, nil
}
