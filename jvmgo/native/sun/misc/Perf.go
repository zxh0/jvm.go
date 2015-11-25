package misc

import (
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/rtda/class"
)

func init() {
	_perf(createLong, "createLong", "(Ljava/lang/String;IIJ)Ljava/nio/ByteBuffer;")
}

func _perf(method interface{}, name, desc string) {
	rtc.RegisterNativeMethod("sun/misc/Perf", name, desc, method)
}

// public native ByteBuffer createLong(String name, int variability, int units, long value);
// (Ljava/lang/String;IIJ)Ljava/nio/ByteBuffer;
func createLong(frame *rtda.Frame) {
	bbClass := frame.ClassLoader().LoadClass("java/nio/ByteBuffer")
	if bbClass.InitializationNotStarted() {
		frame.RevertNextPC()
		frame.Thread().InitClass(bbClass)
		return
	}

	stack := frame.OperandStack()
	stack.PushInt(8)

	allocate := bbClass.GetStaticMethod("allocate", "(I)Ljava/nio/ByteBuffer;")
	frame.Thread().InvokeMethod(allocate)
}
