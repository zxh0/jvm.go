package instructions

import (
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
)

// Invoke interface method
type invokeinterface struct {
	index uint16
	count uint8 // unused
	// 0
}

func (self *invokeinterface) fetchOperands(decoder *InstructionDecoder) {
	self.index = decoder.readUint16()
	self.count = decoder.readUint8()
	decoder.readUint8() // must be 0
}

func (self *invokeinterface) Execute(frame *rtda.Frame) {
	cp := frame.Method().ConstantPool()
	kMethodRef := cp.GetConstant(uint(self.index)).(*rtc.ConstantMethodref)

	stack := frame.OperandStack()
	ref := stack.Top(kMethodRef.ArgCount())
	if ref == nil {
		panic("NPE") // todo
	}

	method := kMethodRef.FindInterfaceMethod(ref.(*rtc.Obj))
	frame.Thread().InvokeMethod(method)
}
