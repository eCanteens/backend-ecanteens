package upload

import (
	"fmt"
	"io"
	"mime/multipart"
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
	File        *multipart.FileHeader
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

func SaveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	if err = os.MkdirAll(filepath.Dir(dst), 0750); err != nil {
		return err
	}

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

func New(option *Option) (*FilePath, error) {
	extracted := ExtractFileName(option.File.Filename)
	filePath := UploadPath(fmt.Sprintf("%s/%s.%s", option.Folder, option.NewFilename, extracted.Ext))

	if err := SaveUploadedFile(option.File, filePath.Path); err != nil {
		return nil, err
	}

	return filePath, nil
}
