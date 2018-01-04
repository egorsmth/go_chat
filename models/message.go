package models

import "time"

type Message struct {
	AuthorID int       `json:"user_id,string"`
	Message  string    `json:"message"`
	Date     time.Time `json:"created,time"`
}
