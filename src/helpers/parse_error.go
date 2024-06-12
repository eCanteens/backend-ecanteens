package helpers

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type ApiError struct {
	Field string `json:"field"`
	Msg   string `json:"message"`
}

func ParseError(err error) ([]ApiError, error) {
	var ve validator.ValidationErrors

	if errors.As(err, &ve) {
		out := make([]	ApiError, len(ve))
		for i, fe := range ve {
			out[i] = ApiError{
				Field: ToSnakeCase(fe.Field()),
				Msg:   MsgForTag(fe.Tag(), fe.Field(), fe.Param()),
			}
		}
		return out, nil
	}

	return nil, err
}
