package auth

import (
	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers/response"
	"github.com/eCanteens/backend-ecanteens/src/helpers/validation"
	"github.com/gin-gonic/gin"
)

func handleCheckRegister(ctx *gin.Context) {
	var body checkRegisterScheme

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	if err := checkUniqueService(body.Email, body.Phone); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 200, gin.H{"message": "Data sudah valid"})
}

func handleRegister(ctx *gin.Context) {
	var body registerScheme

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	if err := registerService(&body); err != nil {
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
		"token": token,
		"data":  data,
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

	response.Success(ctx, 200, gin.H{
		"message": "Logout berhasil",
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

	response.Success(ctx, 200, gin.H{
		"token": token,
	})
}

func handleProfile(ctx *gin.Context) {
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

func handleUpdateProfile(ctx *gin.Context) {
	var body updateProfileScheme

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	user, _ := ctx.Get("user")
	_user := user.(models.User)

	if err := updateProfileService(&body, &_user); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 200, gin.H{"message": "Profil berhasil diperbarui"})
}

func handleUpdateResto(ctx *gin.Context) {
	var body updateRestoScheme

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	user, _ := ctx.Get("user")
	_user := user.(models.User)

	if err := updateRestoService(&body, _user.Restaurant); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	response.Success(ctx, 200, gin.H{"message": "Restoran berhasil diperbarui"})
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

	response.Success(ctx, 200, gin.H{"message": "Password berhasil diperbarui"})
}
