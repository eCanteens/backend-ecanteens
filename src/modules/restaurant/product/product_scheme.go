package product

import "mime/multipart"

type createProduct struct {
	Image       *multipart.FileHeader `binding:"required" form:"image"`
	Name        string                `binding:"required" form:"name" json:"name" mod:"trim"`
	Price       uint                  `binding:"required" form:"price" json:"price" mod:"trim"`
	Stock       uint                  `binding:"required" form:"stock" json:"stock" mod:"trim"`
	Description string                `binding:"required" form:"description" json:"description" mod:"trim"`
	CategoryId  uint                  `binding:"required,numeric" form:"category_id" json:"category_id" mod:"trim"`
}
