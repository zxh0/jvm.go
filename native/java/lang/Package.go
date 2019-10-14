package lang

import (
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
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
	// vars := frame.
	// name := frame.GetRefVar(0)

	sysPkg := frame.GetClassLoader().JLObjectClass().LoadedFrom.String()
	sysPkgObj := heap.JSFromGoStr(sysPkg)

	frame.PushRef(sysPkgObj)
}

// private static native String[] getSystemPackages0();
