package util

import (
	"strings"
	"unicode/utf16"
)

func ReplaceAll(s, old, _new string) string {
	return strings.Replace(s, old, _new, -1)
}

// utf8 -> utf16
func StringToUtf16(s string) []uint16 {
	runes := []rune(s)
	return utf16.Encode(runes) // func Encode(s []rune) []uint16
}

// utf16 -> utf8
func Utf16ToString(s []uint16) string {
	runes := utf16.Decode(s) // func Decode(s []uint16) []rune
	return string(runes)
}
