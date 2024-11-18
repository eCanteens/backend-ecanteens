package dashboard

import (
	"fmt"

	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers/response"
	"github.com/eCanteens/backend-ecanteens/src/helpers/validation"
	"github.com/gin-gonic/gin"
)

type Controller interface {
	toggleOpen(ctx *gin.Context)
	dashboard(ctx *gin.Context)
}

type controller struct {
	service Service
}

func NewController(service Service) Controller {
	return &controller{
		service: service,
	}
}

func (c *controller) toggleOpen(ctx *gin.Context) {
	var body toggleOpenScheme

	user, _ := ctx.Get("user")
	_user := user.(models.User)

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	if err := c.service.toggleOpen(*_user.Restaurant.Id, &body); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	msg := "ditutup"

	if body.IsOpen {
		msg = "dibuka"
	}

	response.Success(ctx, 200, gin.H{
		"message": fmt.Sprintf("Restoran berhasil di%s", msg),
	})
}

func (c *controller) dashboard(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	_user := user.(models.User)

	data, err := c.service.dashboard(*_user.Restaurant.Id);

	if err != nil {
		response.ServiceError(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{
		"data": data,
	})
}