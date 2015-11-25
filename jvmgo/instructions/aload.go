package instructions

import (
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
)

// Load reference from local variable
type aload struct{ Index8Instruction }

func (self *aload) Execute(frame *rtda.Frame) {
	_aload(frame, uint(self.index))
}

type aload_0 struct{ NoOperandsInstruction }

func (self *aload_0) Execute(frame *rtda.Frame) {
	_aload(frame, 0)
}

type aload_1 struct{ NoOperandsInstruction }

func (self *aload_1) Execute(frame *rtda.Frame) {
	_aload(frame, 1)
}

type aload_2 struct{ NoOperandsInstruction }

func (self *aload_2) Execute(frame *rtda.Frame) {
	_aload(frame, 2)
}

type aload_3 struct{ NoOperandsInstruction }

func (self *aload_3) Execute(frame *rtda.Frame) {
	_aload(frame, 3)
}

func _aload(frame *rtda.Frame, index uint) {
	ref := frame.LocalVars().GetRef(index)
	frame.OperandStack().PushRef(ref)
}
