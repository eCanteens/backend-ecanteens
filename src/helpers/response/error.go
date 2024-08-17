package response

import (
	"errors"

	"github.com/eCanteens/backend-ecanteens/src/helpers/customerror"
	"github.com/gin-gonic/gin"
)

func ServiceError(ctx *gin.Context, err error) {
	if err != nil {
		var customErr *customerror.CustomError
		if errors.As(err, &customErr) {
			ctx.JSON(customErr.StatusCode, gin.H{"message": customErr.Message})
		} else {
			ctx.JSON(500, gin.H{"message": "Terjadi kesalahan"})
		}
	}
}

func Error(ctx *gin.Context, msg string, statusCode int) {
	ctx.AbortWithStatusJSON(statusCode, gin.H{
		"status": "error",
		"message": msg,
	})
}
