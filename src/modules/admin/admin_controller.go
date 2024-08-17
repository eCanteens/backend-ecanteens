package admin

import (
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers/response"
	"github.com/eCanteens/backend-ecanteens/src/helpers/validation"
	"github.com/gin-gonic/gin"
)

type Controller interface {
	adminLogin(ctx *gin.Context)
	dashoard(ctx *gin.Context)
	checkWallet(ctx *gin.Context)
	topup(ctx *gin.Context)
	withdraw(ctx *gin.Context)
	transaction(ctx *gin.Context)
	mutasi(ctx *gin.Context)
	adminProfile(ctx *gin.Context)
	updateAdminProfile(ctx *gin.Context)
	updateAdminPassword(ctx *gin.Context)
}

type controller struct {
	service Service
}

func NewController(service Service) Controller {
	return &controller{
		service: service,
	}
}

// login
func (c *controller) adminLogin(ctx *gin.Context) {
	var body adminLoginScheme

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	data, token, err := c.service.adminLogin(&body)

	if err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 200, gin.H{
		"token": token,
		"data":  data,
		"message": "Login berhasil",
	})
}

// dashboard
func (c *controller) dashoard(ctx *gin.Context) {
	data, err := c.service.dashboard()

	if err != nil {
		response.ServiceError(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{
		"data": data,
	})
}

// check wallet
func (c *controller) checkWallet(ctx *gin.Context) {
	phone := ctx.Param("phone")

	data, err := c.service.checkWallet(phone)

	if err != nil {
		response.ServiceError(ctx, err)
		return
	}

	data.Password = ""
	data.Wallet.Pin = ""

	ctx.JSON(200, gin.H{
		"data": data,
	})
}

// topup
func (c *controller) topup(ctx *gin.Context) {
	phone := ctx.Param("phone")
	var body topupWithdrawScheme

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	data, err := c.service.topupWithdraw(phone, &body, "TOPUP")

	if err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 200, gin.H{
		"message": "Top Up berhasil",
		"data":    data,
	})
}

// withdraw
func (c *controller) withdraw(ctx *gin.Context) {
	phone := ctx.Param("phone")
	var body topupWithdrawScheme

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	data, err := c.service.topupWithdraw(phone, &body, "WITHDRAW")

	if err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 200, gin.H{
		"data":    data,
		"message": "Withdraw berhasil",
	})
}

// transaction
func (c *controller) transaction(ctx *gin.Context) {
	id := ctx.Param("id")

	data, err := c.service.transaction(id)

	if err != nil {
		response.ServiceError(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{
		"data": data,
	})
}

// mutasi
func (c *controller) mutasi(ctx *gin.Context) {
	var query mutationQS

	ctx.ShouldBindQuery(&query)

	data, err := c.service.mutasi(&query)

	if err != nil {
		response.ServiceError(ctx, err)
		return
	}

	ctx.JSON(200, data)
}

// profile
func (c *controller) adminProfile(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	_user := user.(models.User)
	_user.Password = ""

	ctx.JSON(200, gin.H{
		"data": _user,
	})
}

func (c *controller) updateAdminProfile(ctx *gin.Context) {
	var body updateAdminProfileScheme
	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	user, _ := ctx.Get("user")
	_user := user.(models.User)

	if err := c.service.updateAdminProfile(&_user, &body); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 200, gin.H{
		"message": "Profil berhasil diperbarui",
		"data":    _user,
	})
}

func (c *controller) updateAdminPassword(ctx *gin.Context) {
	var body updateAdminPasswordScheme
	user, _ := ctx.Get("user")
	_user := user.(models.User)

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	if err := c.service.updateAdminPassword(&_user, &body); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 200, gin.H{
		"message": "Password berhasil diperbarui",
	})
}
