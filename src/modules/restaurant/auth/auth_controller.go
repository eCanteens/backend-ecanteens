package auth

import (
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers/response"
	"github.com/eCanteens/backend-ecanteens/src/helpers/validation"
	"github.com/gin-gonic/gin"
)

type Controller interface {
	checkRegister(ctx *gin.Context)
	register(ctx *gin.Context)
	login(ctx *gin.Context)
	logout(ctx *gin.Context)
	refresh(ctx *gin.Context)
	profile(ctx *gin.Context)
	updateProfile(ctx *gin.Context)
	updateResto(ctx *gin.Context)
	updatePassword(ctx *gin.Context)
}

type controller struct {
	service Service
}

func NewController(service Service) Controller {
	return &controller{
		service: service,
	}
}

func (c *controller) checkRegister(ctx *gin.Context) {
	var body checkRegisterScheme

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	if err := c.service.checkUnique(body.Email, body.Phone); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 200, gin.H{"message": "Data sudah valid"})
}

func (c *controller) register(ctx *gin.Context) {
	var body registerScheme

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	if err := c.service.register(&body); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 201, gin.H{"message": "Register berhasil"})
}

func (c *controller) login(ctx *gin.Context) {
	var body loginScheme

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	data, token, err := c.service.login(&body)

	if err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 200, gin.H{
		"message": "Login berhasil",
		"token":   token,
		"data":    data,
	})
}

func (c *controller) logout(ctx *gin.Context) {
	var body refreshScheme

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	if err := c.service.logout(&body); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 200, gin.H{
		"message": "Logout berhasil",
	})
}

func (c *controller) refresh(ctx *gin.Context) {
	var body refreshScheme

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	token, err := c.service.refresh(&body)

	if err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 200, gin.H{
		"token": token,
	})
}

func (c *controller) profile(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	_user := user.(models.User)
	_user.Password = ""
	isPinSet := _user.Wallet.Pin != ""
	_user.Wallet.Pin = ""

	ctx.JSON(200, gin.H{
		"data":       _user,
		"is_pin_set": isPinSet,
	})
}

func (c *controller) updateProfile(ctx *gin.Context) {
	var body updateProfileScheme

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	user, _ := ctx.Get("user")
	_user := user.(models.User)

	if err := c.service.updateProfile(&body, &_user); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 200, gin.H{"message": "Profil berhasil diperbarui"})
}

func (c *controller) updateResto(ctx *gin.Context) {
	var body updateRestoScheme

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	user, _ := ctx.Get("user")
	_user := user.(models.User)

	if err := c.service.updateResto(&body, _user.Restaurant); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 200, gin.H{"message": "Restoran berhasil diperbarui"})
}

func (c *controller) updatePassword(ctx *gin.Context) {
	var body updatePasswordScheme

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	user, _ := ctx.Get("user")
	_user := user.(models.User)

	if err := c.service.updatePassword(&_user, &body); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 200, gin.H{"message": "Password berhasil diperbarui"})
}
