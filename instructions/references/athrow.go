package references

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

// Throw exception or error
type AThrow struct{ base.NoOperandsInstruction }

func (instr *AThrow) Execute(frame *rtda.Frame) {
	thread := frame.Thread

	ex := frame.PopRef()
	if ex == nil {
		thread.ThrowNPE()
		return
	}

	for {
		frame := thread.CurrentFrame()
		pc := frame.NextPC - 1

		handlerPC := frame.Method.FindExceptionHandler(ex.Class, pc) // TODO
		if handlerPC >= 0 {
			frame.ClearStack()
			frame.PushRef(ex)
			frame.NextPC = handlerPC
			return
		}

		thread.PopFrame()
		if thread.IsStackEmpty() {
			break
		}
	}

	thread.HandleUncaughtException(ex)
}
