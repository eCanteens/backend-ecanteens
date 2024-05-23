package auth

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func registerService(body *RegisterScheme) (*models.User, error) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

	if err := findByEmail(&models.User{}, body.Email); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			user := &models.User{
				Name:     body.Name,
				Email:    body.Email,
				Password: string(hashed),
			}

			if err := create(user); err != nil {
				return nil, err
			}

			user.Password = ""

			return user, nil
		}

		return nil, err
	}

	return nil, errors.New("email sudah digunakan")
}

func loginService(body *LoginScheme) (*models.User, *string, error) {
	var user models.User

	if err := findByEmail(&user, body.Email); err != nil {
		return nil, nil, errors.New("email atau password salah")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return nil, nil, errors.New("email atau password salah")
	}

	tokenString := helpers.GenerateJwt(&jwt.MapClaims{
		"id":    *user.Id.Id,
		"email": user.Email,
		"exp":   0,
	})

	user.Password = ""
	user.Wallet.Pin = ""

	return &user, &tokenString, nil
}

func googleService(body *GoogleScheme) (*models.User, *string, *bool, error) {
	var user models.User
	var isPasswordSet bool

	if err := findByEmail(&user, body.Email); err != nil {
		isPasswordSet = false
		user.Name = body.Name
		user.Email = body.Email
		user.Avatar = body.Avatar

		if err := create(&user); err != nil {
			return nil, nil, nil, err
		}
	}

	isPasswordSet = user.Password != ""
	user.Password = ""
	user.Wallet.Pin = ""

	tokenString := helpers.GenerateJwt(&jwt.MapClaims{
		"id":  *user.Id.Id,
		"exp": 0,
	})

	return &user, &tokenString, &isPasswordSet, nil
}

func forgotService(body *ForgotScheme) error {
	var user models.User

	if err := findByEmail(&user, body.Email); err != nil {
		return errors.New("pengguna tidak ditemukan")
	}

	tokenString := helpers.GenerateJwt(&jwt.MapClaims{
		"id":  *user.Id.Id,
		"exp": float64(time.Now().Add(time.Minute * 10).Unix()),
	})

	absPath, _ := filepath.Abs("./src/templates/reset-password.html")
	t, err := template.ParseFiles(absPath)
	if err != nil {
		return err
	}

	type ResetPasswordProps struct {
		LOGO string
		URL  string
		NAME string
	}

	return helpers.SendMail([]string{body.Email}, &helpers.MailMessage{
		Subject:     "Forgot Password",
		ContentType: helpers.HTML,
		HtmlBody:    t,
		HTMLProps: &ResetPasswordProps{
			LOGO: fmt.Sprintf("%s/public/assets/logo.png", os.Getenv("BASE_URL")),
			URL:  fmt.Sprintf("%s/api/auth/new-password/%s", os.Getenv("BASE_URL"), tokenString),
			NAME: user.Name,
		},
	})
}

func resetService(body *ResetScheme) error {
	claim, err := helpers.ParseJwt(body.Token)

	if err != nil {
		return err
	}

	id, ok := claim["id"].(float64)
	if !ok {
		return errors.New("cant convert claims")
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

	user := &models.User{Password: string(hashed)}

	return updatePassword(uint(id), user)
}

func updateProfileService(ctx *gin.Context, user *models.User, body *UpdateScheme) error {
	user.Name = body.Name
	user.Email = body.Email

	if body.Phone != "" {
		user.Phone = &body.Phone
	}

	sameUser := checkEmailAndPhone(user, user.Id.Id)

	if len(*sameUser) > 1 {
		return errors.New("email dan nomor telepon sudah digunakan")
	} else if len(*sameUser) == 1 {
		var fields []string

		if (*sameUser)[0].Email == user.Email {
			fields = append(fields, "email")
		}

		if (*sameUser)[0].Phone != nil && user.Phone != nil {
			if *(*sameUser)[0].Phone == *user.Phone {
				fields = append(fields, "nomor telepon")
			}
		}

		errMsg := strings.Join(fields, " dan ") + " sudah digunakan"
		return errors.New(errMsg)
	}

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

func updatePasswordService(user *models.User, body *UpdatePasswordScheme) error {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.OldPassword)); err != nil {
		return errors.New("password salah")
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(body.NewPassword), bcrypt.DefaultCost)

	user.Password = string(hashed)

	return save(user)
}

func checkPinService(user *models.User, body *CheckPinScheme) error {
	if user.Wallet.Pin == "" {
		return errors.New("pin belum di set")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Wallet.Pin), []byte(body.Pin)); err != nil {
		return errors.New("pin salah")
	}

	return nil
}

func updatePinService(user *models.User, body *UpdatePinScheme) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(body.Pin), bcrypt.DefaultCost)
	if err != nil {
		return errors.New(err.Error())
	}
	user.Wallet.Pin = string(hashed)

	return save(user.Wallet)
}
