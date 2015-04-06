package invoke

import (
	. "github.com/zxh0/jvm.go/jvmgo/any"
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
)

func init() {
	_mhn(getConstant, "getConstant", "(I)I")
}

func _mhn(method Any, name, desc string) {
	rtc.RegisterNativeMethod("java/lang/invoke/MethodHandleNatives", name, desc, method)
}

// static native int getConstant(int which);
// (I)I
func getConstant(frame *rtda.Frame) {
	vars := frame.LocalVars()
	which := vars.GetInt(0)

	if which == 4 {
		frame.OperandStack().PushInt(1)
	} else {
		frame.OperandStack().PushInt(0)
	}
}
