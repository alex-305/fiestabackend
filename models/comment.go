package models

import "time"

type Comment struct {
	ID        string    `json:"id"`
	Content   string    `json:"content"`
	Fiestaid  string    `json:"fiestaid"`
	Username  string    `json:"username"`
	Post_date time.Time `json:"post_date"`
}
