package native

import (
	"github.com/zxh0/jvm.go/rtda"
	//"github.com/zxh0/jvm.go/rtda/heap"
)

type NativeMethod func(frame *rtda.Frame)
