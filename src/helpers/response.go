package helpers

import (
	"github.com/gin-gonic/gin"
)

type Data map[string]interface{}

func ErrorResponse(errMsg interface{}) *gin.H {
	return &gin.H{
		"message": errMsg,
		"status":  "error",
	}
}

func SuccessResponse(msg interface{}, data ...Data) *gin.H {
	response := &gin.H{
		"message": msg,
		"status":  "success",
	}

	if len(data) > 0 {
		for key, value := range data[0] {
			(*response)[key] = value
		}
	}

	return response
}
