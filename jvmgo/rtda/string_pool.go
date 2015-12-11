package rtda

import (
	"github.com/zxh0/jvm.go/jvmgo/rtda/heap"
)

var _internedStrings = map[string]*heap.Object{}

func getInternedString(goStr string) *heap.Object {
	return _internedStrings[goStr]
}

// todo
func InternString(goStr string, jStr *heap.Object) *heap.Object {
	if internedStr, ok := _internedStrings[goStr]; ok {
		return internedStr
	}

	_internedStrings[goStr] = jStr
	return jStr
}
