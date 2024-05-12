package config

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type FilePath struct {
	Path string
	Url  string
}

func Upload(router *gin.Engine) {
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
}

func UploadPath(filename string) *FilePath {
	splitted := strings.Split(filename, ".")
	var fullname string
	var name string
	var ext string

	if len(splitted) > 1 {
		name = strings.Join(splitted[0:len(splitted)-1], "")
		ext = "." + splitted[len(splitted)-1]
	} else {
		name = filename
	}

	fullname = fmt.Sprintf("%s-%d%s", name, time.Now().Unix(), ext)
	absPath, _ := filepath.Abs("./public/uploads")

	return &FilePath{
		Path: fmt.Sprintf("%s/%s", absPath, fullname),
		Url: fmt.Sprintf("/public/uploads/%s", fullname),
	}
}
