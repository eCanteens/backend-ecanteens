package config

func fieldToLocale(field string) (intlField string) {
	switch field {
	case "Name":
		intlField = "Nama"
	case "Email":
		intlField = "Email"
	case "Phone":
		intlField = "Nomor Telepon"
	case "Password":
		intlField = "Kata Sandi"
	case "OldPassword":
		intlField = "Kata Sandi Lama"
	case "NewPassword":
		intlField = "Kata Sandi Baru"
	case "Pin":
		intlField = "Pin"
	case "NewPin":
		intlField = "Pin Baru"
	default:
		intlField = field
	}

	return
}
