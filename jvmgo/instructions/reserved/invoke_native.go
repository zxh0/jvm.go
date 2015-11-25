package reserved

import (
	"github.com/zxh0/jvm.go/jvmgo/instructions/base"
	"github.com/zxh0/jvm.go/jvmgo/rtda"
)

// Invoke native method
type invoke_native struct{ base.NoOperandsInstruction }

func (self *invoke_native) Execute(frame *rtda.Frame) {
	nativeMethod := frame.Method().NativeMethod().(func(*rtda.Frame))
	nativeMethod(frame)
}
