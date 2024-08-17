package validation

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/eCanteens/backend-ecanteens/src/helpers"
	"github.com/go-playground/validator/v10"
)

type ApiError struct {
	Field string `json:"field"`
	Msg   string `json:"message"`
}

func ParseError(err error) []ApiError {
	var ve validator.ValidationErrors

	if errors.As(err, &ve) {
		out := make([]ApiError, len(ve))

		for i, fe := range ve {
			out[i] = ApiError{
				Field: helpers.ToSnakeCase(fe.Field()),
				Msg:   MsgForTag(fe.Tag(), fe.Field(), fe.Param()),
			}
		}
		return out
	}

	return nil
}

func ParseUnmarshalError(err error) *string {
	var ue *json.UnmarshalTypeError

	if errors.As(err, &ue) {
		return helpers.PointerTo(fmt.Sprintf("Gagal memproses data, pastikan `%s` berupa `%s`", ue.Field, ue.Type))
	}
	return nil
}
