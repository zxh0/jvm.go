package misc

import (
	"github.com/zxh0/jvm.go/native"
	"github.com/zxh0/jvm.go/rtda"
)

func init() {
	_urlcp(getLookupCacheURLs, "getLookupCacheURLs", "(Ljava/lang/ClassLoader;)[Ljava/net/URL;")
}

func _urlcp(method native.Method, name, desc string) {
	native.Register("sun/misc/URLClassPath", name, desc, method)
}

// private static native URL[] getLookupCacheURLs(ClassLoader var0);
// (Ljava/lang/ClassLoader;)[Ljava/net/URL;
func getLookupCacheURLs(frame *rtda.Frame) {
	frame.PushNull()
}
