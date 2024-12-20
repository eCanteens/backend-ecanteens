package restaurant

import (
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers/pagination"
)

type paginationQS struct {
	Page      string `form:"page"`
	Search    string `form:"search" mod:"trim"`
	Limit     string `form:"limit"`
	Order     string `form:"order"`
	Direction string `form:"direction"`
}

type reviewQS struct {
	Filter string `form:"filter"`
}

type getProductsQS struct {
	paginationQS
	CategoryId string `form:"category_id"`
}

type categoryProductsDTO struct {
	Category *categoryDTO `json:"category"`
	*pagination.Pagination[models.Product]
}

type getProductsResponse struct {
	Meta struct {
		Categories []*categoryDTO `json:"categories"`
	} `json:"meta"`
	Data []*categoryProductsDTO `json:"data"`
}

type categoryDTO struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type categoryRestosDTO struct {
	Category *categoryDTO `json:"category"`
	*pagination.Pagination[models.Restaurant]
}
type getRestosResponse struct {
	Meta struct {
		Categories []*categoryDTO `json:"categories"`
	} `json:"meta"`
	Data []*categoryRestosDTO `json:"data"`
}
