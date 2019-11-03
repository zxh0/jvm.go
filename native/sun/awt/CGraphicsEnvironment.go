package awt

import (
	"github.com/zxh0/jvm.go/native"
	"github.com/zxh0/jvm.go/rtda"
)

func init() {
	_cge(cge_initCocoa, "initCocoa", "()V")
	_cge(cge_getMainDisplayID, "getMainDisplayID", "()I")
}

func _cge(method native.Method, name, desc string) {
	native.Register("sun/awt/CGraphicsEnvironment", name, desc, method)
}

func cge_initCocoa(frame *rtda.Frame) {
	//TODO
}

func cge_getMainDisplayID(frame *rtda.Frame) {
	frame.PushInt(1)
}
