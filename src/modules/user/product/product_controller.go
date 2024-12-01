package product

import (
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers/response"
	"github.com/eCanteens/backend-ecanteens/src/helpers/validation"
	"github.com/gin-gonic/gin"
)

type Controller interface {
	checkFeedback(ctx *gin.Context)
	addFeedback(ctx *gin.Context)
	getFavorite(ctx *gin.Context)
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

func (c *controller) checkFeedback(ctx *gin.Context) {
	id := ctx.Param("id")

	user, _ := ctx.Get("user")
	_user := user.(models.User)

	feedback, err := c.service.checkFeedback(*_user.Id, id)

	if err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 200, gin.H{"is_liked": feedback})
}

func (c *controller) addFeedback(ctx *gin.Context) {
	var body feedbackScheme
	id := ctx.Param("id")

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	user, _ := ctx.Get("user")
	_user := user.(models.User)

	if err := c.service.addFeedback(&body, *_user.Id, id); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	msg := "Produk berhasil di"
	if body.IsLiked == nil {
		msg += "unlike/undislike"
	} else if *body.IsLiked {
		msg += "like"
	} else if !*body.IsLiked {
		msg += "dislike"
	}

	response.Success(ctx, 200, gin.H{"message": msg})
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

func (c *controller) addFavorite(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	id := ctx.Param("id")

	if err := c.service.addFavorite(*user.(models.User).Id, id); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 200, gin.H{"message": "Produk berhasil ditambahkan ke favorit"})
}

func (c *controller) removeFavorite(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	id := ctx.Param("id")

	if err := c.service.removeFavorite(*user.(models.User).Id, id); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 200, gin.H{"message": "Produk berhasil dihapus dari favorit"})
}
