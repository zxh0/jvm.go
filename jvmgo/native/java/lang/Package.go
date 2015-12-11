package lang

import (
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	"github.com/zxh0/jvm.go/jvmgo/rtda/heap"
)

func init() {
	_package(getSystemPackage0, "getSystemPackage0", "(Ljava/lang/String;)Ljava/lang/String;")
}

func _package(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("java/lang/Package", name, desc, method)
}

// private static native String getSystemPackage0(String name);
// (Ljava/lang/String;)Ljava/lang/String;
func getSystemPackage0(frame *rtda.Frame) {
	// vars := frame.LocalVars()
	// name := vars.GetRef(0)

	sysPkg := frame.ClassLoader().JLObjectClass().LoadedFrom().String()
	sysPkgObj := rtda.JString(sysPkg)

	stack := frame.OperandStack()
	stack.PushRef(sysPkgObj)
}

// private static native String[] getSystemPackages0();
