package models

type Meta struct {
	CurrentPage int   `json:"current_page"`
	PerPage     int   `json:"per_page"`
	Total       int64 `json:"total"`
}

type Pagination struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}
