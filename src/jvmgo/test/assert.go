package test

import (
    "fmt"
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
    x := fmt.Sprintf("%v", expected)
    y := fmt.Sprintf("%v", actual)
    if x != y {
        msg := fmt.Sprintf("expected: %v actual: %v", expected, actual)
        panic(msg)
    }
}
