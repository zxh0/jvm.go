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

// todo: is there a better way to create String?
// go string -> java.lang.String
func JSFromGoStr(goStr string) *Object {
	if internedStr, found := _internedStrings[goStr]; found {
		return internedStr
	}

	chars := vmutils.UTF8ToUTF16(goStr)
	charArr := NewCharArray(chars)
	jStr := bootLoader.JLStringClass().NewObj()
	jStr.SetFieldValue("value", "[C", NewRefSlot(charArr))
	return JSIntern(goStr, jStr)
}

// java.lang.String -> go string
func JSToGoStr(jStr *Object) string {
	return vmutils.UTF16ToUTF8(JSToChars(jStr))
}

func JSToChars(jStr *Object) []uint16 {
	charArr := jStr.GetFieldValue("value", "[C").Ref
	return charArr.Chars()
}
