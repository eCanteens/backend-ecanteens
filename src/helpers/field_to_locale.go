package helpers

import (
	"strings"

	"github.com/eCanteens/backend-ecanteens/src/config"
)

func fieldToLocale(field string) (string) {
	for k, v := range config.Locale {
		if field == k {
			return v
		}
	}

	words := strings.Split(field, "_")

	for i, word := range words {
		words[i] = strings.ToUpper(string(word[0])) + word[1:]
	}

	return strings.Join(words, " ")
}
