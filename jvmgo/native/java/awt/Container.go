package awt

import (
	. "github.com/zxh0/jvm.go/jvmgo/any"
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
)

func init() {
	_container(container_initIDs, "initIDs", "()V")
}

func _container(method Any, name, desc string) {
	rtc.RegisterNativeMethod("java/awt/Container", name, desc, method)
}

func container_initIDs(frame *rtda.Frame) {
	//TODO
}
