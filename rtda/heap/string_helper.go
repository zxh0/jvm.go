package heap

import (
	"unicode/utf16"
)

// todo: is there a better way to create String?
// go string -> java.lang.String
func JString(goStr string) *Object {
	internedStr := getInternedString(goStr)
	if internedStr != nil {
		return internedStr
	}

	chars := _stringToUtf16(goStr)
	charArr := NewCharArray(chars)
	jStr := BootLoader().JLStringClass().NewObj()
	jStr.SetFieldValue("value", "[C", NewRefSlot(charArr))
	return InternString(goStr, jStr)
}

// java.lang.String -> go string
func GoString(jStr *Object) string {
	charArr := jStr.GetFieldValue("value", "[C").Ref
	return _utf16ToString(charArr.Chars())
}

// utf8 -> utf16
func _stringToUtf16(s string) []uint16 {
	runes := []rune(s)
	return utf16.Encode(runes) // func Encode(s []rune) []uint16
}

// utf16 -> utf8
func _utf16ToString(s []uint16) string {
	runes := utf16.Decode(s) // func Decode(s []uint16) []rune
	return string(runes)
}
