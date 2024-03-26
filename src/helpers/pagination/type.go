package pagination

import "gorm.io/gorm"

type Meta struct {
	CurrentPage int   `json:"current_page"`
	PerPage     int   `json:"per_page"`
	Total       int64 `json:"total"`
	LastPage    int   `json:"last_page"`
}

type Pagination struct {
	Meta *Meta       `json:"meta"`
	Data interface{} `json:"data"`
}

type Params struct {
	Query     *gorm.DB
	Page      interface{}
	Limit     interface{}
	Order     string
	Direction string
}
