package lang

import (
	"strings"

	cp "github.com/zxh0/jvm.go/jvmgo/classpath"
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	"github.com/zxh0/jvm.go/jvmgo/rtda/heap"
)

func init() {
	_class(getClassLoader0, "getClassLoader0", "()Ljava/lang/ClassLoader;")
	_class(getComponentType, "getComponentType", "()Ljava/lang/Class;")
	_class(getConstantPool, "getConstantPool", "()Lsun/reflect/ConstantPool;")
	_class(getDeclaringClass0, "getDeclaringClass0", "()Ljava/lang/Class;")
	_class(getEnclosingMethod0, "getEnclosingMethod0", "()[Ljava/lang/Object;")
	_class(getInterfaces0, "getInterfaces0", "()[Ljava/lang/Class;")
	_class(getModifiers, "getModifiers", "()I")
	_class(getName0, "getName0", "()Ljava/lang/String;")
	_class(getSuperclass, "getSuperclass", "()Ljava/lang/Class;")
	_class(isArray, "isArray", "()Z")
	_class(isAssignableFrom, "isAssignableFrom", "(Ljava/lang/Class;)Z")
	_class(isInstance, "isInstance", "(Ljava/lang/Object;)Z")
	_class(isInterface, "isInterface", "()Z")
	_class(isPrimitive, "isPrimitive", "()Z")
	_class(getGenericSignature0, "getGenericSignature0", "()Ljava/lang/String;")
}

func _class(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("java/lang/Class", name, desc, method)
}

// native ClassLoader getClassLoader0();
// ()Ljava/lang/ClassLoader;
func getClassLoader0(frame *rtda.Frame) {
	class := _popClass(frame)
	from := class.LoadedFrom()

	stack := frame.OperandStack()
	if cp.IsBootClassPath(from) {
		stack.PushRef(nil)
		return
	}

	clClass := heap.BootLoader().LoadClass("java/lang/ClassLoader")
	getSysCl := clClass.GetStaticMethod("getSystemClassLoader", "()Ljava/lang/ClassLoader;")
	frame.Thread().InvokeMethod(getSysCl)
}

// public native Class<?> getComponentType();
// ()Ljava/lang/Class;
func getComponentType(frame *rtda.Frame) {
	class := _popClass(frame)
	componentClass := class.ComponentClass()
	componentClassObj := componentClass.JClass()

	stack := frame.OperandStack()
	stack.PushRef(componentClassObj)
}

// native ConstantPool getConstantPool();
// ()Lsun/reflect/ConstantPool;
func getConstantPool(frame *rtda.Frame) {
	class := _popClass(frame)
	cpClass := heap.BootLoader().LoadClass("sun/reflect/ConstantPool")
	if cpClass.InitializationNotStarted() {
		frame.RevertNextPC()
		frame.Thread().InitClass(cpClass)
		return
	}

	cp := class.ConstantPool()
	cpObj := cpClass.NewObjWithExtra(cp) // todo init cpObj
	frame.OperandStack().PushRef(cpObj)
}

// private native Class<?> getDeclaringClass0();
// ()Ljava/lang/Class;
func getDeclaringClass0(frame *rtda.Frame) {
	class := _popClass(frame)
	if class.IsArray() || class.IsPrimitive() {
		frame.OperandStack().PushRef(nil)
		return
	}

	lastDollarIndex := strings.LastIndex(class.Name(), "$")
	if lastDollarIndex < 0 {
		frame.OperandStack().PushRef(nil)
		return
	}

	// todo
	declaringClassName := class.Name()[:lastDollarIndex]
	declaringClass := frame.ClassLoader().LoadClass(declaringClassName)
	frame.OperandStack().PushRef(declaringClass.JClass())
}

// private native Object[] getEnclosingMethod0();
// ()[Ljava/lang/Object;
func getEnclosingMethod0(frame *rtda.Frame) {
	class := _popClass(frame)
	if class.IsPrimitive() {
		frame.OperandStack().PushNull()
	} else {
		emInfo := class.EnclosingMethod()
		emInfoObj := _createEnclosintMethodInfo(frame.ClassLoader(), emInfo)
		if emInfoObj == nil || heap.ArrayLength(emInfoObj) == 0 {
			frame.OperandStack().PushNull()
		} else {
			frame.OperandStack().PushRef(emInfoObj)
		}
	}
}

