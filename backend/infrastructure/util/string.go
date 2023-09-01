package util

import (
	"regexp"
	"strings"
)

// RemoveAccentedVietnamese return string without accented vietnamese
func RemoveAccentedVietnamese(pattern string) string {

	accentedCharacters := []string{"À", "Á", "Â", "Ã", "È", "É",
		"Ê", "Ì", "Í", "Ò", "Ó", "Ô", "Õ", "Ù", "Ú", "Ý", "Ỳ", "Ỷ", "Ỵ", "Ỹ",
		"à", "á", "â", "ã", "è", "é", "ê", "ì", "í", "ò", "ó", "ô", "õ",
		"ù", "ú", "ý", "ỳ", "ỷ", "ỵ", "ỹ", "Ă", "ă", "Đ", "đ", "Ĩ", "ĩ", "Ũ",
		"ũ", "Ơ", "ơ", "Ư", "ư", "Ạ", "ạ", "Ả", "ả", "Ấ", "ấ", "Ầ", "ầ",
		"Ẩ", "ẩ", "Ẫ", "ẫ", "Ậ", "ậ", "Ắ", "ắ", "Ằ", "ằ", "Ẳ", "ẳ", "Ẵ",
		"ẵ", "Ặ", "ặ", "Ẹ", "ẹ", "Ẻ", "ẻ", "Ẽ", "ẽ", "Ế", "ế", "Ề", "ề",
		"Ể", "ể", "Ễ", "ễ", "Ệ", "ệ", "Ỉ", "ỉ", "Ị", "ị", "Ọ", "ọ", "Ỏ",
		"ỏ", "Ố", "ố", "Ồ", "ồ", "Ổ", "ổ", "Ỗ", "ỗ", "Ộ", "ộ", "Ớ", "ớ",
		"Ờ", "ờ", "Ở", "ở", "Ỡ", "ỡ", "Ợ", "ợ", "Ụ", "ụ", "Ủ", "ủ", "Ứ",
		"ứ", "Ừ", "ừ", "Ử", "ử", "Ữ", "ữ", "Ự", "ự"}

	charactersWithoutAccents := []string{"A", "A", "A", "A", "E",
		"E", "E", "I", "I", "O", "O", "O", "O", "U", "U", "Y", "Y", "Y", "Y",
		"Y", "a", "a", "a", "a", "e", "e", "e", "i", "i", "o", "o", "o",
		"o", "u", "u", "y", "y", "y", "y", "y", "A", "a", "D", "d", "I", "i",
		"U", "u", "O", "o", "U", "u", "A", "a", "A", "a", "A", "a", "A",
		"a", "A", "a", "A", "a", "A", "a", "A", "a", "A", "a", "A", "a",
		"A", "a", "A", "a", "E", "e", "E", "e", "E", "e", "E", "e", "E",
		"e", "E", "e", "E", "e", "E", "e", "I", "i", "I", "i", "O", "o",
		"O", "o", "O", "o", "O", "o", "O", "o", "O", "o", "O", "o", "O",
		"o", "O", "o", "O", "o", "O", "o", "O", "o", "U", "u", "U", "u",
		"U", "u", "U", "u", "U", "u", "U", "u", "U", "u"}
	for i, accentedCharacter := range accentedCharacters {
		pattern = strings.ReplaceAll(pattern, accentedCharacter, charactersWithoutAccents[i])
	}
	return pattern
}

// ToTrimSpace returns a slice of the string s, with all leading and trailing white space removed, as defined by Unicode.
func ToTrimSpace(s string) string {
	return strings.TrimSpace(s)
}

func ToLower(s string) string {
	return strings.ToLower(s)
}

// ToSearchFormat converts string to lowercase and remove accents
func ToSearchFormat(s string) string {
	return ToLower(RemoveAccentedVietnamese(strings.TrimSpace(s)))
}

// CheckStringInSlice checks if a string is present in a slice
func CheckStringInSlice(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

// IsValidDomainUrl check is valid domain
func IsValidDomainUrl(url string) bool {
	pattern := `([a-z0-9A-Z]\.)*[a-z0-9-]+\.([a-z0-9]{2,24})+(\.co\.([a-z0-9]{2,24})|\.([a-z0-9]{2,24}))*`
	match, _ := regexp.MatchString(pattern, url)

	return match
}
