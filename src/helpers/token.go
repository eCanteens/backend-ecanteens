package helpers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func GenerateUserToken(id uint, roleId uint) *Token {
	var token Token

	token.AccessToken = GenerateJwt(&jwt.MapClaims{
		"iss":  os.Getenv("BASE_URL"),
		"sub":  id,
		"iat":  float64(time.Now().Unix()),
		"exp":  float64(time.Now().Add(time.Hour).Unix()),
		"type": "access",
		"role": roleId,
	})

	token.RefreshToken = GenerateJwt(&jwt.MapClaims{
		"iss":  os.Getenv("BASE_URL"),
		"sub":  id,
		"iat":  float64(time.Now().Unix()),
		"exp":  float64(time.Now().Add(time.Hour * 24 * 5).Unix()),
		"type": "refresh",
	})

	return &token
}

func GenerateResetToken(id uint) string {
	return GenerateJwt(&jwt.MapClaims{
		"iss":  os.Getenv("BASE_URL"),
		"sub":  id,
		"iat":  float64(time.Now().Unix()),
		"exp":  float64(time.Now().Add(time.Minute * 10).Unix()),
		"type": "reset",
	})
}
