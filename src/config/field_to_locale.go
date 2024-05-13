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
		intlField = "Password Lama"
	case "NewPassword":
		intlField = "Password Baru"
	case "Pin":
		intlField = "Pin"
	case "NewPin":
		intlField = "Pin Baru"
	case "IsLike":
		intlField = "IsLike"
	default:
		intlField = field
	}

	return
}
