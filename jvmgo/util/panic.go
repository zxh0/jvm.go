package util

import (
	"fmt"
)

func Panicf(format string, a ...interface{}) {
	panic(fmt.Sprintf(format, a...))
}
