package reflect

import (
	. "jvmgo/any"
	"jvmgo/jvm/rtda"
	rtc "jvmgo/jvm/rtda/class"
)

func init() {
	_cp(getLongAt0, "getLongAt0", "(Ljava/lang/Object;I)J")
	_cp(getUTF8At0, "getUTF8At0", "(Ljava/lang/Object;I)Ljava/lang/String;")
}

func _cp(method Any, name, desc string) {
	rtc.RegisterNativeMethod("sun/reflect/ConstantPool", name, desc, method)
}

// private native long getLongAt0(Object o, int i);
// (Ljava/lang/Object;I)J
func getLongAt0(frame *rtda.Frame) {
	cp, index := _getPop(frame)
	val := cp.GetConstant(index).(int64)
	frame.OperandStack().PushLong(val)
}

// private native String getUTF8At0(Object o, int i);
// (Ljava/lang/Object;I)Ljava/lang/String;
func getUTF8At0(frame *rtda.Frame) {
	cp, index := _getPop(frame)
	kUtf8 := cp.GetConstant(index).(*rtc.ConstantUtf8)
	goStr := kUtf8.Str()
	jStr := rtda.NewJString(goStr)
	frame.OperandStack().PushRef(jStr)
}

func _getPop(frame *rtda.Frame) (cp *rtc.ConstantPool, index uint) {
	vars := frame.LocalVars()
	this := vars.GetThis()
	index = uint(vars.GetInt(2))
	cp = this.Extra().(*rtc.ConstantPool)
	return
}
