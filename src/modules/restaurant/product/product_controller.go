package product

import (
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers/response"
	"github.com/eCanteens/backend-ecanteens/src/helpers/validation"
	"github.com/gin-gonic/gin"
)

type Controller interface {
	createProduct(ctx *gin.Context)
	getAllProducts(ctx *gin.Context)
	updateProduct(ctx *gin.Context)
	deleteProduct(ctx *gin.Context)
}

type controller struct {
	service Service
}

func NewController(service Service) Controller {
	return &controller{
		service: service,
	}
}

func (c *controller) createProduct(ctx *gin.Context) {
	var body createProduct

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	user, _ := ctx.Get("user")
	_user := user.(models.User)

	if err := c.service.createProduct(&_user, &body); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 201, gin.H{"message": "Menu berhasil ditambahkan"})
}

func (c *controller) getAllProducts(ctx *gin.Context) {
	var query productQs

	ctx.ShouldBindQuery(&query)

	user, _ := ctx.Get("user")
	_user := user.(models.User)

	data, err := c.service.getAllProducts(&query, &_user)

	if err != nil {
		response.ServiceError(ctx, err)
		return
	}

	ctx.JSON(200, data)
}

func (c *controller) updateProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	var body updateProduct

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	user, _ := ctx.Get("user")
	_user := user.(models.User)

	if err := c.service.updateProduct(&_user, &body, id); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 200, gin.H{"message": "Menu berhasil diupdate"})
}

func (c *controller) deleteProduct(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	_user := user.(models.User)
	id := ctx.Param("id")

	if err := c.service.deleteProduct(&_user, id); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 200, gin.H{"message": "Menu berhasil dihapus"})
}
