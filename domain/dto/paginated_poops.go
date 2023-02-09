package dto

type PaginatedPoops struct {
	Count       int64         `json:"count"`
	Page        int           `json:"page"`
	PagesAmount int           `json:"pages_amount"`
	PageSize    int           `json:"page_size"`
	Poops       []PoopCreated `json:"poops"`
}
