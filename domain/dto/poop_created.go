package dto

type PoopCreated struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
}
