package utils

import "strings"

const (
	UndefinedString   = "undefined"
	ZeroString        = ""
	CensorStringValue = "***"
)

// StringTrimSpace -- trim space of string
func StringTrimSpace(s string) string {
	return strings.TrimSpace(s)
}

// IsStringEmpty -- check if string is empty
func IsStringEmpty(s string) bool {
	return s == ZeroString
}

// IsStringNotEmpty -- check if string is not empty
func IsStringNotEmpty(s string) bool {
	return s != ZeroString
}

// CensorString --
func CensorString(str string) string {
	if len(str) <= 6 {
		return CensorStringValue
	}

	return str[:2] + CensorStringValue + str[len(str)-2:]
}
