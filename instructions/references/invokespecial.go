package references

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
	"github.com/zxh0/jvm.go/rtda/linker"
	"github.com/zxh0/jvm.go/vm"
)

// Invoke instance method;
// special handling for superclass, private, and instance initialization method invocations
type InvokeSpecial struct {
	base.Index16Instruction
	method *heap.Method
}

func (instr *InvokeSpecial) Execute(frame *rtda.Frame) {
	if instr.method == nil {
		methodRef := instr.resolveSymbol(frame)
		class := instr.getClass(frame, methodRef)
		method := instr.lookupMethod(frame, class, methodRef)
		if method.IsAbstract() {
			panic(vm.NewAbstractMethodError(method.String()))
		}
		instr.method = method
	}

	frame.Thread.InvokeMethod(instr.method)
}

func (instr *InvokeSpecial) resolveSymbol(frame *rtda.Frame) *heap.ConstantMethodRef {
	cp := frame.GetConstantPool()
	methodRef := cp.GetConstant(instr.Index).(*heap.ConstantMethodRef)
	method := linker.ResolveMethod(frame.GetBootLoader(), methodRef)
	if method.IsStatic() {
		panic(vm.NewIncompatibleClassChangeError(method.String()))
	}
	if method.IsAbstract() {
		panic(vm.NewAbstractMethodError(method.String()))
	}
	return methodRef
}

func (instr *InvokeSpecial) getClass(frame *rtda.Frame, methodRef *heap.ConstantMethodRef) *heap.Class {
	currentClass := frame.Method.Class
	if !methodRef.ResolvedMethod.IsConstructor() &&
		methodRef.ResolvedClass == currentClass.SuperClass &&
		methodRef.ResolvedClass.IsAccSuper() {
		return currentClass.SuperClass
	}
	return methodRef.ResolvedClass
}

func (instr *InvokeSpecial) lookupMethod(frame *rtda.Frame,
	class *heap.Class, methodRef *heap.ConstantMethodRef) *heap.Method {

	if method := class.GetMethod(methodRef.Name, methodRef.Descriptor); method != nil {
		return method
	}
	if class.IsInterface() {
		objClass := frame.GetBootLoader().JLObjectClass()
		if method := objClass.GetDeclaredMethod(methodRef.Name, methodRef.Descriptor); method != nil {
			if method.IsPublic() {
				return method
			}
		}
	}
	// TODO: maximally-specific superinterface method
	panic(vm.NewNoSuchMethodError(methodRef.String()))
}
