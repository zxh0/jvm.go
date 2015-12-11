package misc

import (
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	"github.com/zxh0/jvm.go/jvmgo/rtda/heap"
)

func init() {
	_urlcp(getLookupCacheURLs, "getLookupCacheURLs", "(Ljava/lang/ClassLoader;)[Ljava/net/URL;")
}

func _urlcp(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("sun/misc/URLClassPath", name, desc, method)
}

// private static native URL[] getLookupCacheURLs(ClassLoader var0);
// (Ljava/lang/ClassLoader;)[Ljava/net/URL;
func getLookupCacheURLs(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PushNull()
}
