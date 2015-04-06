package native

import (
	. "github.com/zxh0/jvm.go/jvmgo/any"
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
	_ "github.com/zxh0/jvm.go/jvmgo/native/java/awt"
	_ "github.com/zxh0/jvm.go/jvmgo/native/java/io"
	_ "github.com/zxh0/jvm.go/jvmgo/native/java/lang"
	_ "github.com/zxh0/jvm.go/jvmgo/native/java/lang/invoke"
	_ "github.com/zxh0/jvm.go/jvmgo/native/java/lang/reflect"
	_ "github.com/zxh0/jvm.go/jvmgo/native/java/net"
	_ "github.com/zxh0/jvm.go/jvmgo/native/java/security"
	_ "github.com/zxh0/jvm.go/jvmgo/native/java/util"
	_ "github.com/zxh0/jvm.go/jvmgo/native/java/util/concurrent/atomic"
	_ "github.com/zxh0/jvm.go/jvmgo/native/java/util/jar"
	_ "github.com/zxh0/jvm.go/jvmgo/native/java/util/zip"
	_ "github.com/zxh0/jvm.go/jvmgo/native/sun/awt"
	_ "github.com/zxh0/jvm.go/jvmgo/native/sun/java2d/opengl"
	_ "github.com/zxh0/jvm.go/jvmgo/native/sun/management"
	_ "github.com/zxh0/jvm.go/jvmgo/native/sun/misc"
	_ "github.com/zxh0/jvm.go/jvmgo/native/sun/reflect"
)

func init() {
	rtc.SetEmptyNativeMethod(emptyNativeMethod)
}

func emptyNativeMethod(frame *rtda.Frame) {
	// do nothing
}
