package native

import (
	. "jvmgo/any"
	"jvmgo/jvm/rtda"
	rtc "jvmgo/jvm/rtda/class"
	_ "jvmgo/native/java/io"
	_ "jvmgo/native/java/lang"
	_ "jvmgo/native/java/lang/reflect"
	_ "jvmgo/native/java/security"
	_ "jvmgo/native/java/util"
	_ "jvmgo/native/java/util/concurrent/atomic"
	_ "jvmgo/native/java/util/zip"
	_ "jvmgo/native/sun/misc"
	_ "jvmgo/native/sun/reflect"
)

// register native methods
func init() {
	rtc.SetRegisterNatives(registerNatives)
}

func registerNatives(frame *rtda.Frame) {
	// todo
}
