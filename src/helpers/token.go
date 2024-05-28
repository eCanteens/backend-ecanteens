package helpers

import (
	"os"
	"time"

	"github.com/eCanteens/backend-ecanteens/src/database/models"
	"github.com/golang-jwt/jwt/v5"
)

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func GenerateUserToken(user *models.User) *Token {
	var token Token

	token.AccessToken = GenerateJwt(&jwt.MapClaims{
		"iss":   os.Getenv("BASE_URL"),
		"sub":   *user.Id.Id,
		"iat":   float64(time.Now().Unix()),
		"exp":   float64(time.Now().Add(time.Hour).Unix()),
		"type":  "access",
		"email": user.Email,
	})

	token.RefreshToken = GenerateJwt(&jwt.MapClaims{
		"iss":  os.Getenv("BASE_URL"),
		"sub":  *user.Id.Id,
		"iat":  float64(time.Now().Unix()),
		"exp":  float64(time.Now().Add(time.Hour * 24 * 5).Unix()),
		"type": "refresh",
	})

	return &token
}

func GenerateResetToken(user *models.User) string {
	return GenerateJwt(&jwt.MapClaims{
		"iss":  os.Getenv("BASE_URL"),
		"sub":  *user.Id.Id,
		"iat":  float64(time.Now().Unix()),
		"exp":  float64(time.Now().Add(time.Minute * 10).Unix()),
		"type": "reset",
	})
}
