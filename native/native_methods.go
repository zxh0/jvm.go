package native

import (
	_ "github.com/zxh0/jvm.go/native/java/awt"
	_ "github.com/zxh0/jvm.go/native/java/io"
	_ "github.com/zxh0/jvm.go/native/java/lang"
	_ "github.com/zxh0/jvm.go/native/java/lang/invoke"
	_ "github.com/zxh0/jvm.go/native/java/lang/ref"
	_ "github.com/zxh0/jvm.go/native/java/lang/reflect"
	_ "github.com/zxh0/jvm.go/native/java/net"
	_ "github.com/zxh0/jvm.go/native/java/security"
	_ "github.com/zxh0/jvm.go/native/java/util"
	_ "github.com/zxh0/jvm.go/native/java/util/concurrent/atomic"
	_ "github.com/zxh0/jvm.go/native/java/util/jar"
	_ "github.com/zxh0/jvm.go/native/java/util/zip"
	_ "github.com/zxh0/jvm.go/native/jdk/internal_/loader"
	_ "github.com/zxh0/jvm.go/native/jdk/internal_/misc"
	_ "github.com/zxh0/jvm.go/native/jdk/internal_/reflect"
	_ "github.com/zxh0/jvm.go/native/jdk/internal_/util"
	_ "github.com/zxh0/jvm.go/native/sun/awt"
	_ "github.com/zxh0/jvm.go/native/sun/io"
	_ "github.com/zxh0/jvm.go/native/sun/java2d/opengl"
	_ "github.com/zxh0/jvm.go/native/sun/management"
	_ "github.com/zxh0/jvm.go/native/sun/nio/ch"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

// type NativeMethod func(frame *rtda.Frame)

func init() {
	heap.SetEmptyNativeMethod(emptyNativeMethod)
}

func emptyNativeMethod(frame *rtda.Frame) {
	// do nothing
}
