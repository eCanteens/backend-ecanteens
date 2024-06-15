package jwt

import (
	"crypto/rand"
	"encoding/base64"
	"os"
	"time"

	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/golang-jwt/jwt/v5"
)

type UserToken struct {
	Type         string `json:"type"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func generateRefreshToken() string {
	bytes := make([]byte, 40)
	rand.Read(bytes)
	return base64.URLEncoding.EncodeToString(bytes)
}

func GenerateUserToken(id uint, roleId uint) *UserToken {
	var token UserToken

	token.AccessToken = New(&jwt.MapClaims{
		"iss":  os.Getenv("BASE_URL"),
		"sub":  id,
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(config.App.Auth.AccessTokenExpiresIn).Unix(),
		"type": "access",
		"role": roleId,
	})

	token.RefreshToken = generateRefreshToken()
	token.Type = "Bearer"

	return &token
}

func GenerateResetToken(id uint) string {
	return New(&jwt.MapClaims{
		"iss":  os.Getenv("BASE_URL"),
		"sub":  id,
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Minute * 10).Unix(),
		"type": "reset",
	})
}
