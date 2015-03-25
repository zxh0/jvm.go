package instructions

import (
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
)

// Invoke interface method
type invokeinterface struct {
	index uint
	// count uint8
	// zero uint8

	// optimization
	kMethodRef   *rtc.ConstantMethodref
	argSlotCount uint
}

func (self *invokeinterface) fetchOperands(decoder *InstructionDecoder) {
	self.index = uint(decoder.readUint16())
	decoder.readUint8() // count
	decoder.readUint8() // must be 0
}

func (self *invokeinterface) Execute(frame *rtda.Frame) {
	if self.kMethodRef == nil {
		cp := frame.Method().ConstantPool()
		self.kMethodRef = cp.GetConstant(self.index).(*rtc.ConstantMethodref)
		self.argSlotCount = self.kMethodRef.ArgSlotCount()
	}

	stack := frame.OperandStack()
	ref := stack.TopRef(self.argSlotCount)
	if ref == nil {
		panic("NPE") // todo
	}

	method := self.kMethodRef.FindInterfaceMethod(ref)
	frame.Thread().InvokeMethod(method)
}
