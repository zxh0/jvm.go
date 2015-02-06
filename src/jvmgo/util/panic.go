package util

import (
    "fmt"
    . "jvmgo/any"
)

func Panicf(format string, a ...Any) {
    panic(fmt.Sprintf(format, a))
}
