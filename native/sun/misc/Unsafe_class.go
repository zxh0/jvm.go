package misc

import (
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
	"github.com/zxh0/jvm.go/vmutils"
)

func init() {
	_unsafe(allocateInstance, "allocateInstance", "(Ljava/lang/Class;)Ljava/lang/Object;")
	_unsafe(defineClass, "defineClass", "(Ljava/lang/String;[BIILjava/lang/ClassLoader;Ljava/security/ProtectionDomain;)Ljava/lang/Class;")
	_unsafe(shouldBeInitialized, "shouldBeInitialized", "(Ljava/lang/Class;)Z")
	_unsafe(ensureClassInitialized, "ensureClassInitialized", "(Ljava/lang/Class;)V")
	_unsafe(staticFieldOffset, "staticFieldOffset", "(Ljava/lang/reflect/Field;)J")
	_unsafe(staticFieldBase, "staticFieldBase", "(Ljava/lang/reflect/Field;)Ljava/lang/Object;")
}

// public native Object allocateInstance(Class<?> type) throws InstantiationException;
// (Ljava/lang/Class;)Ljava/lang/Object;
func allocateInstance(frame *rtda.Frame) {
	classObj := frame.GetRefVar(1)

	class := classObj.GetGoClass()
	obj := class.NewObj()

	frame.PushRef(obj)
}

// public native Class defineClass(String name, byte[] b, int off, int len,
//  		ClassLoader loader, ProtectionDomain protectionDomain)
// (Ljava/lang/String;[BIILjava/lang/ClassLoader;Ljava/security/ProtectionDomain;)Ljava/lang/Class;
func defineClass(frame *rtda.Frame) {
	nameObj := frame.GetRefVar(1)
	byteArr := frame.GetRefVar(2)
	off := frame.GetIntVar(3)
	_len := frame.GetIntVar(4)
	//loaderObj := frame.GetRefVar(5)
	//pd := frame.GetRefVar(6)

	name := nameObj.JSToGoStr()
	name = vmutils.DotToSlash(name)
	data := byteArr.GetGoBytes()
	data = data[off : off+_len]

	// todo
	class := frame.GetClassLoader().DefineClass(name, data)
	frame.PushRef(class.JClass)
}

// public native boolean shouldBeInitialized(Class<?> c);
// (Ljava/lang/Class;)V
func shouldBeInitialized(frame *rtda.Frame) {
	// this := frame.GetRefVar(0)
	classObj := frame.GetRefVar(1)

	goClass := classObj.GetGoClass()
	ret := goClass.InitializationNotStarted() // TODO
	frame.PushBoolean(ret)
}

// public native void ensureClassInitialized(Class<?> c);
// (Ljava/lang/Class;)V
func ensureClassInitialized(frame *rtda.Frame) {
	// this := frame.GetRefVar(0)
	classObj := frame.GetRefVar(1)

	goClass := classObj.GetGoClass()
	if goClass.InitializationNotStarted() {
		// undo ensureClassInitialized()
		frame.RevertNextPC()
		// init
		frame.Thread.InitClass(goClass)
	}
}

// public native long staticFieldOffset(Field f);
// (Ljava/lang/reflect/Field;)J
func staticFieldOffset(frame *rtda.Frame) {
	// frame.GetRefVar(0) // this
	fieldObj := frame.GetRefVar(1)

	offset := fieldObj.GetFieldValue("slot", "I").IntValue()
	frame.PushLong(int64(offset))
}

// public native Object staticFieldBase(Field f);
// (Ljava/lang/reflect/Field;)Ljava/lang/Object;
func staticFieldBase(frame *rtda.Frame) {
	// frame.GetRefVar(0) // this
	fieldObj := frame.GetRefVar(1)

	goField := _getGoField(fieldObj)
	obj := goField.Class.AsObj()

	frame.PushRef(obj)
}

func _getGoField(fieldObj *heap.Object) *heap.Field {
	extra := fieldObj.Extra
	if extra != nil {
		return extra.(*heap.Field)
	}

	root := fieldObj.GetFieldValue("root", "Ljava/lang/reflect/Field;").Ref
	return root.Extra.(*heap.Field)
}
