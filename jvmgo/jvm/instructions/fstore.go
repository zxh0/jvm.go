package instructions

import (
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
)

// Store float into local variable
type fstore struct{ Index8Instruction }

func (self *fstore) Execute(frame *rtda.Frame) {
	_fstore(frame, uint(self.index))
}

type fstore_0 struct{ NoOperandsInstruction }

func (self *fstore_0) Execute(frame *rtda.Frame) {
	_fstore(frame, 0)
}

type fstore_1 struct{ NoOperandsInstruction }

func (self *fstore_1) Execute(frame *rtda.Frame) {
	_fstore(frame, 1)
}

type fstore_2 struct{ NoOperandsInstruction }

func (self *fstore_2) Execute(frame *rtda.Frame) {
	_fstore(frame, 2)
}

type fstore_3 struct{ NoOperandsInstruction }

func (self *fstore_3) Execute(frame *rtda.Frame) {
	_fstore(frame, 3)
}

func _fstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopFloat()
	frame.LocalVars().SetFloat(index, val)
}
