package loads

import (
	"github.com/zxh0/jvm.go/jvmgo/instructions/base"
	"github.com/zxh0/jvm.go/jvmgo/rtda"
)

// Load reference from local variable
type aload struct{ base.Index8Instruction }

func (self *aload) Execute(frame *rtda.Frame) {
	_aload(frame, uint(self.Index))
}

type aload_0 struct{ base.NoOperandsInstruction }

func (self *aload_0) Execute(frame *rtda.Frame) {
	_aload(frame, 0)
}

type aload_1 struct{ base.NoOperandsInstruction }

func (self *aload_1) Execute(frame *rtda.Frame) {
	_aload(frame, 1)
}

type aload_2 struct{ base.NoOperandsInstruction }

func (self *aload_2) Execute(frame *rtda.Frame) {
	_aload(frame, 2)
}

type aload_3 struct{ base.NoOperandsInstruction }

func (self *aload_3) Execute(frame *rtda.Frame) {
	_aload(frame, 3)
}

func _aload(frame *rtda.Frame, index uint) {
	ref := frame.LocalVars().GetRef(index)
	frame.OperandStack().PushRef(ref)
}
