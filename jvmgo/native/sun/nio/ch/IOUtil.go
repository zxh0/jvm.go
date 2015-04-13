package ch

import (
	. "github.com/zxh0/jvm.go/jvmgo/any"
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
)

func init() {
	_ioUtil(iou_iovMax, "iovMax", "()I")
	_ioUtil(iou_setfdVal, "setfdVal", "(Ljava/io/FileDescriptor;I)V")
	_ioUtil(iou_fdVal, "fdVal", "(Ljava/io/FileDescriptor;)I")
}

func _ioUtil(method Any, name, desc string) {
	rtc.RegisterNativeMethod("sun/nio/ch/IOUtil", name, desc, method)
}

func iou_iovMax(frame *rtda.Frame) {
	//read config _SC_IOV_MAX
	//default = 16
	frame.OperandStack().PushInt(16)
}

//static native void setfdVal(FileDescriptor var0, int var1)
// (Ljava/io/FileDescriptor;I)V
func iou_setfdVal(frame *rtda.Frame) {
	//vars := frame.LocalVars()
	//fdVal := vars.GetInt(1)
	//fd := vars.GetRef(0)
	//fmt.Println(fd)
	//fmt.Println(fdVal)
}

//static native int fdVal(FileDescriptor fd);
//(Ljava/io/FileDescriptor;)I
func iou_fdVal(frame *rtda.Frame) {
	frame.OperandStack().PushInt(100)
}
