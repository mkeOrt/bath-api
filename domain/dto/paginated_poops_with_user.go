package dto

type PaginatedPoopsWithUser struct {
	Count       int64                 `json:"count"`
	Page        int                   `json:"page"`
	PagesAmount int                   `json:"pages_amount"`
	PageSize    int                   `json:"page_size"`
	Poops       []PoopCreatedWithUser `json:"poops"`
}
