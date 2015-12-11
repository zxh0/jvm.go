package native

import (
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	//"github.com/zxh0/jvm.go/jvmgo/rtda/heap"
)

type NativeMethod func(frame *rtda.Frame)
