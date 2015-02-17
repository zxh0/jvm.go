package util

import (
	"unicode/utf16"
	// "unicode/utf8"
)

// utf8 -> utf16
func StringToUtf16(str string) []uint16 {
	runes := []rune(str)
	return utf16.Encode(runes) // func Encode(s []rune) []uint16
}

// utf16 -> utf8
func Utf16ToString(s []uint16) string {
	runes := utf16.Decode(s) // func Decode(s []uint16) []rune
	return string(runes)
}
