package instructions

import (
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/rtda/class"
)

// Invoke instance method;
// special handling for superclass, private, and instance initialization method invocations
type invokespecial struct{ Index16Instruction }

func (self *invokespecial) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	k := cp.GetConstant(self.index)
	if kMethodRef, ok := k.(*rtc.ConstantMethodref); ok {
		method := kMethodRef.SpecialMethod()
		frame.Thread().InvokeMethod(method)
	} else {
		method := k.(*rtc.ConstantInterfaceMethodref).SpecialMethod()
		frame.Thread().InvokeMethod(method)
	}
}
