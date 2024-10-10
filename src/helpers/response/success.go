package response

import (
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil"
)

func Success(ctx *gin.Context, statusCode int, data ...gin.H) {
	var response = gin.H{
		"status": "success",
	}

	if len(data) > 0 {
		for key, value := range data[0] {
			if !goutil.IsEmpty(value) {
				response[key] = value
			}

		}
	}

	ctx.JSON(statusCode, &response)
}
