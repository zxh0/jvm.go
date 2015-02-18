package lang

import (
	. "jvmgo/any"
	"jvmgo/jvm/rtda"
	rtc "jvmgo/jvm/rtda/class"
	"jvmgo/util"
)

func init() {
	_class(desiredAssertionStatus0, "desiredAssertionStatus0", "(Ljava/lang/Class;)Z")
	_class(forName0, "forName0", "(Ljava/lang/String;ZLjava/lang/ClassLoader;)Ljava/lang/Class;")
	_class(getClassLoader0, "getClassLoader0", "()Ljava/lang/ClassLoader;")
	_class(getComponentType, "getComponentType", "()Ljava/lang/Class;")
	_class(getDeclaredConstructors0, "getDeclaredConstructors0", "(Z)[Ljava/lang/reflect/Constructor;")
	_class(getDeclaredFields0, "getDeclaredFields0", "(Z)[Ljava/lang/reflect/Field;")
	_class(getDeclaredMethods0, "getDeclaredMethods0", "(Z)[Ljava/lang/reflect/Method;")
	_class(getInterfaces, "getInterfaces", "()[Ljava/lang/Class;")
	_class(getModifiers, "getModifiers", "()I")
	_class(getName0, "getName0", "()Ljava/lang/String;")
	_class(getPrimitiveClass, "getPrimitiveClass", "(Ljava/lang/String;)Ljava/lang/Class;")
	_class(getSuperclass, "getSuperclass", "()Ljava/lang/Class;")
	_class(isArray, "isArray", "()Z")
	_class(isInterface, "isInterface", "()Z")
	_class(isPrimitive, "isPrimitive", "()Z")
}

func _class(method Any, name, desc string) {
	rtc.RegisterNativeMethod("java/lang/Class", name, desc, method)
}

// private native Class<?>[]   getDeclaredClasses0();

// private static native boolean desiredAssertionStatus0(Class<?> clazz);
// (Ljava/lang/Class;)Z
func desiredAssertionStatus0(frame *rtda.Frame) {
	// todo
	stack := frame.OperandStack()
	//stack.PopRef() // this
	stack.PushBoolean(false)
}

// private static native Class<?> forName0(String name, boolean initialize, ClassLoader loader) throws ClassNotFoundException;
// (Ljava/lang/String;ZLjava/lang/ClassLoader;)Ljava/lang/Class;
func forName0(frame *rtda.Frame) {
	vars := frame.LocalVars()
	jName := vars.GetRef(0)
	initialize := vars.GetBoolean(1)
	//jLoader := vars.GetRef(2)

	goName := rtda.GoString(jName)
	goName = util.ReplaceAll(goName, ".", "/")
	goClass := frame.Method().Class().ClassLoader().LoadClass(goName)
	jClass := goClass.JClass()

	if initialize && goClass.InitializationNotStarted() {
		// undo forName0
		thread := frame.Thread()
		frame.SetNextPC(thread.PC())
		// init class
		thread.InitClass(goClass)
	} else {
		stack := frame.OperandStack()
		stack.PushRef(jClass)
	}
}

// native ClassLoader getClassLoader0();
// ()Ljava/lang/ClassLoader;
func getClassLoader0(frame *rtda.Frame) {
	// todo
	// _ = stack.PopRef() // this
	stack := frame.OperandStack()
	stack.PushRef(nil)
}

// public native Class<?> getComponentType();
// ()Ljava/lang/Class;
func getComponentType(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()

	class := this.Extra().(*rtc.Class)
	componentClass := class.ComponentClass()
	componentClassObj := componentClass.JClass()

	stack := frame.OperandStack()
	stack.PushRef(componentClassObj)
}

// private native Class<?>[] getInterfaces();
// ()[Ljava/lang/Class;
func getInterfaces(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()

	class := this.Extra().(*rtc.Class)
	interfaces := class.Interfaces()
	interfaceObjs := make([]*rtc.Obj, len(interfaces))
	for i, iface := range interfaces {
		interfaceObjs[i] = iface.JClass()
	}

	jlClassClass := class.ClassLoader().JLClassClass()
	interfaceArr := rtc.NewRefArray2(jlClassClass, interfaceObjs)

	stack := frame.OperandStack()
	stack.PushRef(interfaceArr)
}

// private native String getName0();
// ()Ljava/lang/String;
func getName0(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()

	class := this.Extra().(*rtc.Class)
	name := class.JlsName()
	nameObj := rtda.NewJString(name, frame)

	stack := frame.OperandStack()
	stack.PushRef(nameObj)
}

// public native int getModifiers();
// ()I
func getModifiers(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()

	class := this.Extra().(*rtc.Class)
	modifiers := class.GetAccessFlags()

	stack := frame.OperandStack()
	stack.PushInt(int32(modifiers))
}

// static native Class<?> getPrimitiveClass(String name);
// (Ljava/lang/String;)Ljava/lang/Class;
func getPrimitiveClass(frame *rtda.Frame) {
	vars := frame.LocalVars()
	nameObj := vars.GetRef(0)

	name := rtda.GoString(nameObj)
	classLoader := frame.ClassLoader()
	class := classLoader.GetPrimitiveClass(name)
	classObj := class.JClass()

	stack := frame.OperandStack()
	stack.PushRef(classObj)
}

// public native Class<? super T> getSuperclass();
// ()Ljava/lang/Class;
func getSuperclass(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()

	class := this.Extra().(*rtc.Class)
	superClass := class.SuperClass()

	stack := frame.OperandStack()
	if superClass != nil {
		stack.PushRef(superClass.JClass())
	} else {
		stack.PushNull()
	}
}

// public native boolean isArray();
// ()Z
func isArray(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()

	class := this.Extra().(*rtc.Class)
	stack := frame.OperandStack()
	stack.PushBoolean(class.IsArray())
}

// public native boolean isInterface();
// ()Z
func isInterface(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()

	class := this.Extra().(*rtc.Class)
	stack := frame.OperandStack()
	stack.PushBoolean(class.IsInterface())
}

// public native boolean isPrimitive();
// ()Z
func isPrimitive(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()

	class := this.Extra().(*rtc.Class)
	stack := frame.OperandStack()
	stack.PushBoolean(class.IsPrimitive())
}
