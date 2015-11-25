package stores

import (
	"github.com/zxh0/jvm.go/jvmgo/instructions/base"
	"github.com/zxh0/jvm.go/jvmgo/rtda"
)

// Store reference into local variable
type astore struct{ base.Index8Instruction }

func (self *astore) Execute(frame *rtda.Frame) {
	_astore(frame, uint(self.Index))
}

type astore_0 struct{ base.NoOperandsInstruction }

func (self *astore_0) Execute(frame *rtda.Frame) {
	_astore(frame, 0)
}

type astore_1 struct{ base.NoOperandsInstruction }

func (self *astore_1) Execute(frame *rtda.Frame) {
	_astore(frame, 1)
}

type astore_2 struct{ base.NoOperandsInstruction }

func (self *astore_2) Execute(frame *rtda.Frame) {
	_astore(frame, 2)
}

type astore_3 struct{ base.NoOperandsInstruction }

func (self *astore_3) Execute(frame *rtda.Frame) {
	_astore(frame, 3)
}

func _astore(frame *rtda.Frame, index uint) {
	ref := frame.OperandStack().PopRef()
	frame.LocalVars().SetRef(index, ref)
}
