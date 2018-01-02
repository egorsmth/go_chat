package db

import (
	"encoding/base64"
	"encoding/json"
	"strings"
)

type User struct {
	Id       *int
	Username *string
}

func GetUserById(id string) (*User, error) {
	row := db.QueryRow("select id, username from auth_user where id=$1", id)
	user := User{}
	err := row.Scan(&user.Id, &user.Username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserFromSession(skey string) (*User, error) {
	row := db.QueryRow("select * from django_session where session_key=$1", skey)
	sess := Session{}
	err := row.Scan(&sess.session_key, &sess.session_data, &sess.expire_data)
	if err != nil {
		return nil, err
	}

	user := UserSess{}
	err = genUser(&user, *sess.session_data)
	if err != nil {
		return nil, err
	}

	userDb, err := GetUserById(*user.Auth_user_id)
	if err != nil {
		return nil, err
	}
	return userDb, nil
}

func dec(str string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func genUser(user *UserSess, session_data string) error {
	decoded, err := dec(session_data)
	if err != nil {
		return err
	}
	data := strings.SplitN(decoded, ":", 2)
	err = json.Unmarshal([]byte(data[1]), &user)
	return err
}
