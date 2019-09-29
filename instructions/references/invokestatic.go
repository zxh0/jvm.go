package references

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

// Invoke a class (static) method
type InvokeStatic struct {
	base.Index16Instruction
	method *heap.Method
}

func (instr *InvokeStatic) Execute(frame *rtda.Frame) {
	if instr.method == nil {
		cp := frame.GetConstantPool()
		k := cp.GetConstant(instr.Index)
		if kMethodRef, ok := k.(*heap.ConstantMethodRef); ok {
			instr.method = kMethodRef.StaticMethod()
		} else {
			instr.method = k.(*heap.ConstantInterfaceMethodRef).StaticMethod()
		}
	}

	// init class
	class := instr.method.Class
	if class.InitializationNotStarted() {
		frame.RevertNextPC()
		frame.Thread().InitClass(class)
		return
	}

	frame.Thread().InvokeMethod(instr.method)
}
