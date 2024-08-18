package transaction

import (
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers/response"
	"github.com/eCanteens/backend-ecanteens/src/helpers/validation"
	"github.com/gin-gonic/gin"
)

type Controller interface {
	getOrder(ctx *gin.Context)
	updateOrder(ctx *gin.Context)
}

type controller struct {
	service Service
}

func NewController(service Service) Controller {
	return &controller{
		service: service,
	}
}

func (c *controller) getOrder(ctx *gin.Context) {
	var query getOrderQS
	user, _ := ctx.Get("user")
	_user := user.(models.User)

	ctx.ShouldBindQuery(&query)

	data, err := c.service.getOrder(*_user.Restaurant.Id, &query)

	if err != nil {
		response.ServiceError(ctx, err)
		return
	}

	ctx.JSON(200, data)
}

func (c *controller) updateOrder(ctx *gin.Context) {
	var body updateOrderScheme
	id := ctx.Param("id")

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	user, _ := ctx.Get("user")
	_user := user.(models.User)

	if err := c.service.updateOrder(id, &_user, &body); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	msg := "Pesanan berhasil"

	switch body.Status {
	case "INPROGRESS":
		msg += " diterima"
	case "CANCELED":
		msg += " ditolak"
	case "READY":
		msg = "Status pesanan berhasil diperbarui"
	}

	response.Success(ctx, 200, gin.H{"message": msg})
}
