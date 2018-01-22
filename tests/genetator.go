package tests

import (
	"github.com/egorsmth/go_chat/shared"
)

func genUser() error {
	_, err := shared.Db.Exec("insert into auth_user " +
		"(id, username, password, is_superuser, email, first_name, last_name, is_staff, is_active, date_joined) values " +
		"(1, 'mattspr', 'password', false, 'matt@spr.com', '', '', false, true, '2018-01-22')")
	return err
}

func genChatRooms() error {
	_, err := shared.Db.Exec("insert into user_profile_chatroom (id, last_message_id, type) values (1, 1, 'pair'),(2, 2, 'pair')")
	return err
}

func clearDb() {
	_, _ = shared.Db.Exec("DELETE FROM auth_user")
	_, _ = shared.Db.Exec("DELETE FROM user_profile_chatroom")
}
