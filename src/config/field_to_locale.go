package config

import "strings"

func fieldToLocale(field string) (intlField string) {
	switch field {
	case "name":
		intlField = "Nama"
	case "phone":
		intlField = "Nomor Telepon"
	case "password":
		intlField = "Kata Sandi"
	case "old_password":
		intlField = "Kata Sandi Lama"
	case "new_password":
		intlField = "Kata Sandi Baru"
	case "refresh_token":
		intlField = "Refresh Token"
	case "id_token":
		intlField = "ID Token"
	case "restaurant_avatar":
		intlField = "Avatar Restoran"
	case "restaurant_name":
		intlField = "Nama Restoran"
	case "category_id":
		intlField = "Kategori"
	case "fullfilment_date":
		intlField = "Tanggal Preorder"
	case "payment_method":
		intlField = "Metode Pembayaran"
	default:
		intlField = strings.ToUpper(string(field[0])) + field[1:]
	}

	return
}
