package misc

import (
	. "github.com/zxh0/jvm.go/jvmgo/any"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
)

func init() {

}

func _unsafe(method Any, name, desc string) {
	rtc.RegisterNativeMethod("sun/misc/Unsafe", name, desc, method)
}
