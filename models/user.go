package models

import (
	"encoding/base64"
	"encoding/json"
	"strings"
	"time"

	"github.com/egorsmth/go_chat/shared"
)

type session struct {
	session_key  string
	session_data string
	expire_data  time.Time
}

type userSess struct {
	Auth_user_hash    string `json:"_auth_user_hash"`
	Auth_user_id      string `json:"_auth_user_id"`
	Auth_user_backend string `json:"_auth_user_backend"`
}

// User is auth_user table model
type User struct {
	ID       int    `json:"id,string"`
	Username string `json:"username,string"`
	Avatar   string `json:"avatar,string"`
}

// GetUserByID get user by id
func GetUserByID(id string) (*User, error) {
	row := shared.Db.QueryRow("select id, username from auth_user where id=$1", id)
	user := User{}
	err := row.Scan(&user.ID, &user.Username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserFromSession(skey string) (*User, error) {
	row := shared.Db.QueryRow(
		"select * from django_session where session_key=$1",
		skey)
	sess := session{}
	err := row.Scan(&sess.session_key, &sess.session_data, &sess.expire_data)
	if err != nil {
		return nil, err
	}

	user := userSess{}
	err = genUser(&user, sess.session_data)
	if err != nil {
		return nil, err
	}

	userDb, err := GetUserByID(user.Auth_user_id)
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

func genUser(user *userSess, session_data string) error {
	decoded, err := dec(session_data)
	if err != nil {
		return err
	}
	data := strings.SplitN(decoded, ":", 2)
	err = json.Unmarshal([]byte(data[1]), &user)
	return err
}
