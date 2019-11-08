package linker

import (
	"github.com/zxh0/jvm.go/rtda/heap"
	"github.com/zxh0/jvm.go/vm"
)

// https://docs.oracle.com/javase/specs/jvms/se13/html/jvms-5.html#jvms-5.4.6
func SelectMethod(obj *heap.Object, ref *heap.ConstantMethodRef) *heap.Method {
	if ref.ResolvedMethod.IsPrivate() {
		return ref.ResolvedMethod
	}
	if method := obj.Class.GetMethod(ref.Name, ref.Descriptor); method != nil {
		return method
	}
	// TODO: maximally-specific superinterface methods
	panic(vm.NewNoSuchMethodError(ref.String()))
}
