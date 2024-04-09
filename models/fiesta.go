package models

import "time"

type Fiesta struct {
	Title     string    `json:"title"`
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Images    []string  `json:"images"`
	Post_date time.Time `json:"post_date"`
	CanEdit   bool      `json:"can_edit"`
}

type FiestaDetails struct {
	Username string
	FiestaID string
}

type SmallFiesta struct {
	Title         string `json:"title"`
	ID            string `json:"id"`
	Username      string `json:"username"`
	CoverImageURL string `json:"coverimage"`
}
