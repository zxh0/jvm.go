package ch

import (
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	"github.com/zxh0/jvm.go/jvmgo/rtda/heap"
)

func init() {
	_ioUtil(iou_iovMax, "iovMax", "()I")
	_ioUtil(iou_setfdVal, "setfdVal", "(Ljava/io/FileDescriptor;I)V")
	_ioUtil(iou_fdVal, "fdVal", "(Ljava/io/FileDescriptor;)I")
	_ioUtil(iou_makePipe, "makePipe", "(Z)J")
}

func _ioUtil(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("sun/nio/ch/IOUtil", name, desc, method)
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

//makePipe~(Z)J
func iou_makePipe(frame *rtda.Frame) {

}
