package auth

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/eCanteens/backend-ecanteens/src/helpers"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func RegisterService(user *models.User) error {
	sameUser := CheckEmailAndPhone(user)

	if len(sameUser) > 1 {
		return errors.New("email dan nomor telepon sudah digunakan")
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
		return errors.New(errMsg)
	}

	return RegisterRepo(user)
}

func LoginService(body *LoginSchema) (*string, error) {
	var user models.User

	if err := FindByEmail(&user, body.Email); err != nil {
		return nil, errors.New("email atau password salah")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return nil, errors.New("email atau password salah")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":    user.Name,
		"email":   user.Email,
		"phone":   user.Phone,
		"avatar":  user.Avatar,
		"balance": user.Balance,
		"exp":     float64(time.Now().Add(time.Hour * 24).Unix()),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}

func ForgotService(body *ForgotSchema) error {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": body.Email,
		"exp":   float64(time.Now().Add(time.Minute * 5).Unix()),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return err
	}

	return helpers.SendMail([]string{body.Email}, &helpers.MailMessage{
		Subject:     "Forgot Password",
		ContentType: helpers.HTML,
		Body: fmt.Sprintf(
			"<html><body><a href=\"%s/api/auth/reset-password/%s\">Reset Password</a></body></html>",
			os.Getenv("BASE_URL"),
			tokenString,
		),
	})
}

func ResetService(body *ResetSchema, token string) error {
	claim, err := helpers.ParseJwt(token)

	if err != nil {
		return err
	}

	user := models.User{Password: body.Password}

	return UpdatePassword(&user, claim["email"].(string))
}
