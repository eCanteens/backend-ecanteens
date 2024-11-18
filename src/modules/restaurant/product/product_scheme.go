package product

import (
	"mime/multipart"

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

type createProduct struct {
	Image       *multipart.FileHeader `binding:"required" form:"image"`
	Name        string                `binding:"required" form:"name" json:"name" mod:"trim"`
	Price       uint                  `binding:"required" form:"price" json:"price" mod:"trim"`
	Stock       uint                  `binding:"required" form:"stock" json:"stock" mod:"trim"`
	Description string                `binding:"required" form:"description" json:"description" mod:"trim"`
	CategoryId  uint                  `binding:"required,numeric" form:"category_id" json:"category_id" mod:"trim"`
}

type updateProduct struct {
	Image       *multipart.FileHeader `form:"image"`
	Name        string                `binding:"required" mod:"trim" json:"name" form:"name"`
	Price       uint                  `binding:"required" mod:"trim" json:"price" form:"price"`
	Stock       uint                  `binding:"required" mod:"trim" json:"stock" form:"stock"`
	Description string                `binding:"required" mod:"trim" json:"description" form:"description"`
	CategoryId  uint                  `binding:"required,numeric" mod:"trim" json:"category_id" form:"category_id"`
}

type productQs struct {
	paginationQS
	CategoryId string `form:"category_id"`
}

type categoryProductsDTO struct {
	Category *categoryDTO `json:"category"`
	*pagination.Pagination[models.Product]
}

type categoryDTO struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type getProductsResponse struct {
	Meta struct {
		Categories []*categoryDTO `json:"categories"`
	} `json:"meta"`
	Data []*categoryProductsDTO `json:"data"`
}