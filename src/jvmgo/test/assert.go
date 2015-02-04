package test

import (
    "fmt"
    "reflect"
    . "jvmgo/any"
)

func AssertTrue(b bool) {
    if !b {
        panic("Not true!")
    }
}
func AssertFalse(b bool) {
    if b {
        panic("Not false!")
    }
}

func AssertNil(any Any) {
    v := fmt.Sprintf("%v", any)
    if v != "<nil>" {
        panic("Not nil! " + v)   
    }
}
func AssertNotNil(any Any) {
    if any == nil {
        panic("nil!")
    }
}

func AssertSame(expected, actual Any) {
    if expected != actual {
        msg := fmt.Sprintf("Not same! expected: %v actual: %v", expected, actual)
        panic(msg)
    }
}

func AssertEquals(expected, actual Any) {
    switch expected.(type) {
    case int8, uint8, int16, uint16, int32, uint32, int64, uint64, int, uint, float32, float64:
        x := fmt.Sprintf("%v", expected)
        y := fmt.Sprintf("%v", actual)
        if x == y {
            return
        }
    }

    if reflect.DeepEqual(expected, actual) {
        return
    }

    msg := fmt.Sprintf("Not equals! expected: %v actual: %v", expected, actual)
    panic(msg)
}
