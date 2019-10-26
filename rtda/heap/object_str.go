package heap

import (
	"github.com/zxh0/jvm.go/vmutils"
)

// java.lang.String -> go string
func (obj *Object) JSToGoStr() string {
	charArr := obj.GetFieldValue("value", "[C").Ref
	return vmutils.UTF16ToUTF8(charArr.GetChars())
}
