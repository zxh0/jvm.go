package rtda

import (
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
	"github.com/zxh0/jvm.go/jvmgo/util"
)

// todo: is there a better way to create String?
func NewJString(goStr string) *rtc.Obj {
	internedStr := getInternedString(goStr)
	if internedStr != nil {
		return internedStr
	}

	chars := util.StringToUtf16(goStr)
	charArr := rtc.NewCharArray(chars)
	jStr := rtc.BootLoader().JLStringClass().NewObj()
	jStr.SetFieldValue("value", "[C", charArr)
	return InternString(goStr, jStr)
}

func GoString(jStr *rtc.Obj) string {
	utf16 := JStringChars(jStr)
	return util.Utf16ToString(utf16)
}

func JStringChars(jStr *rtc.Obj) []uint16 {
	charArr := jStr.GetFieldValue("value", "[C").(*rtc.Obj)
	return charArr.Fields().([]uint16)
}
