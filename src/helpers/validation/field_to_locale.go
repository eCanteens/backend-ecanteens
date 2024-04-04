package validation

func FieldToLocale(field string) (intlField string) {
	switch field {
	case "Name":
		intlField = "Nama"
	case "Email":
		intlField = "Email"
	case "Password":
		intlField = "Kata Sandi"
	}

	return
}
