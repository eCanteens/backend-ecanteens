package auth

import (
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers/response"
	"github.com/eCanteens/backend-ecanteens/src/helpers/validation"
	"github.com/gin-gonic/gin"
)

type Controller interface {
	register(ctx *gin.Context)
	login(ctx *gin.Context)
	logout(ctx *gin.Context)
	google(ctx *gin.Context)
	setup(ctx *gin.Context)
	refresh(ctx *gin.Context)
	forgot(ctx *gin.Context)
	reset(ctx *gin.Context)
	profile(ctx *gin.Context)
	updateProfile(ctx *gin.Context)
	updatePassword(ctx *gin.Context)
	checkPin(ctx *gin.Context)
	updatePin(ctx *gin.Context)
}

type controller struct {
	service Service
}

func NewController(service Service) Controller {
	return &controller{
		service: service,
	}
}

func (c *controller) register(ctx *gin.Context) {
	var body registerScheme

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	err := c.service.register(&body)

	if err != nil {
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

	response.Success(ctx, 200, gin.H{"message": "Logout berhasil"})
}

func (c *controller) google(ctx *gin.Context) {
	var body googleScheme

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	data, token, err := c.service.google(&body)

	if err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 200, gin.H{
		"message":   "Login berhasil",
		"token":     token,
		"data":      data,
		"is_setted": data.Phone != nil,
	})
}

func (c *controller) setup(ctx *gin.Context) {
	var body setupScheme

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	user, _ := ctx.Get("user")
	_user := user.(models.User)

	if err := c.service.setupGoogle(&body, &_user); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 200, gin.H{
		"message": "Nomor telepon dan Kata sandi berhasil disimpan",
		"data":    _user,
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

	ctx.JSON(200, gin.H{
		"data": token,
	})
}

func (c *controller) forgot(ctx *gin.Context) {
	var body forgotScheme

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	if err := c.service.forgot(&body); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 200, gin.H{
		"message": "Email telah dikirim",
	})
}

func (c *controller) reset(ctx *gin.Context) {
	var body resetScheme

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	if err := c.service.reset(&body); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 200, gin.H{
		"message": "Password berhasil direset",
	})
}

func (c *controller) profile(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	_user := user.(models.User)
	_user.Password = ""
	_user.Wallet.IsPinSet = _user.Wallet.Pin != ""
	_user.Wallet.Pin = ""

	ctx.JSON(200, gin.H{
		"data": _user,
	})
}

func (c *controller) updateProfile(ctx *gin.Context) {
	var body updateScheme

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	user, _ := ctx.Get("user")
	_user := user.(models.User)

	if err := c.service.updateProfile(&_user, &body); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 200, gin.H{
		"message": "Profil berhasil diperbarui",
		"data":      _user,
	})
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

	response.Success(ctx, 200, gin.H{
		"message": "Password berhasil diperbarui",
	})
}

func (c *controller) checkPin(ctx *gin.Context) {
	var body checkPinScheme

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	user, _ := ctx.Get("user")
	_user := user.(models.User)

	if err := c.service.checkPin(&_user, &body); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 200, gin.H{
		"message": "Pin benar",
	})
}

func (c *controller) updatePin(ctx *gin.Context) {
	var body updatePinScheme

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	user, _ := ctx.Get("user")
	_user := user.(models.User)

	if err := c.service.updatePin(&_user, &body); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 200, gin.H{
		"message": "Pin berhasil diperbarui",
	})
}