func _createEnclosintMethodInfo(classLoader *heap.ClassLoader, emInfo *heap.EnclosingMethod) *heap.Object {
	if emInfo == nil {
		return nil
	}

	enclosingClass := classLoader.LoadClass(emInfo.ClassName())
	enclosingClassObj := enclosingClass.JClass()
	var methodNameObj, methodDescriptorObj *heap.Object
	if emInfo.MethodName() != "" {
		methodNameObj = rtda.JString(emInfo.MethodName())
		methodDescriptorObj = rtda.JString(emInfo.MethodDescriptor())
	} else {
		methodNameObj, methodDescriptorObj = nil, nil
	}

	objs := []*heap.Object{enclosingClassObj, methodNameObj, methodDescriptorObj}
	return heap.NewRefArray2(classLoader.JLObjectClass(), objs) // Object[]
}

// private native Class<?>[] getInterfaces0();
// ()[Ljava/lang/Class;
func getInterfaces0(frame *rtda.Frame) {
	class := _popClass(frame)
	interfaces := class.Interfaces()
	interfaceObjs := make([]*heap.Object, len(interfaces))
	for i, iface := range interfaces {
		interfaceObjs[i] = iface.JClass()
	}

	jlClassClass := heap.BootLoader().JLClassClass()
	interfaceArr := heap.NewRefArray2(jlClassClass, interfaceObjs)

	stack := frame.OperandStack()
	stack.PushRef(interfaceArr)
}

// private native String getName0();
// ()Ljava/lang/String;
func getName0(frame *rtda.Frame) {
	class := _popClass(frame)
	name := class.NameJlsFormat()
	nameObj := rtda.JString(name)

	stack := frame.OperandStack()
	stack.PushRef(nameObj)
}

// public native int getModifiers();
// ()I
func getModifiers(frame *rtda.Frame) {
	class := _popClass(frame)
	modifiers := class.GetAccessFlags()

	stack := frame.OperandStack()
	stack.PushInt(int32(modifiers))
}

// public native Class<? super T> getSuperclass();
// ()Ljava/lang/Class;
func getSuperclass(frame *rtda.Frame) {
	class := _popClass(frame)
	superClass := class.SuperClass()

	stack := frame.OperandStack()
	if superClass != nil {
		stack.PushRef(superClass.JClass())
	} else {
		stack.PushNull()
	}
}

// public native boolean isAssignableFrom(Class<?> cls);
// (Ljava/lang/Class;)Z
func isAssignableFrom(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()
	cls := vars.GetRef(1)

	thisClass := this.Extra().(*heap.Class)
	clsClass := cls.Extra().(*heap.Class)
	ok := thisClass.IsAssignableFrom(clsClass)

	stack := frame.OperandStack()
	stack.PushBoolean(ok)
}

// public native boolean isInstance(Object obj);
// (Ljava/lang/Object;)Z
func isInstance(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()
	obj := vars.GetRef(1)

	class := this.Extra().(*heap.Class)
	ok := obj.IsInstanceOf(class)

	stack := frame.OperandStack()
	stack.PushBoolean(ok)
}

// public native boolean isArray();
// ()Z
func isArray(frame *rtda.Frame) {
	class := _popClass(frame)
	stack := frame.OperandStack()
	stack.PushBoolean(class.IsArray())
}

// public native boolean isInterface();
// ()Z
func isInterface(frame *rtda.Frame) {
	class := _popClass(frame)
	stack := frame.OperandStack()
	stack.PushBoolean(class.IsInterface())
}

// public native boolean isPrimitive();
// ()Z
func isPrimitive(frame *rtda.Frame) {
	class := _popClass(frame)
	stack := frame.OperandStack()
	stack.PushBoolean(class.IsPrimitive())
}

// private native String getGenericSignature0();
// ()Ljava/lang/String;
func getGenericSignature0(frame *rtda.Frame) {
	class := _popClass(frame)
	if class == nil {
		panic("illegal class")
	}

	// Return null for arrays and primatives
	if !class.IsPrimitive() {
		signature := class.Signature()
		if signature == "" {
			frame.OperandStack().PushNull()
		} else {
			frame.OperandStack().PushRef(rtda.JString(signature))
		}
		return
	}

	frame.OperandStack().PushNull()
}

func _popClass(frame *rtda.Frame) *heap.Class {
	vars := frame.LocalVars()
	this := vars.GetThis()
	return this.Extra().(*heap.Class)
}
