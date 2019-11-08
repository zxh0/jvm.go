package lang

import (
	"github.com/zxh0/jvm.go/native"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/vmutils"
)

func init() {
	native.ForClass("java/lang/StringUTF16").
		Register(isBigEndian, "()Z")
}

// private static native boolean isBigEndian();
func isBigEndian(frame *rtda.Frame) {
	frame.PushBoolean(vmutils.IsBigEndian())
}
