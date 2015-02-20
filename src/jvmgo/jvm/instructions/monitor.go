package instructions

import "jvmgo/jvm/rtda"

// Enter monitor for object
type monitorenter struct{ NoOperandsInstruction }

func (self *monitorenter) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		frame.RevertNextPC()
		thread.ThrowException("java/lang/NullPointerException", "()V", nil)
		return
	}

	ref.Monitor().Enter(thread)
}

// Exit monitor for object
type monitorexit struct{ NoOperandsInstruction }

func (self *monitorexit) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		frame.RevertNextPC()
		thread.ThrowException("java/lang/NullPointerException", "()V", nil)
		return
	}

	ref.Monitor().Exit(thread)
}
