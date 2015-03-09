package zip

import (
	. "jvmgo/any"
	"jvmgo/jvm/rtda"
	rtc "jvmgo/jvm/rtda/class"
)

func init() {
	_zf(initIDs, "initIDs", "()V")
	_zf(getTotal, "getTotal", "(J)I")
	_zf(open, "open", "(Ljava/lang/String;IJZ)J")
	_zf(startsWithLOC, "startsWithLOC", "(J)Z")
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

// private static native boolean startsWithLOC(long jzfile);
// (J)Z
func startsWithLOC(frame *rtda.Frame) {
	// todo
	stack := frame.OperandStack()
	stack.PushBoolean(true)
}

// private static native int getTotal(long jzfile);
// (J)I
func getTotal(frame *rtda.Frame) {
	vars := frame.LocalVars()
	jzfile := vars.GetLong(0)

	total := getEntryCount(jzfile)

	stack := frame.OperandStack()
	stack.PushInt(total)
}
