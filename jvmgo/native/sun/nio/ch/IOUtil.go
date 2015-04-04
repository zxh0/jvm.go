package ch

import (
	. "github.com/zxh0/jvm.go/jvmgo/any"
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
)

func init() {
	_ioUtil(iou_iovMax, "iovMax", "()I")
	//sun/nio/ch/IOUtil~iovMax~
}

func _ioUtil(method Any, name, desc string) {
	rtc.RegisterNativeMethod("sun/nio/ch/IOUtil", name, desc, method)
}

func iou_iovMax(frame *rtda.Frame) {
	//read config _SC_IOV_MAX
	//default = 16
	frame.OperandStack().PushInt(16)
}
