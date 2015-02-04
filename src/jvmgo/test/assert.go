package test

import (
    "fmt"
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
    x := fmt.Sprintf("%v", expected)
    y := fmt.Sprintf("%v", actual)
    if x != y {
        msg := fmt.Sprintf("Not equals! expected: %v actual: %v", expected, actual)
        panic(msg)
    }
}
