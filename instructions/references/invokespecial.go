package references

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

// Invoke instance method;
// special handling for superclass, private, and instance initialization method invocations
type INVOKE_SPECIAL struct{ base.Index16Instruction }

func (instr *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	k := cp.GetConstant(instr.Index)
	if kMethodRef, ok := k.(*heap.ConstantMethodref); ok {
		method := kMethodRef.SpecialMethod()
		frame.Thread().InvokeMethod(method)
	} else {
		method := k.(*heap.ConstantInterfaceMethodref).SpecialMethod()
		frame.Thread().InvokeMethod(method)
	}
}
