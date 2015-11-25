package references

import (
	"github.com/zxh0/jvm.go/jvmgo/instructions/base"
	"github.com/zxh0/jvm.go/jvmgo/rtda"
)

// Enter monitor for object
type monitorenter struct{ base.NoOperandsInstruction }

func (self *monitorenter) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		frame.RevertNextPC()
		thread.ThrowNPE()
	} else {
		ref.Monitor().Enter(thread)
	}
}

// Exit monitor for object
type monitorexit struct{ base.NoOperandsInstruction }

func (self *monitorexit) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		frame.RevertNextPC()
		thread.ThrowNPE()
	} else {
		ref.Monitor().Exit(thread)
	}
}
