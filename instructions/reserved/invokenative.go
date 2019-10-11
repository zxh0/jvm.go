package reserved

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

// Invoke native method
type InvokeNative struct{ base.NoOperandsInstruction }

func (instr *InvokeNative) Execute(frame *rtda.Frame) {
	nativeMethod := frame.Method.GetNativeMethod().(func(*rtda.Frame))
	nativeMethod(frame)
}
