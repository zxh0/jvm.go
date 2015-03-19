package instructions

import (
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
)

// Invoke a class (static) method
type invokestatic struct {
	Index16Instruction
	method *rtc.Method
}

func (self *invokestatic) Execute(frame *rtda.Frame) {
	if self.method == nil {
		cp := frame.Method().Class().ConstantPool()
		kMethodRef := cp.GetConstant(self.index).(*rtc.ConstantMethodref)
		self.method = kMethodRef.StaticMethod()
	}

	// init class
	class := self.method.Class()
	if class.InitializationNotStarted() {
		frame.RevertNextPC()
		frame.Thread().InitClass(class)
		return
	}

	frame.Thread().InvokeMethod(self.method)
}
