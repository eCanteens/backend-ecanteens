package middleware

import (
	"strconv"
	"strings"

	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/helpers"
	"github.com/gin-gonic/gin"
)

func Cors(ctx *gin.Context) {
	origin := ctx.Request.Header.Get("Origin")

	exist := helpers.Find(&config.App.Cors.AllowOrigin, func(t *string) bool {
		return *t == origin || *t == "*"
	})

	if exist != nil {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", *exist)
	}
	ctx.Writer.Header().Set("Access-Control-Allow-Credentials", strconv.FormatBool(config.App.Cors.AllowCredential))
	ctx.Writer.Header().Set("Access-Control-Allow-Headers", strings.Join(config.App.Cors.AllowHeaders, ", "))
	ctx.Writer.Header().Set("Access-Control-Allow-Methods", strings.Join(config.App.Cors.AllowMethod, ", "))

	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(204)
		return
	}

	ctx.Next()
}