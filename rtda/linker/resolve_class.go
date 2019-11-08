package linker

import (
	"github.com/zxh0/jvm.go/rtda/heap"
)

// https://docs.oracle.com/javase/specs/jvms/se13/html/jvms-5.html#jvms-5.4.3.1
func ResolveClass(loader *heap.ClassLoader, ref *heap.ConstantClass) *heap.Class {
	if ref.Resolved == nil {
		ref.Resolved = loader.LoadClass(ref.Name) // TODO
	}
	return ref.Resolved
}
