package validation

func MsgForTag(tag string, field string, param string) string {
	intlField := FieldToLocale(field)

	switch tag {
	case "required":
		return intlField + " harus diisi"
	case "email":
		return "Email tidak valid"
	case "min":
		return intlField + " minimal " + param + " karakter"
	case "max":
		return intlField + " maximal " + param + " karakter"
	case "len":
		return intlField + " harus " + param + " karakter"
	case "unique":
		return intlField + " sudah digunakan"
	case "numeric":
		return intlField + " harus berupa angka"
	}
	return ""
}