package transaction

import (
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers/response"
	"github.com/eCanteens/backend-ecanteens/src/helpers/validation"
	"github.com/gin-gonic/gin"
)

func getCart(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	_user := user.(models.User)

	data, err := getCartService(&_user)

	if err != nil {
		response.ServiceError(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{
		"data": data,
	})
}

func updateCart(ctx *gin.Context) {
	id := ctx.Param("id")
	var body updateCartNoteScheme

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	user, _ := ctx.Get("user")
	_user := user.(models.User)

	if err := updateCartService(id, &body, *_user.Id); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 200, gin.H{"message": "Catatan berhasil ditambahkan"})
}

func addCart(ctx *gin.Context) {
	var body addCartScheme

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	user, _ := ctx.Get("user")
	_user := user.(models.User)

	if err := addCartService(&_user, &body); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	if *body.Quantity == 0 {
		response.Success(ctx, 200, gin.H{"message": "Produk berhasil dihapus dari keranjang"})
		} else {
		response.Success(ctx, 201, gin.H{"message": "Produk berhasil ditambahkan ke keranjang"})
	}
}

func getOrder(ctx *gin.Context) {
	var query getOrderQS
	user, _ := ctx.Get("user")
	_user := user.(models.User)

	ctx.ShouldBindQuery(&query)

	data, err := getOrderService(*_user.Id, &query)

	if err != nil {
		response.ServiceError(ctx, err)
		return
	}

	ctx.JSON(200, data)
}

func handleOrder(ctx *gin.Context) {
	var body orderScheme

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	user, _ := ctx.Get("user")
	_user := user.(models.User)

	data, err := orderService(&body, &_user)

	if err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 201, gin.H{
		"message": "Pesanan berhasil dibuat",
		"data": data,
	})
}

func handlePostReview(ctx *gin.Context) {
	var body postReviewScheme
	id := ctx.Param("id")

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	user, _ := ctx.Get("user")
	_user := user.(models.User)

	if err := postReviewService(&body, id, *_user.Id); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 200, gin.H{"message": "Ulasan berhasil dibuat"})
}

func handleUpdateOrder(ctx *gin.Context) {
	var body updateOrderScheme
	id := ctx.Param("id")

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	user, _ := ctx.Get("user")
	_user := user.(models.User)

	if err := updateOrderService(&body, id, &_user); err != nil {
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
