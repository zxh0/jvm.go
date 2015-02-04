package jvm

import (
    "fmt"
    rtc "jvmgo/jvm/rtda/class"
)

func LogJString(jStr *rtc.Obj) {
    charArr := rtc.GetFieldValue(jStr, "value", "[C").(*rtc.Obj)
    chars := charArr.Fields().([]uint16)
    // todo
    for _, char := range chars {
        fmt.Printf("%c", char)
    }
    fmt.Println()
}
