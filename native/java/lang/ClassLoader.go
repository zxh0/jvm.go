package lang

import (
	"github.com/zxh0/jvm.go/native"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
	"github.com/zxh0/jvm.go/vmutils"
)

func init() {
	_cl(defineClass1, "defineClass1", "(Ljava/lang/String;[BIILjava/security/ProtectionDomain;Ljava/lang/String;)Ljava/lang/Class;")
	_cl(findBootstrapClass, "findBootstrapClass", "(Ljava/lang/String;)Ljava/lang/Class;")
	_cl(findLoadedClass0, "findLoadedClass0", "(Ljava/lang/String;)Ljava/lang/Class;")
}

func _cl(method native.Method, name, desc string) {
	native.Register("java/lang/ClassLoader", name, desc, method)
}

// private native Class<?> defineClass0(String name, byte[] b, int off, int len,
//                                      ProtectionDomain pd);

// private native Class<?> defineClass1(String name, byte[] b, int off, int len,
//                                      ProtectionDomain pd, String source);
// (Ljava/lang/String;[BIILjava/security/ProtectionDomain;Ljava/lang/String;)Ljava/lang/Class;
func defineClass1(frame *rtda.Frame) {
	//this := frame.GetThis()
	name := frame.GetRefVar(1)
	byteArr := frame.GetRefVar(2)
	off := frame.GetIntVar(3)
	_len := frame.GetIntVar(4)
	// pd := frame.GetRefVar(5)
	// source := frame.GetRefVar(6)

	goBytes := byteArr.GetGoBytes()
	goBytes = goBytes[off : off+_len]

	//println(this.Extra)
	panic(name.JSToGoStr())

}

// private native Class<?> defineClass2(String name, java.nio.ByteBuffer b,
//                                      int off, int len, ProtectionDomain pd,
//                                      String source);

// private native Class<?> findBootstrapClass(String name);
// (Ljava/lang/String;)Ljava/lang/Class;
func findBootstrapClass(frame *rtda.Frame) {
	//this := frame.GetThis()
	name := frame.GetRefVar(1)

	className := vmutils.DotToSlash(name.JSToGoStr())
	class := frame.GetBootLoader().LoadClass(className)

	// todo: init class?
	frame.PushRef(class.JClass)

	// todo
	if r := recover(); r != nil {
		frame.PushRef(nil)
	}
}

// private native final Class<?> findLoadedClass0(String name);
// (Ljava/lang/String;)Ljava/lang/Class;
func findLoadedClass0(frame *rtda.Frame) {
	this := frame.GetThis()
	name := frame.GetRefVar(1)

	className := vmutils.DotToSlash(name.JSToGoStr())

	if isAppClassLoader(this) {
		class := frame.GetBootLoader().FindLoadedClass(className)
		if class != nil {
			frame.PushRef(class.JClass)
		} else {
			frame.PushRef(nil)
		}
		return
	}

	// todo
	frame.PushRef(nil)
}

// todo
func isAppClassLoader(loader *heap.Object) bool {
	return loader.Class.Name == "sun/misc/Launcher$AppClassLoader"
}
