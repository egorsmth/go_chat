package tests

import (
	"github.com/egorsmth/go_chat/shared"
)

func genUser() error {
	_, err := shared.Db.Exec("insert into auth_user " +
		"(id, username, password, is_superuser, email, first_name, last_name, is_staff, is_active, date_joined) values " +
		"(1, 'mattspr', 'password', false, 'matt@spr.com', '', '', false, true, '2018-01-22'), " +
		"(2, 'Zack', 'pswwtf', false, 'zax@mail.com', '', '', false, true, '2018-01-25')")
	return err
}

func genPairChatroom() error {
	_, err := shared.Db.Exec("insert into user_profile_usertopairchatroom " +
		"(user_id, pair_id, chat_room_id) values " +
		"(1, 2, 1)")
	return err
}

func genChatRooms() error {
	_, err := shared.Db.Exec("insert into user_profile_chatroom (id, last_message_id, type) values (1, 1, 'pair'),(2, 2, 'pair')")
	return err
}

func genMessages() error {
	_, err := shared.Db.Exec("insert into user_profile_message (id, chat_room_id, message, created_at, user_id, status) values " +
		"(1, 1, 'first message', '2018-02-07 15:33', 2, 'unread'), (2, 1, 'second message', '2018-02-07 16:11', 1, 'unread')")
	return err
}

func clearDb() error {
	_, err := shared.Db.Exec("DELETE FROM user_profile_message")
	if err != nil {
		return err
	}
	_, err = shared.Db.Exec("DELETE FROM user_profile_usertopairchatroom")
	if err != nil {
		return err
	}
	_, err = shared.Db.Exec("DELETE FROM auth_user")
	if err != nil {
		return err
	}

	_, err = shared.Db.Exec("DELETE FROM user_profile_chatroom")
	if err != nil {
		return err
	}
	return nil
}
