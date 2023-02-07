package dto

import "time"

type PoopCreatedWithUser struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Latitude    string    `json:"latitude"`
	Longitude   string    `json:"longitude"`
	LastUpdate  time.Time `json:"last_update"`
	User        User      `json:"user"`
}
