package management

import (
	. "jvmgo/any"
	"jvmgo/jvm/rtda"
	rtc "jvmgo/jvm/rtda/class"
)

func init() {
	_vmm(getVersion0, "getVersion0", "()Ljava/lang/String;")
	_vmm(initOptionalSupportFields, "initOptionalSupportFields", "()V")
}

func _vmm(method Any, name, desc string) {
	rtc.RegisterNativeMethod("sun/management/VMManagementImpl", name, desc, method)
}

// private static native String getVersion0();
// ()Ljava/lang/String;
func getVersion0(frame *rtda.Frame) {
	// todo
	version := rtda.NewJString("0", frame)

	stack := frame.OperandStack()
	stack.PushRef(version)
}

// private static native void initOptionalSupportFields();
// ()V
func initOptionalSupportFields(frame *rtda.Frame) {
	// todo
}
