package middleware

import (
	"fmt"

	"github.com/eCanteens/backend-ecanteens/src/config"
	"github.com/eCanteens/backend-ecanteens/src/helpers/response"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// rate (20/s) and burst (30/s)
var limiter = rate.NewLimiter(rate.Limit(config.App.Limiter.Rate), config.App.Limiter.Burst)

func RateLimiter(ctx *gin.Context) {
	if !limiter.Allow() {
		response.Error(ctx, "Too Many Request", 429)
		fmt.Println("Too Many Request")
	}
}
