package test

import (
    "fmt"
    "reflect"
    . "jvmgo/any"
)

func AssertNil(any Any) {
    AssertEquals(nil, any)
}
func AssertNotNil(any Any) {
    if any == nil {
        panic("nil!")
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

    msg := fmt.Sprintf("expected: %v actual: %v", expected, actual)
    panic(msg)
}
