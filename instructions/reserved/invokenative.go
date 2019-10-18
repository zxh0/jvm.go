package reserved

import (
	"fmt"

	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

// Invoke native method
type InvokeNative struct{ base.NoOperandsInstruction }

func (instr *InvokeNative) Execute(frame *rtda.Frame) {
	if frame.Thread.VMOptions.VerboseJNI {
		fmt.Printf("invokenative: %s.%s%s\n",
			frame.Method.Class.Name, frame.Method.Name, frame.Method.Descriptor)
	}

	nativeMethod := frame.Method.GetNativeMethod().(func(*rtda.Frame))
	nativeMethod(frame)
}
