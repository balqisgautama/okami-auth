package util

import "unicode"

func IsPasswordStandardValid(password string) (bool, string, string) {
	if len(password) < 8 {
		return false, "NEED_MORE_THAN", "8"
	} else if len(password) >= 8 && len(password) <= 20 {
	next:
		for name, classes := range map[string][]*unicode.RangeTable{
			"UPPERCASE": {unicode.Upper, unicode.Title},
			"LOWERCASE": {unicode.Lower},
			"NUMERIC":   {unicode.Number, unicode.Digit},
			"SPECIAL":   {unicode.Space, unicode.Symbol, unicode.Punct, unicode.Mark},
		} {
			for _, r := range password {
				if unicode.IsOneOf(classes, r) {
					continue next
				}
			}
			return false, name, ""
		}
	} else {
		return false, "NEED_LESS_THAN", "20"
	}
	return true, "", ""

	// RegexPassword => Minimum eight and maximum 20 characters, at least one uppercase letter,
	// one lowercase letter, one number and one special character
}
