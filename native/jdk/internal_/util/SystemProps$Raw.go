package util

import (
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

func init() {
	sysPropsRaw(vmProperties, "vmProperties", "()[Ljava/lang/String;")
	sysPropsRaw(platformProperties, "platformProperties", "()[Ljava/lang/String;")
}

func sysPropsRaw(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("jdk/internal/util/SystemProps$Raw", name, desc, method)
}

// private static native String[] vmProperties();
func vmProperties(frame *rtda.Frame) {
	// TODO
	frame.PushRef(frame.GetRuntime().NewStringArray(nil))
}

// private static native String[] platformProperties();
func platformProperties(frame *rtda.Frame) {
	// TODO
	jStrs := make([]*heap.Object, 40)
	frame.PushRef(frame.GetRuntime().NewStringArray(jStrs))
}
