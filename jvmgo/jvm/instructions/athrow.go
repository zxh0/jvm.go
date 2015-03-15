package instructions

import (
	"jvmgo/jvm/rtda"
)

// Throw exception or error
type athrow struct{ NoOperandsInstruction }

func (self *athrow) Execute(frame *rtda.Frame) {
	ex := frame.OperandStack().PopRef()
	if ex == nil {
		frame.Thread().ThrowNPE()
		return
	}

	thread := frame.Thread()
	for {
		frame := thread.CurrentFrame()
		pc := frame.NextPC() - 1

		handler := frame.Method().FindExceptionHandler(ex.Class(), pc)
		if handler != nil {
			stack := frame.OperandStack()
			stack.Clear()
			stack.PushRef(ex)
			frame.SetNextPC(handler.HandlerPc())
			return
		}

		thread.PopFrame()
		if thread.IsStackEmpty() {
			break
		}
	}

	thread.HandleUncaughtException(ex)
}
