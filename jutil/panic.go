package jutil

import (
	"fmt"
)

func Panicf(format string, a ...interface{}) {
	panic(fmt.Errorf(format, a...))
}
