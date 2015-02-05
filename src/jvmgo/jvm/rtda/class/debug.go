package class

import (
    "fmt"
)

func LogJString(jStr *Obj) {
    charArr := jStr.GetFieldValue("value", "[C").(*Obj)
    chars := charArr.fields.([]uint16)
    // todo
    for _, char := range chars {
        fmt.Printf("%c", char)
    }
    fmt.Println()
}
