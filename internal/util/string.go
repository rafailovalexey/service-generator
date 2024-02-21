package util

import (
	"strings"
)

func GetWithUpperCaseFirstLetter(value string) string {
	if value == "" {
		return value
	}

	return strings.ToUpper(value[:1]) + value[1:]
}

func GetFirstLetterLowerCase(value string) string {
	return strings.ToLower(value[:1])
}
