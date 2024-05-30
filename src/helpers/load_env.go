package helpers

import (
	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	if err := godotenv.Load(); err != nil {
		panic(err.Error())
	}
}
