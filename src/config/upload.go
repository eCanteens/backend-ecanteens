package config

import (
	"github.com/gin-gonic/gin"
)

func Upload(router *gin.Engine) {
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
}
