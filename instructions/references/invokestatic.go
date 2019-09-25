package references

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

// Invoke a class (static) method
type INVOKE_STATIC struct {
	base.Index16Instruction
	method *heap.Method
}

func (instr *INVOKE_STATIC) Execute(frame *rtda.Frame) {
	if instr.method == nil {
		cp := frame.Method().Class().ConstantPool()
		k := cp.GetConstant(instr.Index)
		if kMethodRef, ok := k.(*heap.ConstantMethodref); ok {
			instr.method = kMethodRef.StaticMethod()
		} else {
			instr.method = k.(*heap.ConstantInterfaceMethodref).StaticMethod()
		}
	}

	// init class
	class := instr.method.Class()
	if class.InitializationNotStarted() {
		frame.RevertNextPC()
		frame.Thread().InitClass(class)
		return
	}

	frame.Thread().InvokeMethod(instr.method)
}
