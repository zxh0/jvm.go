package rtda

import (
    "unicode/utf8"
    "unicode/utf16"
    rtc "jvmgo/rtda/class"
)

// todo: is there a better way to create String?
func NewJString(goStr string, frame *Frame) ([]uint16, *rtc.Obj) {
    classLoader := frame.Method().Class().ClassLoader()
    stringClass := classLoader.StringClass()
    chars := string2chars(goStr) // utf16
    jCharArr := rtc.NewCharArray(chars, classLoader)
    jStr := stringClass.NewObj()
    stringClass.GetField("value", "[C").PutValue(jStr, jCharArr)
    return chars, jStr
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
