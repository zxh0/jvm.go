package ch

import (
	"github.com/zxh0/jvm.go/native"
	"github.com/zxh0/jvm.go/rtda"
)

func init() {
	_ioUtil(iou_iovMax, "iovMax", "()I")
	_ioUtil(iou_setfdVal, "setfdVal", "(Ljava/io/FileDescriptor;I)V")
	_ioUtil(iou_fdVal, "fdVal", "(Ljava/io/FileDescriptor;)I")
	_ioUtil(iou_makePipe, "makePipe", "(Z)J")
}

func _ioUtil(method native.Method, name, desc string) {
	native.Register("sun/nio/ch/IOUtil", name, desc, method)
}

func iou_iovMax(frame *rtda.Frame) {
	//read config _SC_IOV_MAX
	//default = 16
	frame.PushInt(16)
}

//static native void setfdVal(FileDescriptor var0, int var1)
// (Ljava/io/FileDescriptor;I)V
func iou_setfdVal(frame *rtda.Frame) {
	//vars := frame.
	//fdVal := frame.GetIntVar(1)
	//fd := frame.GetRefVar(0)
	//fmt.Println(fd)
	//fmt.Println(fdVal)
}

//static native int fdVal(FileDescriptor fd);
//(Ljava/io/FileDescriptor;)I
func iou_fdVal(frame *rtda.Frame) {
	frame.PushInt(100)
}

//makePipe~(Z)J
func iou_makePipe(frame *rtda.Frame) {

}
