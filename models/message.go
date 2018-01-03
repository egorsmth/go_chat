package models

import "time"

type Message struct {
	Author  *models.User `json:"user"`
	Message string       `json:"message"`
	Date    time.Time    `json:"created"`
}
