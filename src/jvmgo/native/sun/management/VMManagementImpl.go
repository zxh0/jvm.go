package management

import (
	. "jvmgo/any"
	"jvmgo/jvm/rtda"
	rtc "jvmgo/jvm/rtda/class"
)

func init() {
	_vmm(getStartupTime, "getStartupTime", "()J")
	_vmm(getVersion0, "getVersion0", "()Ljava/lang/String;")
	_vmm(initOptionalSupportFields, "initOptionalSupportFields", "()V")
}

func _vmm(method Any, name, desc string) {
	rtc.RegisterNativeMethod("sun/management/VMManagementImpl", name, desc, method)
}

// public native long getStartupTime();
// ()J
func getStartupTime(frame *rtda.Frame) {
	// todo
	startupTime := int64(0)

	stack := frame.OperandStack()
	stack.PushLong(startupTime)
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
