package instructions

import (
	"jvmgo/jvm/rtda"
	rtc "jvmgo/jvm/rtda/class"
)

// Invoke instance method;
// special handling for superclass, private, and instance initialization method invocations
type invokespecial struct{ Index16Instruction }

func (self *invokespecial) Execute(frame *rtda.Frame) {
	cp := frame.ConstantPool()
	k := cp.GetConstant(self.index)

	if kMethodRef, ok := k.(*rtc.ConstantMethodref); ok {
		method := kMethodRef.SpecialMethod()
		frame.Thread().InvokeMethod(method)
		return
	}

	kInterfaceMethodRef := k.(*rtc.ConstantInterfaceMethodref)
	method := kInterfaceMethodRef.SpecialMethod()
	frame.Thread().InvokeMethod(method)
}
