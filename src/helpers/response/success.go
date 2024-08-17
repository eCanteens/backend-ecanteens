package response

import "github.com/gin-gonic/gin"

func Success(ctx *gin.Context, statusCode int, data ...gin.H) {
	var response = gin.H{
		"status":  "success",
	}

	if len(data) > 0 {
		for key, value := range data[0] {
			response[key] = value
		}
	}

	ctx.JSON(statusCode, &response)
}