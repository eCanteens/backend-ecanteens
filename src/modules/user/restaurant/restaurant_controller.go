package restaurant

import (
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers/response"
	"github.com/gin-gonic/gin"
)

type Controller interface {
	getFavorite(ctx *gin.Context)
	getAll(ctx *gin.Context)
	getReviews(ctx *gin.Context)
	getDetail(ctx *gin.Context)
	getRestosProducts(ctx *gin.Context)
	addFavorite(ctx *gin.Context)
	removeFavorite(ctx *gin.Context)
}

type controller struct {
	service Service
}

func NewController(service Service) Controller {
	return &controller{
		service: service,
	}
}

func (c *controller) getFavorite(ctx *gin.Context) {
	var query paginationQS
	user, _ := ctx.Get("user")

	ctx.ShouldBindQuery(&query)

	data, err := c.service.getFavorite(*user.(models.User).Id, &query)

	if err != nil {
		response.ServiceError(ctx, err)
		return
	}

	ctx.JSON(200, data)
}

func (c *controller) getAll(ctx *gin.Context) {
	var query paginationQS

	ctx.ShouldBindQuery(&query)

	data, err := c.service.getAll(&query)

	if err != nil {
		response.ServiceError(ctx, err)
		return
	}

	ctx.JSON(200, data)
}

func (c *controller) getReviews(ctx *gin.Context) {
	id := ctx.Param("id")
	var query reviewQS

	ctx.ShouldBindQuery(&query)

	data, err := c.service.getReviews(id, &query)

	if err != nil {
		response.ServiceError(ctx, err)
		return
	}

	ctx.JSON(200, data)
}

func (c *controller) getDetail(ctx *gin.Context) {
	id := ctx.Param("id")
	user, _ := ctx.Get("user")

	data, err := c.service.getDetail(id, *user.(models.User).Id)

	if err != nil {
		response.ServiceError(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{
		"data": data,
	})
}

func (c *controller) getRestosProducts(ctx *gin.Context) {
	id := ctx.Param("id")
	var query getProductsQS

	user, _ := ctx.Get("user")

	ctx.ShouldBindQuery(&query)

	data, err := c.service.getRestosProducts(id, &query, *user.(models.User).Id)

	if err != nil {
		response.ServiceError(ctx, err)
		return
	}

	ctx.JSON(200, data)
}

func (c *controller) addFavorite(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	id := ctx.Param("id")

	if err := c.service.addFavorite(*user.(models.User).Id, id); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 200, gin.H{"message": "Restoran berhasil ditambahkan ke favorit"})
}

func (c *controller) removeFavorite(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	id := ctx.Param("id")

	if err := c.service.removeFavorite(*user.(models.User).Id, id); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 200, gin.H{"message": "Restoran berhasil dihapus dari favorit"})
}
