package constants

import (
	"github.com/zxh0/jvm.go/jvmgo/instructions/base"
	"github.com/zxh0/jvm.go/jvmgo/rtda"
)

// Do nothing
type nop struct{ base.NoOperandsInstruction }

func (self *nop) Execute(frame *rtda.Frame) {
	// really do nothing
}
