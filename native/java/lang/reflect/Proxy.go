package reflect

import (
	"github.com/zxh0/jvm.go/native"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/vmutils"
)

func init() {
	_proxy(defineClass0, "defineClass0", "(Ljava/lang/ClassLoader;Ljava/lang/String;[BII)Ljava/lang/Class;")
}

func _proxy(method native.Method, name, desc string) {
	native.Register("java/lang/reflect/Proxy", name, desc, method)
}

// private static native Class<?> defineClass0(ClassLoader loader, String name,
//                                             byte[] b, int off, int len);
// (Ljava/lang/ClassLoader;Ljava/lang/String;[BII)Ljava/lang/Class;
func defineClass0(frame *rtda.Frame) {
	if frame.IsStackEmpty() {
		_loadClass(frame)
	}

	// init class
	class := frame.TopRef(0).GetGoClass()
	if class.InitializationNotStarted() {
		frame.RevertNextPC()
		frame.Thread.InitClass(class)
	}
}

func _loadClass(frame *rtda.Frame) {
	//loader := frame.GetRefVar(0)
	nameObj := frame.GetRefVar(1)
	byteArr := frame.GetRefVar(2)
	off := frame.GetIntVar(3)
	_len := frame.GetIntVar(4)

	name := nameObj.JSToGoStr()
	name = vmutils.DotToSlash(name)
	data := byteArr.GetGoBytes()
	data = data[off : off+_len]

	// todo
	class := frame.GetClassLoader().DefineClass(name, data)
	frame.PushRef(class.JClass)
}
