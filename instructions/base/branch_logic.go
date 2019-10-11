package base

import (
	"github.com/zxh0/jvm.go/rtda"
)

func Branch(frame *rtda.Frame, offset int) {
	pc := frame.Thread.PC
	frame.NextPC = pc + offset
}
