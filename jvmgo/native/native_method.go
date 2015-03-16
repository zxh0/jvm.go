package native

import (
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	//rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
)

type NativeMethod func(frame *rtda.Frame)
