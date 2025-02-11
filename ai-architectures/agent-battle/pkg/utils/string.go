package utils

import (
	"math/rand"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	UndefinedString   = "undefined"
	ZeroString        = ""
	CensorStringValue = "***"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

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

// StringPrefixInSlice --
func StringPrefixInSlice(list []string, str string) bool {
	for _, v := range list {
		if strings.HasPrefix(str, v) {
			return true
		}
	}
	return false
}

// StringInSliceEqualFold --
func StringInSliceEqualFold(list []string, str string) bool {
	for _, v := range list {
		if strings.EqualFold(v, str) {
			return true
		}
	}
	return false
}

// StringKeys --
func StringKeys(mmap map[string]interface{}) []string {
	keys := make([]string, 0, len(mmap))
	for k := range mmap {
		keys = append(keys, k)
	}
	return keys
}

// IsUndefinedValue
func IsUndefinedValue(s string) bool {
	return strings.EqualFold(s, UndefinedString)
}

// RandString give a random string
func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GetFirstLastFromFullname(fullName string) (string, string) {
	splitFullname := strings.SplitN(strings.TrimSpace(fullName), " ", 2)
	caser := cases.Title(language.English)
	firstName := caser.String(splitFullname[0])
	if len(splitFullname) == 1 {
		return firstName, ""
	}
	lastName := caser.String(splitFullname[1])
	return firstName, lastName
}
