package heap

import (
	"github.com/zxh0/jvm.go/vmutils"
)

// java.lang.String -> go string
func JSToGoStr(jStr *Object) string {
	return vmutils.UTF16ToUTF8(JSToChars(jStr))
}

// java.lang.String -> java char[]
func JSToChars(jStr *Object) []uint16 {
	charArr := jStr.GetFieldValue("value", "[C").Ref
	return charArr.GetChars()
}
