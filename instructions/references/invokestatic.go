package references

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
	"github.com/zxh0/jvm.go/rtda/linker"
	"github.com/zxh0/jvm.go/vm"
)

// Invoke a class (static) method
type InvokeStatic struct {
	base.Index16Instruction
	method *heap.Method
}

func (instr *InvokeStatic) Execute(frame *rtda.Frame) {
	if instr.method == nil {
		cp := frame.GetConstantPool()
		methodRef := cp.GetConstant(instr.Index).(*heap.ConstantMethodRef)
		method := linker.ResolveMethod(frame.GetBootLoader(), methodRef)
		if !method.IsStatic() || method.IsClinit() || method.IsConstructor() {
			panic(vm.NewIncompatibleClassChangeError(method.String()))
		}
		instr.method = method
	}

	// init class
	class := instr.method.Class
	if class.InitializationNotStarted() {
		frame.RevertNextPC()
		frame.Thread.InitClass(class)
		return
	}

	frame.Thread.InvokeMethod(instr.method)
}
