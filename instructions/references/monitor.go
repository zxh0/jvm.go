package references

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

// Enter monitor for object
type MONITOR_ENTER struct{ base.NoOperandsInstruction }

func (instr *MONITOR_ENTER) Execute(frame *rtda.Frame) {
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
type MONITOR_EXIT struct{ base.NoOperandsInstruction }

func (instr *MONITOR_EXIT) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		frame.RevertNextPC()
		thread.ThrowNPE()
	} else {
		ref.Monitor().Exit(thread)
	}
}
