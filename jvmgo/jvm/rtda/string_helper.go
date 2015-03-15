package rtda

import (
	rtc "jvmgo/jvm/rtda/class"
	"jvmgo/util"
)

// todo: is there a better way to create String?
func NewJString(goStr string) *rtc.Obj {
	chars := util.StringToUtf16(goStr)
	internedStr := getInternedString(chars)
	if internedStr != nil {
		return internedStr
	}

	charArr := rtc.NewCharArray(chars)
	jStr := rtc.BootLoader().JLStringClass().NewObj()
	jStr.SetFieldValue("value", "[C", charArr)
	return InternString(chars, jStr)
}

func GoString(jStr *rtc.Obj) string {
	utf16 := JStringChars(jStr)
	return util.Utf16ToString(utf16)
}

func JStringChars(jStr *rtc.Obj) []uint16 {
	charArr := jStr.GetFieldValue("value", "[C").(*rtc.Obj)
	return charArr.Fields().([]uint16)
}
