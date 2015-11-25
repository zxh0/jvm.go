package native

import (
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	//rtc "github.com/zxh0/jvm.go/jvmgo/rtda/class"
)

type NativeMethod func(frame *rtda.Frame)
