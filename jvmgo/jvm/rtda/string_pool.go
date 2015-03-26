package rtda

import (
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
)

var _internedStrings = map[string]*rtc.Obj{}

func getInternedString(goStr string) *rtc.Obj {
	return _internedStrings[goStr]
}

// todo
func InternString(goStr string, jStr *rtc.Obj) *rtc.Obj {
	if internedStr, ok := _internedStrings[goStr]; ok {
		return internedStr
	}

	_internedStrings[goStr] = jStr
	return jStr
}
