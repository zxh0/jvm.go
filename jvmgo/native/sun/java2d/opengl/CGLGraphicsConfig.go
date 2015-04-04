package awt

import (
	. "github.com/zxh0/jvm.go/jvmgo/any"
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
)

func init() {
	_cgl(cgl_initCGL, "initCGL", "()Z")
}

func _cgl(method Any, name, desc string) {
	rtc.RegisterNativeMethod("sun/java2d/opengl/CGLGraphicsConfig", name, desc, method)
}

func cgl_initCGL(frame *rtda.Frame) {
	//TODO
	frame.OperandStack().PushBoolean(true)
}
