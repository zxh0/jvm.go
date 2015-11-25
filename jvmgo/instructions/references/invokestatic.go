package references

import (
	"github.com/zxh0/jvm.go/jvmgo/instructions/base"
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/rtda/class"
)

// Invoke a class (static) method
type invokestatic struct {
	base.Index16Instruction
	method *rtc.Method
}

func (self *invokestatic) Execute(frame *rtda.Frame) {
	if self.method == nil {
		cp := frame.Method().Class().ConstantPool()
		k := cp.GetConstant(self.Index)
		if kMethodRef, ok := k.(*rtc.ConstantMethodref); ok {
			self.method = kMethodRef.StaticMethod()
		} else {
			self.method = k.(*rtc.ConstantInterfaceMethodref).StaticMethod()
		}
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
