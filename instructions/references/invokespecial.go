package references

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

// Invoke instance method;
// special handling for superclass, private, and instance initialization method invocations
type InvokeSpecial struct{ base.Index16Instruction }

func (instr *InvokeSpecial) Execute(frame *rtda.Frame) {
	cp := frame.GetConstantPool()
	k := cp.GetConstant(instr.Index)
	if kMethodRef, ok := k.(*heap.ConstantMethodRef); ok {
		method := kMethodRef.GetMethod(false)
		frame.Thread.InvokeMethod(method)
	} else {
		method := k.(*heap.ConstantInterfaceMethodRef).GetMethod(false)
		frame.Thread.InvokeMethod(method)
	}
}
