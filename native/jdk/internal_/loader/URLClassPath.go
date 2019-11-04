package loader

import (
	"github.com/zxh0/jvm.go/native"
	"github.com/zxh0/jvm.go/rtda"
)

func init() {
	native.ForClass("jdk/internal/loader/URLClassPath").
		Register(getLookupCacheURLs, "(Ljava/lang/ClassLoader;)[Ljava/net/URL;")
}

// private static native URL[] getLookupCacheURLs(ClassLoader var0);
// (Ljava/lang/ClassLoader;)[Ljava/net/URL;
func getLookupCacheURLs(frame *rtda.Frame) {
	frame.PushNull()
}
