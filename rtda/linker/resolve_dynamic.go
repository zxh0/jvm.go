package linker

import (
	"github.com/zxh0/jvm.go/classfile"
	"github.com/zxh0/jvm.go/rtda/heap"
)

func ResolveInvokeDynamic(loader *heap.ClassLoader, ref *heap.ConstantInvokeDynamic) {
	if ref.FieldRef != nil {
		isStatic := ref.RefKind == classfile.RefGetStatic || ref.RefKind == classfile.RefPutStatic
		ResolveField(loader, ref.FieldRef, isStatic)
	} else if ref.MethodRef != nil {
		ResolveMethod(loader, ref.MethodRef)
	}
}
