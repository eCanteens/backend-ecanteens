package customerror

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func GormError(err error, data string) *CustomError {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return New(fmt.Sprintf("%s tidak ditemukan", data), 404)
	}

	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return New("Terdapat key yang duplikat", 400)
	}

	if errors.Is(err, gorm.ErrForeignKeyViolated) {
		return New("Foreign key tidak valid", 400)
	}

	var customErr *CustomError
	if errors.As(err, &customErr) {
		return customErr
	}

	fmt.Println(err.Error())
	return New("Terjadi kesalahan saat mengambil data", 500)
}
