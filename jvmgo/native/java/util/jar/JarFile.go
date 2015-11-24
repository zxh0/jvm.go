package jar

import (
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
)

func init() {
	_jf(getMetaInfEntryNames, "getMetaInfEntryNames", "()[Ljava/lang/String;")
}

func _jf(method interface{}, name, desc string) {
	rtc.RegisterNativeMethod("java/util/jar/JarFile", name, desc, method)
}

// private native String[] getMetaInfEntryNames();
// ()[Ljava/lang/String;
func getMetaInfEntryNames(frame *rtda.Frame) {
	// todo
	stack := frame.OperandStack()
	stack.PushNull()
}
