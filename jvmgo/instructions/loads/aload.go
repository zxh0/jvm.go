package loads

import (
	"github.com/zxh0/jvm.go/jvmgo/instructions/base"
	"github.com/zxh0/jvm.go/jvmgo/rtda"
)

// Load reference from local variable
type ALOAD struct{ base.Index8Instruction }

func (self *ALOAD) Execute(frame *rtda.Frame) {
	_aload(frame, uint(self.Index))
}

type ALOAD_0 struct{ base.NoOperandsInstruction }

func (self *ALOAD_0) Execute(frame *rtda.Frame) {
	_aload(frame, 0)
}

type ALOAD_1 struct{ base.NoOperandsInstruction }

func (self *ALOAD_1) Execute(frame *rtda.Frame) {
	_aload(frame, 1)
}

type ALOAD_2 struct{ base.NoOperandsInstruction }

func (self *ALOAD_2) Execute(frame *rtda.Frame) {
	_aload(frame, 2)
}

type ALOAD_3 struct{ base.NoOperandsInstruction }

func (self *ALOAD_3) Execute(frame *rtda.Frame) {
	_aload(frame, 3)
}

func _aload(frame *rtda.Frame, index uint) {
	ref := frame.LocalVars().GetRef(index)
	frame.OperandStack().PushRef(ref)
}
