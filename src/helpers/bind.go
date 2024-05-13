package helpers

import (
	"github.com/gin-gonic/gin"
)

func Bind(ctx *gin.Context, model interface{}) *gin.H {
	if err := ctx.ShouldBind(model); err != nil {
		parsed, parseErr := ParseError(err)

		if parseErr == nil {
			return ErrorResponse(&parsed)
		} else {
			return ErrorResponse(err.Error())
		}
	}

	return nil
}
