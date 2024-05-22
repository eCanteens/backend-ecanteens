package helpers

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type Data map[string]interface{}

func ErrorResponse(errMsg interface{}, data ...map[string]interface{}) *gin.H {
	var response gin.H

	switch v := errMsg.(type) {
	case string:
		msg := strings.ToUpper(string(v[0])) + v[1:]

		response = gin.H{
			"message": msg,
			"status":  "error",
		}

	default:
		response = gin.H{
			"message": errMsg,
			"status":  "error",
		}
	}

	if len(data) > 0 {
		for key, value := range data[0] {
			response[key] = value
		}
	}

	return &response
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
