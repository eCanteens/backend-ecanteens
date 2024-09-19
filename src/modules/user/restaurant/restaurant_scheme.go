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

type productCategoryDTO struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

// CategoryProductsDTO represents the paginated products for a specific category.
type CategoryProductsDTO struct {
	Category *productCategoryDTO `json:"category"`
	*pagination.Pagination[models.Product]
}

// MainResponseDTO represents the main response containing categories and data.
type GetProductsResponse struct {
	Meta struct {
		Categories []*productCategoryDTO `json:"categories"`
	} `json:"meta"`
	Data []*CategoryProductsDTO `json:"data"`
}
