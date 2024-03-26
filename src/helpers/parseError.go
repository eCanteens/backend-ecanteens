package helpers

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type ApiError struct {
	Field string `json:"field"`
	Msg   string `json:"msg"`
}

func ParseError(err error) ([]ApiError, error) {
	var ve validator.ValidationErrors

	if errors.As(err, &ve) {
		out := make([]ApiError, len(ve))
		for i, fe := range ve {
			out[i] = ApiError{
				Field: fe.Field(),
				Msg:   msgForTag(fe.Tag(), fe.Field()),
			}
		}
		return out, nil
	}

	return nil, err
}

func msgForTag(tag string, field string) string {
	var intlField string

	switch field {
	case "Name":
		intlField = "Nama"
	case "Email":
		intlField = "Email"
	case "Password":
		intlField = "Kata Sandi"
	}

	switch tag {
	case "required":
		return "Kolom " + intlField + " harus diisi"
	case "email":
		return "Email tidak valid"
	}
	return ""
}
