package db

import (
	"database/sql"
	"log"
	"fmt"
	"time"
	"encoding/base64"
	"encoding/json"
	"strings"
	_ "github.com/lib/pq"
)

func checAndExit(err error, db *sql.DB) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

func dec(str string) string {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println("error:", err)
		panic(err)
	}
	return string(data)
}

type User struct {
	Auth_user_hash *string `json:"_auth_user_hash"`
	Auth_user_id *string `json:"_auth_user_id"`
	Auth_user_backend *string `json:"_auth_user_backend"`
}
func GetUser(sessionKey string) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		"root", "root", "social_net")
	db, err := sql.Open("postgres", dbinfo)
	checAndExit(err, db)
	defer db.Close()
	row := db.QueryRow("select * from django_session where session_key='"+ sessionKey +"'")
	var session_key string
    var session_data string
    var expire_data time.Time
	err = row.Scan(&session_key, &session_data, &expire_data)
	checAndExit(err, db)
	//fmt.Printf("%v | %v | %v\n", session_key, session_data, expire_data)
	data := strings.SplitN(dec(session_data), ":", 2)
	u := User{}
	err = json.Unmarshal([]byte(data[1]), &u)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(*u.Auth_user_id)
}