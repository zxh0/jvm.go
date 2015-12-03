package rtda

import (
	rtc "github.com/zxh0/jvm.go/jvmgo/rtda/class"
)

var _internedStrings = map[string]*rtc.Object{}

func getInternedString(goStr string) *rtc.Object {
	return _internedStrings[goStr]
}

// todo
func InternString(goStr string, jStr *rtc.Object) *rtc.Object {
	if internedStr, ok := _internedStrings[goStr]; ok {
		return internedStr
	}

	_internedStrings[goStr] = jStr
	return jStr
}
