package db

import (
	"database/sql"
	"time"
	"encoding/base64"
	"encoding/json"
	"strings"
	_ "github.com/lib/pq"
)

var db *sql.DB

type Session struct {
	session_key *string
    session_data *string
    expire_data *time.Time
}

type UserSess struct {
	Auth_user_hash *string `json:"_auth_user_hash"`
	Auth_user_id *string `json:"_auth_user_id"`
	Auth_user_backend *string `json:"_auth_user_backend"`
}


func Init(info string) error {
	var err error
	db, err = sql.Open("postgres", info)
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
        return err
	}
	return nil
}