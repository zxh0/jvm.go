package reflect

import (
	. "jvmgo/any"
	"jvmgo/jvm/rtda"
	rtc "jvmgo/jvm/rtda/class"
)

func init() {
	_cp(getUTF8At0, "getUTF8At0", "(Ljava/lang/Object;I)Ljava/lang/String;")
}

func _cp(method Any, name, desc string) {
	rtc.RegisterNativeMethod("sun/reflect/ConstantPool", name, desc, method)
}

// private native String getUTF8At0(Object o, int i);
// (Ljava/lang/Object;I)Ljava/lang/String;
func getUTF8At0(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()
	i := uint(vars.GetInt(2))

	cp := this.Extra().(*rtc.ConstantPool)
	kUtf8 := cp.GetConstant(i).(*rtc.ConstantUtf8)
	goStr := kUtf8.Str()
	jStr := rtda.NewJString(goStr, frame)
	frame.OperandStack().PushRef(jStr)
}
