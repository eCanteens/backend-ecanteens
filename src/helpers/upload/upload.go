package upload

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type FilePath struct {
	Path string
	Url  string
}

type FileName struct {
	Name string
	Ext  string
}

type Option struct {
	Folder      string
	Filename    string
	NewFilename string
}

func ExtractFileName(filename string) *FileName {
	splitted := strings.Split(filename, ".")
	var name string
	var ext string

	if len(splitted) > 1 {
		name = strings.Join(splitted[0:len(splitted)-1], "")
		ext = splitted[len(splitted)-1]
	} else {
		name = filename
	}

	return &FileName{
		Name: name,
		Ext:  ext,
	}
}

func UploadPath(filename string) *FilePath {
	absPath, _ := filepath.Abs("./public/uploads")

	return &FilePath{
		Path: fmt.Sprintf("%s/%s", absPath, filename),
		Url:  fmt.Sprintf("%s/public/uploads/%s", os.Getenv("BASE_URL"), filename),
	}
}

func New(option *Option) *FilePath {
	extracted := ExtractFileName(option.Filename)
	return UploadPath(fmt.Sprintf("%s/%s.%s", option.Folder, option.NewFilename, extracted.Ext))
}
