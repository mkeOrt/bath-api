package model

import "gorm.io/gorm"

type Poop struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"Longitude"`
	UserID      int
	User        User
}
