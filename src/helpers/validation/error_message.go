package validation

func MsgForTag(tag string, field string, param string) string {
	intlField := FieldToLocale(field)

	switch tag {
	case "required":
		return intlField + " harus diisi"
	case "email":
		return "Email tidak valid"
	case "min":
		return intlField + " minimal " + param
	case "unique":
		return intlField + " harus unik"
	}
	return ""
}