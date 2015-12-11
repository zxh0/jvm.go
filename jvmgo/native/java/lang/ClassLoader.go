package lang

import (
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	"github.com/zxh0/jvm.go/jvmgo/rtda/heap"
)

func init() {
	_cl(defineClass1, "defineClass1", "(Ljava/lang/String;[BIILjava/security/ProtectionDomain;Ljava/lang/String;)Ljava/lang/Class;")
	_cl(findBootstrapClass, "findBootstrapClass", "(Ljava/lang/String;)Ljava/lang/Class;")
	_cl(findLoadedClass0, "findLoadedClass0", "(Ljava/lang/String;)Ljava/lang/Class;")
}

func _cl(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("java/lang/ClassLoader", name, desc, method)
}

// private native Class<?> defineClass0(String name, byte[] b, int off, int len,
//                                      ProtectionDomain pd);

// private native Class<?> defineClass1(String name, byte[] b, int off, int len,
//                                      ProtectionDomain pd, String source);
// (Ljava/lang/String;[BIILjava/security/ProtectionDomain;Ljava/lang/String;)Ljava/lang/Class;
func defineClass1(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()
	name := vars.GetRef(1)
	byteArr := vars.GetRef(2)
	off := vars.GetInt(3)
	_len := vars.GetInt(4)
	// pd := vars.GetRef(5)
	// source := vars.GetRef(6)

	goBytes := byteArr.GoBytes()
	goBytes = goBytes[off : off+_len]

	println(this.Extra())
	panic(rtda.GoString(name))

}

// private native Class<?> defineClass2(String name, java.nio.ByteBuffer b,
//                                      int off, int len, ProtectionDomain pd,
//                                      String source);

// private native Class<?> findBootstrapClass(String name);
// (Ljava/lang/String;)Ljava/lang/Class;
func findBootstrapClass(frame *rtda.Frame) {
	vars := frame.LocalVars()
	//this := vars.GetThis()
	name := vars.GetRef(1)

	className := heap.DotToSlash(rtda.GoString(name))
	class := heap.BootLoader().LoadClass(className)

	// todo: init class?
	stack := frame.OperandStack()
	stack.PushRef(class.JClass())

	// todo
	if r := recover(); r != nil {
		frame.OperandStack().PushRef(nil)
	}
}

// private native final Class<?> findLoadedClass0(String name);
// (Ljava/lang/String;)Ljava/lang/Class;
func findLoadedClass0(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()
	name := vars.GetRef(1)

	className := heap.DotToSlash(rtda.GoString(name))

	if isAppClassLoader(this) {
		class := heap.BootLoader().FindLoadedClass(className)
		if class != nil {
			frame.OperandStack().PushRef(class.JClass())
		} else {
			frame.OperandStack().PushRef(nil)
		}
		return
	}

	// todo
	frame.OperandStack().PushRef(nil)
}

// todo
func isAppClassLoader(loader *heap.Object) bool {
	return loader.Class().Name() == "sun/misc/Launcher$AppClassLoader"
}
