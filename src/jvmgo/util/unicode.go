package util

import (
	"unicode/utf16"
	"unicode/utf8"
)

// utf8 -> utf16
// todo bytes: func Runes(s []byte) []rune
func StringToUtf16(str string) []uint16 {
	runeCount := utf8.RuneCountInString(str)
	runes := make([]rune, runeCount)
	i := 0
	for len(str) > 0 {
		r, size := utf8.DecodeRuneInString(str)
		runes[i] = r
		i++
		str = str[size:]
	}

	// func Encode(s []rune) []uint16
	return utf16.Encode(runes)
}

// utf16 -> utf8
func Utf16ToString(s []uint16) string {
	// func Decode(s []uint16) []rune
	runes := utf16.Decode(s)

	byteCount := 0
	for _, r := range runes {
		byteCount += utf8.RuneLen(r)
	}

	i := 0
	bytes := make([]byte, byteCount)
	for _, r := range runes {
		i += utf8.EncodeRune(bytes[i:], r)
	}

	return string(bytes)
}
