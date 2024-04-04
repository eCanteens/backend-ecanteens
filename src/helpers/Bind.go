package helpers

import (
	"github.com/eCanteens/backend-ecanteens/src/helpers/validation"
	"github.com/gin-gonic/gin"
)

func Bind(ctx *gin.Context, model interface{}) *gin.H {
	if err := ctx.ShouldBindJSON(model); err != nil {
		parsed, parseErr := validation.ParseError(err)

		if parseErr == nil {
			return &gin.H{
				"error": &parsed,
			}
		} else {
			return &gin.H{
				"error": err.Error(),
			}
		}
	}

	return nil
}