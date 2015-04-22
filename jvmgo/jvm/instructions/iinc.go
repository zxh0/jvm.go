package instructions

import (
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
)

// Increment local variable by constant
type iinc struct {
	index  uint
	_const int32
}

func (self *iinc) fetchOperands(decoder *InstructionDecoder) {
	self.index = uint(decoder.readUint8())
	self._const = int32(decoder.readInt8())
}

func (self *iinc) Execute(frame *rtda.Frame) {
	localVars := frame.LocalVars()
	val := localVars.GetInt(self.index)
	val += self._const
	localVars.SetInt(self.index, val)
}
