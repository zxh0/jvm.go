package math

import (
	"github.com/zxh0/jvm.go/jvmgo/instructions/base"
	"github.com/zxh0/jvm.go/jvmgo/rtda"
)

// Increment local variable by constant
type iinc struct {
	index  uint
	_const int32
}

func (self *iinc) fetchOperands(reader *base.BytecodeReader) {
	self.index = uint(reader.ReadUint8())
	self._const = int32(reader.ReadInt8())
}

func (self *iinc) Execute(frame *rtda.Frame) {
	localVars := frame.LocalVars()
	val := localVars.GetInt(self.index)
	val += self._const
	localVars.SetInt(self.index, val)
}
