package rtda

import (
    "unicode/utf8"
    "unicode/utf16"
    rtc "jvmgo/jvm/rtda/class"
)

func JStringChars(jStr *rtc.Obj) ([]uint16) {
    charArr := jStr.GetFieldValue("value", "[C").(*rtc.Obj)
    return charArr.Fields().([]uint16)
}

// todo: is there a better way to create String?
func NewJString(goStr string, frame *Frame) (*rtc.Obj) {
    chars := string2chars(goStr) // utf16
    internedStr := getInternedString(chars)
    if internedStr != nil {
        return internedStr
    }

    classLoader := frame.Method().Class().ClassLoader()
    stringClass := classLoader.StringClass()
    jCharArr := rtc.NewCharArray(chars, classLoader)
    jStr := stringClass.NewObj()
    jStr.SetFieldValue("value", "[C", jCharArr)
    return InternString(chars, jStr)
}

func string2chars(str string) ([]uint16) {
    runeCount := utf8.RuneCountInString(str)
    codePoints := make([]rune, runeCount)
    i := 0
    for len(str) > 0 {
        r, size := utf8.DecodeRuneInString(str)
        codePoints[i] = r
        i++
        str = str[size:]
    }

    // func Encode(s []rune) []uint16
    return utf16.Encode(codePoints)
}
