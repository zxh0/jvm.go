package test

import (
    "fmt"
    . "jvmgo/any"
)

func AssertNotNil(any Any) {
    if any == nil {
        panic("nil!")
    }
}

func AssertEquals(expected, actual Any) {
    if expected == nil && actual == nil {
        return
    }
    if expected == nil || actual == nil {
        panicNotEquals(expected, actual)
    }

    switch expected.(type) {
    case int8, int16, int32, int64:
        switch actual.(type) {
            case int8, int16, int32, int64:
                if expected == actual {
                    return
                }
        }
    }

    panicNotEquals(expected, actual)
}

func panicNotEquals(expected, actual Any) {
    msg := fmt.Sprintf("expected: %v actual: %v", expected, actual)
    panic(msg)
}
