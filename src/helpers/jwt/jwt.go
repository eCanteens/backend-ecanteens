package jwt

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type MapClaims = jwt.MapClaims

func New(claim *jwt.MapClaims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_KEY")))

	return tokenString
}

func Parse(token string) (jwt.MapClaims, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("JWT_KEY")), nil
	})

	if err != nil {
		return nil, err
	}

	claim, _ := parsedToken.Claims.(jwt.MapClaims)
	return claim, err
}
