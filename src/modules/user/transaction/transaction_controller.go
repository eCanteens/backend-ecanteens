package transaction

import (
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers/response"
	"github.com/eCanteens/backend-ecanteens/src/helpers/validation"
	"github.com/gin-gonic/gin"
)

type Controller interface {
	getCarts(ctx *gin.Context)
	getRestaurantCart(ctx *gin.Context)
	updateCart(ctx *gin.Context)
	addCart(ctx *gin.Context)
	getOrders(ctx *gin.Context)
	getOneOrder(ctx *gin.Context)
	createOrder(ctx *gin.Context)
	postReview(ctx *gin.Context)
	updateOrder(ctx *gin.Context)
	getTrxHistory(ctx *gin.Context)
}

type controller struct {
	service Service
}

func NewController(service Service) Controller {
	return &controller{
		service: service,
	}
}

func (c *controller) getCarts(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	_user := user.(models.User)

	data, err := c.service.getCarts(&_user)

	if err != nil {
		response.ServiceError(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{
		"data": data,
	})
}
func (c *controller) getRestaurantCart(ctx *gin.Context) {
	var query restaurantCartQS

	user, _ := ctx.Get("user")
	_user := user.(models.User)

	ctx.ShouldBindQuery(&query)

	cart, _ := c.service.getRestaurantCart(query.RestaurantId, &_user)

	ctx.JSON(200, gin.H{
		"data": cart,
	})

}

func (c *controller) updateCart(ctx *gin.Context) {
	id := ctx.Param("id")
	var body updateCartNoteScheme

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	user, _ := ctx.Get("user")
	_user := user.(models.User)

	if err := c.service.updateCart(id, &body, *_user.Id); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 200, gin.H{"message": "Catatan berhasil ditambahkan"})
}

func (c *controller) addCart(ctx *gin.Context) {
	var body addCartScheme

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	user, _ := ctx.Get("user")
	_user := user.(models.User)

	data, err := c.service.addCart(&_user, &body)

	if err != nil {
		response.ServiceError(ctx, err)
		return
	}

	if *body.Quantity == 0 {
		response.Success(ctx, 200, gin.H{
			"message": "Produk berhasil dihapus dari keranjang",
			"data":    data,
		})
	} else {
		response.Success(ctx, 201, gin.H{
			"message": "Produk berhasil ditambahkan ke keranjang",
			"data":    data,
		})
	}
}

func (c *controller) getOrders(ctx *gin.Context) {
	var query getOrderQS
	user, _ := ctx.Get("user")
	_user := user.(models.User)

	ctx.ShouldBindQuery(&query)

	data, err := c.service.getOrders(*_user.Id, &query)

	if err != nil {
		response.ServiceError(ctx, err)
		return
	}

	ctx.JSON(200, data)
}

func (c *controller) getOneOrder(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	_user := user.(models.User)

	data, err := c.service.getOneOrder(ctx.Param("id"), *_user.Id)

	if err != nil {
		response.ServiceError(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{
		"data": data,
	})
}

func (c *controller) createOrder(ctx *gin.Context) {
	var body orderScheme

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	user, _ := ctx.Get("user")
	_user := user.(models.User)

	data, err := c.service.order(&body, &_user)

	if err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 201, gin.H{
		"message": "Pesanan berhasil dibuat",
		"data":    data,
	})
}

func (c *controller) postReview(ctx *gin.Context) {
	var body postReviewScheme
	id := ctx.Param("id")

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	user, _ := ctx.Get("user")
	_user := user.(models.User)

	if err := c.service.postReview(&body, id, *_user.Id); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 200, gin.H{"message": "Ulasan berhasil dibuat"})
}

func (c *controller) updateOrder(ctx *gin.Context) {
	var body updateOrderScheme
	id := ctx.Param("id")

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	user, _ := ctx.Get("user")
	_user := user.(models.User)

	if err := c.service.updateOrder(&body, id, &_user); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	msg := "Pesanan berhasil"

	switch body.Status {
	case "SUCCESS":
		msg += " dikonfirmasi"
	case "CANCELED":
		msg += " dibatalkan"
	}

	response.Success(ctx, 200, gin.H{"message": msg})
}

func (c *controller) getTrxHistory(ctx *gin.Context) {
	var query getTrxHistoryQS
	user, _ := ctx.Get("user")
	_user := user.(models.User)

	ctx.ShouldBindQuery(&query)

	data, err := c.service.getTrxHistory(*_user.Id, &query)

	if err != nil {
		response.ServiceError(ctx, err)
		return
	}

	ctx.JSON(200, data)
}
