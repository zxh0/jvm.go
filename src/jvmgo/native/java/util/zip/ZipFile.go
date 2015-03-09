package zip

import (
	. "jvmgo/any"
	"jvmgo/jvm/rtda"
	rtc "jvmgo/jvm/rtda/class"
)

func init() {
	_zf(initIDs, "initIDs", "()V")
	_zf(open, "open", "(Ljava/lang/String;IJZ)J")
}

func _zf(method Any, name, desc string) {
	rtc.RegisterNativeMethod("java/util/zip/ZipFile", name, desc, method)
}

// private static native void initIDs();
// ()V
func initIDs(frame *rtda.Frame) {
	// todo
}

// private static native long open(String name, int mode, long lastModified,
//                                 boolean usemmap) throws IOException;
// (Ljava/lang/String;IJZ)J
func open(frame *rtda.Frame) {
	vars := frame.LocalVars()
	nameObj := vars.GetRef(0)

	name := rtda.GoString(nameObj)
	jzfile, err := openZip(name)
	if err != nil {
		// todo
		panic("IOException")
	}

	stack := frame.OperandStack()
	stack.PushLong(jzfile)
}
