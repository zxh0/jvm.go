package heap

import (
	"github.com/zxh0/jvm.go/vmutils"
)

// java.lang.String -> go string
func (obj *Object) JSToGoStr() string {
	jByteArr := obj.GetFieldValue("value", "[B").Ref
	coder := obj.GetFieldValue("coder", "B").IntValue()

	if coder == 0 { // latin
		return string(jByteArr.GetGoBytes())
	} else {
		jBytes := jByteArr.GetBytes()
		uint16s := vmutils.CastInt8sToUint16s(jBytes)
		bytes := vmutils.UTF16ToUTF8(uint16s)
		return bytes
	}
}
