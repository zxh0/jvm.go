package instructions

import (
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
)

// Invoke instance method;
// special handling for superclass, private, and instance initialization method invocations
type invokespecial struct{ Index16Instruction }

func (self *invokespecial) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	kMethodRef := cp.GetConstant(self.index).(*rtc.ConstantMethodref)
	method := kMethodRef.SpecialMethod()
	frame.Thread().InvokeMethod(method)
}
