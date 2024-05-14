package admin

import (
	"errors"
	"os"
	"time"

	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func adminLoginService(body *LoginScheme) (*models.User, *string, error) {
	var user models.User

	if err := findAdminEmail(&user, body.Email); err != nil {
		return nil, nil, errors.New("email admin salah")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return nil, nil, errors.New("password admin salah")
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
	user.Pin = nil

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
		"users": userCount,
		"restaurants": restaurantCount,
		"total": userCount + restaurantCount,
	}

	return data, nil
}