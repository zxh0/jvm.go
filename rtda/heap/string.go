package heap

import (
	"github.com/zxh0/jvm.go/vmutils"
)

var _internedStrings = map[string]*Object{}

func getInternedString(goStr string) *Object {
	return _internedStrings[goStr]
}

// todo
func InternString(goStr string, jStr *Object) *Object {
	if internedStr, ok := _internedStrings[goStr]; ok {
		return internedStr
	}

	_internedStrings[goStr] = jStr
	return jStr
}

// todo: is there a better way to create String?
// go string -> java.lang.String
func JString(goStr string) *Object {
	internedStr := getInternedString(goStr)
	if internedStr != nil {
		return internedStr
	}

	chars := vmutils.UTF8ToUTF16(goStr)
	charArr := NewCharArray(chars)
	jStr := BootLoader().JLStringClass().NewObj()
	jStr.SetFieldValue("value", "[C", NewRefSlot(charArr))
	return InternString(goStr, jStr)
}

// java.lang.String -> go string
func GoString(jStr *Object) string {
	charArr := jStr.GetFieldValue("value", "[C").Ref
	return vmutils.UTF16ToUTF8(charArr.Chars())
}
