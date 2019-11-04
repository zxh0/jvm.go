package lang

import (
	"github.com/zxh0/jvm.go/native"
	"github.com/zxh0/jvm.go/rtda"
)

func init() {
	native.ForClass("java/lang/String").
		Register(intern, "()Ljava/lang/String;")
}

// public native String intern();
// ()Ljava/lang/String;
func intern(frame *rtda.Frame) {
	jStr := frame.GetThis()

	goStr := jStr.JSToGoStr()
	internedStr := frame.GetRuntime().JSIntern(goStr, jStr)

	frame.PushRef(internedStr)
}
