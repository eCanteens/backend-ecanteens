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
)

func registerService(body *RegisterScheme) (*models.User, error) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

	user := &models.User{
		Name:     body.Name,
		Email:    body.Email,
		Password: string(hashed),
	}

	sameUser := checkEmailAndPhone(user)

	if len(sameUser) > 1 {
		return nil, errors.New("email dan nomor telepon sudah digunakan")
	} else if len(sameUser) == 1 {
		var fields []string

		if sameUser[0].Email == user.Email {
			fields = append(fields, "email")
		}

		if sameUser[0].Phone != nil && user.Phone != nil {
			if *sameUser[0].Phone == *user.Phone {
				fields = append(fields, "phone")
			}
		}

		errMsg := strings.Join(fields, " dan ") + " sudah digunakan"
		return nil, errors.New(errMsg)
	}

	if err := create(user); err != nil {
		return nil, err
	}

	user.Password = ""

	return user, nil
}

func loginService(body *LoginScheme) (*models.User, *string, error) {
	var user models.User

	if err := findByEmail(&user, body.Email); err != nil {
		return nil, nil, errors.New("email atau password salah")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return nil, nil, errors.New("email atau password salah")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":   user.Email,
		"exp":     float64(time.Now().Add(time.Hour * 24).Unix()),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return nil, nil, err
	}

	user.Password = ""
	user.Pin = ""

	return &user, &tokenString, nil
}

func forgotService(body *ForgotScheme) error {
	var user models.User

	if err := findByEmail(&user, body.Email); err != nil {
		return errors.New("pengguna tidak ditemukan")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  *user.Id.Id,
		"exp": float64(time.Now().Add(time.Minute * 5).Unix()),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return err
	}

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

	id, ok := claim["user_id"].(float64)
	if !ok {
		return errors.New("cant convert claims")
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

	user := &models.User{Password: string(hashed)}

	return updatePassword(uint(id), user)
}

func updateProfileService(ctx *gin.Context, user *models.User, body *UpdateScheme) (*models.User, error) {
	user.Name = body.Name
	user.Email = body.Email
	user.Phone = body.Phone

	sameUser := checkEmailAndPhone(user, user.Id.Id)

	if len(sameUser) > 1 {
		return nil, errors.New("email dan nomor telepon sudah digunakan")
	} else if len(sameUser) == 1 {
		var fields []string

		if sameUser[0].Email == user.Email {
			fields = append(fields, "email")
		}

		if sameUser[0].Phone != nil && user.Phone != nil {
			if *sameUser[0].Phone == *user.Phone {
				fields = append(fields, "nomor telepon")
			}
		}

		errMsg := strings.Join(fields, " dan ") + " sudah digunakan"
		return nil, errors.New(errMsg)
	}

	if body.Avatar != nil {
		extracted := helpers.ExtractFileName(body.Avatar.Filename)
		filePath := helpers.UploadPath(fmt.Sprintf("avatar/user/%d.%s", *user.Id.Id, extracted.Ext))
		
		if err := ctx.SaveUploadedFile(body.Avatar, filePath.Path); err != nil {
			return nil, err
		}

		user.Avatar = &filePath.Url
	}

	if err := save(user); err != nil {
		return nil, err
	}

	user.Password = ""
	user.Pin = ""

	return user, nil
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
	if err := bcrypt.CompareHashAndPassword([]byte(user.Pin), []byte(body.Pin)); err != nil {
		return errors.New("pin salah")
	}

	return nil
}

func updatePinService(user *models.User, body *UpdatePinScheme) error {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(body.Pin), bcrypt.DefaultCost)
	user.Pin = string(hashed)

	return save(user)
}