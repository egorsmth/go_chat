package db

import (
	"database/sql"
	"fmt"
	"time"
	"encoding/base64"
	"encoding/json"
	"strings"
	_ "github.com/lib/pq"
)

func dec(str string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

type UserSess struct {
	Auth_user_hash *string `json:"_auth_user_hash"`
	Auth_user_id *string `json:"_auth_user_id"`
	Auth_user_backend *string `json:"_auth_user_backend"`
}

type User struct {
	Id *int
	Username *string
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

func GetUser(sessionKey string) (*User, error) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		"root", "root", "social_net")
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	row := db.QueryRow("select * from django_session where session_key=$1", sessionKey)
	var session_key string
    var session_data string
    var expire_data time.Time
	err = row.Scan(&session_key, &session_data, &expire_data)
	if err != nil {
		return nil, err
	}

	user := UserSess{}
	err = genUser(&user, session_data)
	if err != nil {
		return nil, err
	}

	row = db.QueryRow("select id, username from auth_user where id=$1", *user.Auth_user_id)
	userDb := User{}
	err = row.Scan(&userDb.Id, &userDb.Username)
	if err != nil {
		return nil, err
	}
	return &userDb, nil
}
