package validation

import (
	"github.com/gin-gonic/gin"
)

func Bind(ctx *gin.Context, model interface{}) (isValid bool) {
	if err := ctx.ShouldBind(model); err != nil {
		if parsed := ParseError(err); parsed != nil {
			ctx.AbortWithStatusJSON(400, gin.H{
				"errors":  &parsed,
				"message": parsed[0].Msg,
			})
			return
		}

		if parsed := ParseUnmarshalError(err); parsed != nil {
			ctx.AbortWithStatusJSON(500, gin.H{
				"message": parsed,
			})
			return
		}

		ctx.AbortWithStatusJSON(500, gin.H{
			"message": err.Error(),
		})
		return 
	}

	return true
}
