package stores

import (
	"github.com/zxh0/jvm.go/jvmgo/instructions/base"
	"github.com/zxh0/jvm.go/jvmgo/rtda"
)

// Store int into local variable
type istore struct{ base.Index8Instruction }

func (self *istore) Execute(frame *rtda.Frame) {
	_istore(frame, uint(self.Index))
}

type istore_0 struct{ base.NoOperandsInstruction }

func (self *istore_0) Execute(frame *rtda.Frame) {
	_istore(frame, 0)
}

type istore_1 struct{ base.NoOperandsInstruction }

func (self *istore_1) Execute(frame *rtda.Frame) {
	_istore(frame, 1)
}

type istore_2 struct{ base.NoOperandsInstruction }

func (self *istore_2) Execute(frame *rtda.Frame) {
	_istore(frame, 2)
}

type istore_3 struct{ base.NoOperandsInstruction }

func (self *istore_3) Execute(frame *rtda.Frame) {
	_istore(frame, 3)
}

func _istore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopInt()
	frame.LocalVars().SetInt(index, val)
}
