package rtda

import (
    "unicode/utf8"
    rtc "jvmgo/rtda/class"
)

func NewJString(goStr string, thread *Thread) (*rtc.Obj) {
    // new string
    stringClass := thread.CurrentFrame().Method().Class().ClassLoader().StringClass()
    jStr := stringClass.NewObj()
    
    // init string
    codePoints := string2CodePoints(goStr)
    initMethod := stringClass.GetMethod("<init>", "([III)V") //public String(int[] codePoints, int offset, int count)
    newFrame := thread.NewFrame(initMethod)
    localVars := newFrame.LocalVars()
    localVars.SetRef(0, jStr) // this
    localVars.SetRef(1, rtc.NewIntArray(codePoints))
    localVars.SetInt(2, 0)
    localVars.SetInt(3, int32(len(codePoints)))
    thread.PushFrame(newFrame)

    return jStr
}

func string2CodePoints(str string) ([]rune) {
    runeCount := utf8.RuneCountInString(str)
    codePoints := make([]rune, runeCount)
    i := 0
    for len(str) > 0 {
        r, size := utf8.DecodeRuneInString(str)
        codePoints[i] = r
        i++
        str = str[size:]
    }
    return codePoints
}
