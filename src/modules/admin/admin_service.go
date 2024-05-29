package admin

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func adminLoginService(body *AdminLoginScheme) (*models.User, *string, error) {
	var user models.User

	if err := findAdminEmail(&user, body.Email); err != nil {
		return nil, nil, errors.New("email admin salah")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return nil, nil, errors.New("password admin salah")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   *user.Id.Id,
		"email": user.Email,
		"exp":   0,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return nil, nil, err
	}

	user.Password = ""
	user.Wallet.Pin = ""

	return &user, &tokenString, nil
}

func dashboardService() (map[string]interface{}, error) {
	var userCount int64
	var restaurantCount int64

	if err := count("users", &userCount); err != nil {
		return nil, err
	}

	if err := count("restaurants", &restaurantCount); err != nil {
		return nil, err
	}

	data := map[string]interface{}{
		"users":       userCount,
		"restaurants": restaurantCount,
		"total":       userCount + restaurantCount,
	}

	return data, nil
}

func checkWalletService(phone string) (*models.User, error) {
	var user models.User

	if err := findUser(&user, phone); err != nil {
		return nil, errors.New("user tidak ditemukan")
	}

	return &user, nil
}

func topupWithdrawService(phone string, body *TopupWithdrawScheme, tipe string) (*models.Transaction, error) {
	var user models.User

	if err := findUser(&user, phone); err != nil {
		return nil, errors.New("user tidak ditemukan")
	}

	if err := topupWithdraw(body.Amount, &user, tipe); err != nil {
		return nil, err
	}

	data, err := createTransaction(&user, body.Amount, models.TransactionType(tipe))

	if err != nil {
		return nil, err
	}

	return data, nil
}

func transactionService(id string) (*models.Transaction, error) {
	var transaction models.Transaction

	data, err := findTransaction(&transaction, id)

	if err != nil {
		return nil, errors.New("transaksi tidak ditemukan")
	}

	return data, nil
}

func updateAdminProfileService(ctx *gin.Context, user *models.User, body *UpdateAdminProfileScheme) error {
	if err := checkUniqueService(body.Email, *user.Id.Id); err != nil {
		return err
	}

	user.Name = body.Name
	user.Email = body.Email

	if body.Avatar != nil {
		extracted := helpers.ExtractFileName(body.Avatar.Filename)
		filePath := helpers.UploadPath(fmt.Sprintf("avatar/user/%d.%s", *user.Id.Id, extracted.Ext))

		if err := ctx.SaveUploadedFile(body.Avatar, filePath.Path); err != nil {
			return err
		}

		user.Avatar = &filePath.Url
	}

	if err := save(user); err != nil {
		return err
	}

	user.Password = ""
	user.Wallet.Pin = ""

	return nil
}

func updateAdminPasswordService(user *models.User, body *UpdateAdminPasswordScheme) error {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.OldPassword)); err != nil {
		return errors.New("password salah")
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(body.NewPassword), bcrypt.DefaultCost)

	user.Password = string(hashed)

	return save(user)
}

func checkUniqueService(email string, id ...uint) error {
	sameUser := checkEmail(email, id...)

	if len(*sameUser) > 1 {
		return errors.New("email dan nomor telepon sudah digunakan")
	} else if len(*sameUser) == 1 {
		var fields []string

		if (*sameUser)[0].Email == email {
			fields = append(fields, "email")
		}

		errMsg := strings.Join(fields, " dan ") + " sudah digunakan"
		return errors.New(errMsg)
	}

	return nil
}
