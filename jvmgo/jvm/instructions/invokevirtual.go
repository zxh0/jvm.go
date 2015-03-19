package instructions

import (
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
)

// Invoke instance method; dispatch based on class
type invokevirtual struct {
	Index16Instruction
	kMethodRef *rtc.ConstantMethodref
	argCount   uint
}

func (self *invokevirtual) Execute(frame *rtda.Frame) {
	if self.kMethodRef == nil {
		cp := frame.Method().ConstantPool()
		self.kMethodRef = cp.GetConstant(self.index).(*rtc.ConstantMethodref)
		self.argCount = self.kMethodRef.ArgCount()
	}

	stack := frame.OperandStack()
	ref := stack.TopRef(self.argCount)
	if ref == nil {
		// frame.Thread().ThrowNPE()
		panic("NPE")
	}

	method := self.kMethodRef.GetVirtualMethod(ref)
	frame.Thread().InvokeMethod(method)
}
