package instructions

import (
	//"fmt"
	"jvmgo/jvm/rtda"
)

// Throw exception or error
type athrow struct{ NoOperandsInstruction }

func (self *athrow) Execute(frame *rtda.Frame) {
	ex := frame.OperandStack().PopRef()
	if ex == nil {
		// todo NPE
		panic("athrow NPE")
	}

	thread := frame.Thread()
	for {
		frame := thread.CurrentFrame()
		stack := frame.OperandStack()

		handler := frame.Method().FindExceptionHandler(ex.Class(), thread.PC())
		if handler != nil {
			stack.PushRef(ex)
			frame.SetNextPC(handler.HandlerPc())
			return
		}

		thread.PopFrame()
		if !thread.IsStackEmpty() {
			thread.SetPC(thread.TopFrame().NextPC())
		} else {
			break
		}
	}

	thread.HandleUncaughtException(ex)
}
