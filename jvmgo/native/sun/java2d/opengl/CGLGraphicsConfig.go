package awt

import (
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/rtda/class"
)

func init() {
	_cgl(cgl_initCGL, "initCGL", "()Z")
}

func _cgl(method interface{}, name, desc string) {
	rtc.RegisterNativeMethod("sun/java2d/opengl/CGLGraphicsConfig", name, desc, method)
}

func cgl_initCGL(frame *rtda.Frame) {
	//TODO
	frame.OperandStack().PushBoolean(true)
}
