package instructions

import "jvmgo/jvm/rtda"

// Enter monitor for object
type monitorenter struct{ NoOperandsInstruction }

func (self *monitorenter) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		ref.Monitor().Enter(thread)
	} else {
		frame.RevertNextPC()
		thread.ThrowException("java/lang/NullPointerException", "()V", nil)
	}
}

// Exit monitor for object
type monitorexit struct{ NoOperandsInstruction }

func (self *monitorexit) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		ref.Monitor().Exit(thread)
	} else {
		// todo
		panic("NPE")
	}
}
