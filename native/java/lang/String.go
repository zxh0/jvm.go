package lang

import (
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

func init() {
	_string(intern, "intern", "()Ljava/lang/String;")
}

func _string(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("java/lang/String", name, desc, method)
}

// public native String intern();
// ()Ljava/lang/String;
func intern(frame *rtda.Frame) {
	jStr := frame.GetThis()

	goStr := jStr.JSToGoStr()
	internedStr := frame.GetRuntime().JSIntern(goStr, jStr)

	frame.PushRef(internedStr)
}
