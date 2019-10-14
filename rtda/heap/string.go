package heap

import (
	"github.com/zxh0/jvm.go/vmutils"
)

var _internedStrings = map[string]*Object{}

// todo
func JSIntern(goStr string, jStr *Object) *Object {
	if internedStr, ok := _internedStrings[goStr]; ok {
		return internedStr
	}

	_internedStrings[goStr] = jStr
	return jStr
}

// go string -> java.lang.String
func JSFromGoStr(goStr string) *Object {
	if internedStr, found := _internedStrings[goStr]; found {
		return internedStr
	}

	chars := vmutils.UTF8ToUTF16(goStr)
	jStr := JSFromChars(chars)
	jStr = JSIntern(goStr, jStr) // TODO
	return jStr
}

// java char[] -> java.lang.String
func JSFromChars(chars []uint16) *Object {
	charArr := NewCharArray(chars)
	jStr := bootLoader.JLStringClass().NewObj()
	jStr.SetFieldValue("value", "[C", NewRefSlot(charArr))
	return jStr
}

// java.lang.String -> go string
func JSToGoStr(jStr *Object) string {
	return vmutils.UTF16ToUTF8(JSToChars(jStr))
}

// java.lang.String -> java char[]
func JSToChars(jStr *Object) []uint16 {
	charArr := jStr.GetFieldValue("value", "[C").Ref
	return charArr.Chars()
}
