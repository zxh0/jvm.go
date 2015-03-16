package lang

import (
	. "github.com/zxh0/jvm.go/jvmgo/any"
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
)

func init() {
	_package(getSystemPackage0, "getSystemPackage0", "(Ljava/lang/String;)Ljava/lang/String;")
}

func _package(method Any, name, desc string) {
	rtc.RegisterNativeMethod("java/lang/Package", name, desc, method)
}

// private static native String getSystemPackage0(String name);
// (Ljava/lang/String;)Ljava/lang/String;
func getSystemPackage0(frame *rtda.Frame) {
	// vars := frame.LocalVars()
	// name := vars.GetRef(0)

	sysPkg := frame.ClassLoader().JLObjectClass().LoadedFrom().String()
	sysPkgObj := rtda.NewJString(sysPkg)

	stack := frame.OperandStack()
	stack.PushRef(sysPkgObj)
}

// private static native String[] getSystemPackages0();
