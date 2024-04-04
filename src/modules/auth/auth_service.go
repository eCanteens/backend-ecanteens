package auth

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/eCanteens/backend-ecanteens/src/database/models"
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
		return nil, errors.New("bad credentials")
	}

	fmt.Println(user)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return nil, errors.New("bad credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}
