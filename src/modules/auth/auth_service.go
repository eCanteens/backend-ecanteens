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

type ResetPasswordProps struct {
	LOGO string
	URL  string
	NAME string
}

func ForgotService(body *ForgotSchema) error {
	var user models.User

	if err := FindByEmail(&user, body.Email); err != nil {
		return err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": body.Email,
		"exp":   float64(time.Now().Add(time.Minute * 5).Unix()),
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

	return helpers.SendMail([]string{body.Email}, &helpers.MailMessage{
		Subject:     "Forgot Password",
		ContentType: helpers.HTML,
		HtmlBody:    t,
		HTMLProps: &ResetPasswordProps{
			LOGO: fmt.Sprintf("%s/assets/logo.svg", os.Getenv("BASE_URL")),
			URL:  fmt.Sprintf("%s/api/auth/reset-password/%s", os.Getenv("BASE_URL"), tokenString),
			NAME: user.Name,
		},
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
