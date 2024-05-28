package config

func fieldToLocale(field string) (intlField string) {
	switch field {
	case "name":
		intlField = "Nama"
	case "email":
		intlField = "Email"
	case "phone":
		intlField = "Nomor Telepon"
	case "password":
		intlField = "Kata Sandi"
	case "old_password":
		intlField = "Kata Sandi Lama"
	case "new_password":
		intlField = "Kata Sandi Baru"
	case "pin":
		intlField = "Pin"
	case "refresh_token":
		intlField = "Refresh Token"
	case "id_token":
		intlField = "ID Token"
	default:
		intlField = field
	}

	return
}
