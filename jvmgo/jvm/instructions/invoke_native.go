package instructions

import "jvmgo/jvm/rtda"

// Invoke native method
type invoke_native struct{ NoOperandsInstruction }

func (self *invoke_native) Execute(frame *rtda.Frame) {
	nativeMethod := frame.Method().NativeMethod().(func(*rtda.Frame))
	nativeMethod(frame)
}
