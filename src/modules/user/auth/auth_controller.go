package auth

import (
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers/response"
	"github.com/eCanteens/backend-ecanteens/src/helpers/validation"
	"github.com/gin-gonic/gin"
)

func handleRegister(ctx *gin.Context) {
	var body registerScheme

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	err := registerService(&body)

	if err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 201, gin.H{"message": "Register berhasil"})
}

func handleLogin(ctx *gin.Context) {
	var body loginScheme

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	data, token, err := loginService(&body)

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

func handleLogout(ctx *gin.Context) {
	var body refreshScheme

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	if err := logoutService(&body); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 200, gin.H{"message": "Logout berhasil"})
}

func handleGoogle(ctx *gin.Context) {
	var body googleScheme

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	data, token, err := googleService(&body)

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

func handleSetup(ctx *gin.Context) {
	var body setupScheme

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	user, _ := ctx.Get("user")
	_user := user.(models.User)

	if err := setupGoogleService(&body, &_user); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 200, gin.H{
		"message": "Nomor telepon dan Kata sandi berhasil disimpan",
		"data":    _user,
	})
}

func handleRefresh(ctx *gin.Context) {
	var body refreshScheme

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	token, err := refreshService(&body)

	if err != nil {
		response.ServiceError(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{
		"data": token,
	})
}

func handleForgot(ctx *gin.Context) {
	var body forgotScheme

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	if err := forgotService(&body); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 200, gin.H{
		"message": "Email telah dikirim",
	})
}

func handleReset(ctx *gin.Context) {
	var body resetScheme

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	if err := resetService(&body); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 200, gin.H{
		"message": "Password berhasil direset",
	})
}

func handleProfile(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	_user := user.(models.User)
	_user.Password = ""
	_user.Wallet.IsPinSet = _user.Wallet.Pin != ""
	_user.Wallet.Pin = ""

	ctx.JSON(200, gin.H{
		"data": _user,
	})
}

func handleUpdateProfile(ctx *gin.Context) {
	var body updateScheme

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	user, _ := ctx.Get("user")
	_user := user.(models.User)

	if err := updateProfileService(&_user, &body); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 200, gin.H{
		"message": "Profil berhasil diperbarui",
		"data":      _user,
	})
}

func handleUpdatePassword(ctx *gin.Context) {
	var body updatePasswordScheme

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	user, _ := ctx.Get("user")
	_user := user.(models.User)

	if err := updatePasswordService(&_user, &body); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 200, gin.H{
		"message": "Password berhasil diperbarui",
	})
}

func handleCheckPin(ctx *gin.Context) {
	var body checkPinScheme

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	user, _ := ctx.Get("user")
	_user := user.(models.User)

	if err := checkPinService(&_user, &body); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 200, gin.H{
		"message": "Pin benar",
	})
}

func handleUpdatePin(ctx *gin.Context) {
	var body updatePinScheme

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	user, _ := ctx.Get("user")
	_user := user.(models.User)

	if err := updatePinService(&_user, &body); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 200, gin.H{
		"message": "Pin berhasil diperbarui",
	})
}
