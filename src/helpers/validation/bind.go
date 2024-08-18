package validation

import (
	"github.com/gin-gonic/gin"
)

func Bind(ctx *gin.Context, model interface{}) (isValid bool) {
	if err := ctx.ShouldBind(model); err != nil {
		if parsed := ParseError(err); parsed != nil {
			ctx.AbortWithStatusJSON(400, gin.H{
				"status": "error",
				"message": parsed[0].Msg,
				"errors":  &parsed,
			})
			return
		}

		if parsed := ParseUnmarshalError(err); parsed != nil {
			ctx.AbortWithStatusJSON(400, gin.H{
				"status": "error",
				"message": parsed,
			})
			return
		}

		ctx.AbortWithStatusJSON(500, gin.H{
			"status": "error",
			"message": err.Error(),
		})
		return 
	}

	return true
}
