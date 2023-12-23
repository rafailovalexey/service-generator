package utils

import "strings"

func Capitalize(value string) string {
	if value == "" {
		return value
	}

	return strings.ToUpper(value[:1]) + value[1:]
}

func Lowercase(value string) string {
	return strings.ToLower(value)
}

func SingularForm(value string) string {
	if value == "" {
		return value
	}

	return value[:1] + value[1:len(value)-1]
}

func FirstLetter(value string) string {
	return value[:1]
}
