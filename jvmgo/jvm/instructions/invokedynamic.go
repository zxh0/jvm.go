package instructions

import (
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
)

// Invoke dynamic method
type invokedynamic struct {
	index uint16
	// 0
	// 0
}

func (self *invokedynamic) fetchOperands(decoder *InstructionDecoder) {
	self.index = decoder.readUint16()
	decoder.readUint8() // must be 0
	decoder.readUint8() // must be 0
}
func (self *invokedynamic) Execute(frame *rtda.Frame) {
	// todo
	panic("todo invokedynamic")
}
