package reserved

import (
	"fmt"

	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/native"
	"github.com/zxh0/jvm.go/rtda"
)

// Invoke native method
type InvokeNative struct{ base.NoOperandsInstruction }

func (instr *InvokeNative) Execute(frame *rtda.Frame) {
	method := frame.Method
	if frame.Thread.VMOptions.VerboseJNI {
		fmt.Printf("invokenative: %s.%s%s\n",
			method.Class.Name, method.Name, method.Descriptor)
	}

	// TODO: cache native method
	nativeMethod := native.FindNativeMethod(method)
	nativeMethod(frame)
}
