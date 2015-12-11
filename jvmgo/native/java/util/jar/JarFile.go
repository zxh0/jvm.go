package jar

import (
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	"github.com/zxh0/jvm.go/jvmgo/rtda/heap"
)

func init() {
	_jf(getMetaInfEntryNames, "getMetaInfEntryNames", "()[Ljava/lang/String;")
}

func _jf(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("java/util/jar/JarFile", name, desc, method)
}

// private native String[] getMetaInfEntryNames();
// ()[Ljava/lang/String;
func getMetaInfEntryNames(frame *rtda.Frame) {
	// todo
	stack := frame.OperandStack()
	stack.PushNull()
}
