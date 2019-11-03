package lang

import (
	"strings"

	"github.com/zxh0/jvm.go/native"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

func init() {
	_class(getComponentType, "getComponentType", "()Ljava/lang/Class;")
	_class(getConstantPool, "getConstantPool", "()Lsun/reflect/ConstantPool;")
	_class(getDeclaringClass0, "getDeclaringClass0", "()Ljava/lang/Class;")
	_class(getEnclosingMethod0, "getEnclosingMethod0", "()[Ljava/lang/Object;")
	_class(getInterfaces0, "getInterfaces0", "()[Ljava/lang/Class;")
	_class(getModifiers, "getModifiers", "()I")
	_class(initClassName, "initClassName", "()Ljava/lang/String;")
	_class(getSuperclass, "getSuperclass", "()Ljava/lang/Class;")
	_class(isArray, "isArray", "()Z")
	_class(isAssignableFrom, "isAssignableFrom", "(Ljava/lang/Class;)Z")
	_class(isInstance, "isInstance", "(Ljava/lang/Object;)Z")
	_class(isInterface, "isInterface", "()Z")
	_class(isPrimitive, "isPrimitive", "()Z")
	_class(getGenericSignature0, "getGenericSignature0", "()Ljava/lang/String;")
}

func _class(method native.Method, name, desc string) {
	native.Register("java/lang/Class", name, desc, method)
}

// public native Class<?> getComponentType();
// ()Ljava/lang/Class;
func getComponentType(frame *rtda.Frame) {
	class := frame.GetThis().GetGoClass()
	componentClass := class.GetComponentClass()
	componentClassObj := componentClass.JClass

	frame.PushRef(componentClassObj)
}

// native ConstantPool getConstantPool();
// ()Lsun/reflect/ConstantPool;
func getConstantPool(frame *rtda.Frame) {
	class := frame.GetThis().GetGoClass()
	cpClass := frame.GetBootLoader().LoadClass("sun/reflect/ConstantPool")
	if cpClass.InitializationNotStarted() {
		frame.RevertNextPC()
		frame.Thread.InitClass(cpClass)
		return
	}

	cpObj := cpClass.NewObjWithExtra(class.ConstantPool) // todo init cpObj
	frame.PushRef(cpObj)
}

// private native Class<?> getDeclaringClass0();
// ()Ljava/lang/Class;
func getDeclaringClass0(frame *rtda.Frame) {
	class := frame.GetThis().GetGoClass()
	if class.IsArray() || class.IsPrimitive() {
		frame.PushRef(nil)
		return
	}

	lastDollarIndex := strings.LastIndex(class.Name, "$")
	if lastDollarIndex < 0 {
		frame.PushRef(nil)
		return
	}

	// todo
	declaringClassName := class.Name[:lastDollarIndex]
	declaringClass := frame.GetClassLoader().LoadClass(declaringClassName)
	frame.PushRef(declaringClass.JClass)
}

// private native Object[] getEnclosingMethod0();
// ()[Ljava/lang/Object;
func getEnclosingMethod0(frame *rtda.Frame) {
	class := frame.GetThis().GetGoClass()
	if class.IsPrimitive() {
		frame.PushNull()
	} else {
		emInfo := class.EnclosingMethod
		emInfoObj := _createEnclosintMethodInfo(frame.GetRuntime(), emInfo)
		if emInfoObj == nil || emInfoObj.ArrayLength() == 0 {
			frame.PushNull()
		} else {
			frame.PushRef(emInfoObj)
		}
	}
}

func _createEnclosintMethodInfo(rt *heap.Runtime, emInfo *heap.EnclosingMethod) *heap.Object {
	if emInfo == nil {
		return nil
	}

	bootLoader := rt.BootLoader()
	enclosingClass := bootLoader.LoadClass(emInfo.ClassName)
	enclosingClassObj := enclosingClass.JClass
	var methodNameObj, methodDescriptorObj *heap.Object
	if emInfo.MethodName != "" {
		methodNameObj = rt.JSFromGoStr(emInfo.MethodName)
		methodDescriptorObj = rt.JSFromGoStr(emInfo.MethodDescriptor)
	} else {
		methodNameObj, methodDescriptorObj = nil, nil
	}

	objs := []*heap.Object{enclosingClassObj, methodNameObj, methodDescriptorObj}
	return rt.NewObjectArray(objs)
}

// private native Class<?>[] getInterfaces0();
// ()[Ljava/lang/Class;
func getInterfaces0(frame *rtda.Frame) {
	class := frame.GetThis().GetGoClass()
	interfaces := class.Interfaces
	interfaceObjs := make([]*heap.Object, len(interfaces))
	for i, iface := range interfaces {
		interfaceObjs[i] = iface.JClass
	}

	interfaceArr := frame.GetRuntime().NewClassArray(interfaceObjs)
	frame.PushRef(interfaceArr)
}

// private native String initClassName();
// ()Ljava/lang/String;
func initClassName(frame *rtda.Frame) {
	class := frame.GetThis().GetGoClass()
	name := class.NameJlsFormat()
	jsName := frame.GetRuntime().JSFromGoStr(name)

	frame.PushRef(jsName)
}

// public native int getModifiers();
// ()I
func getModifiers(frame *rtda.Frame) {
	class := frame.GetThis().GetGoClass()
	frame.PushInt(int32(class.AccessFlags))
}

// public native Class<? super T> getSuperclass();
// ()Ljava/lang/Class;
func getSuperclass(frame *rtda.Frame) {
	class := frame.GetThis().GetGoClass()
	superClass := class.SuperClass

	if superClass != nil {
		frame.PushRef(superClass.JClass)
	} else {
		frame.PushNull()
	}
}

// public native boolean isAssignableFrom(Class<?> cls);
// (Ljava/lang/Class;)Z
func isAssignableFrom(frame *rtda.Frame) {
	this := frame.GetThis()
	cls := frame.GetRefVar(1)

	thisClass := this.GetGoClass()
	clsClass := cls.GetGoClass()
	ok := thisClass.IsAssignableFrom(clsClass)

	frame.PushBoolean(ok)
}

// public native boolean isInstance(Object obj);
// (Ljava/lang/Object;)Z
func isInstance(frame *rtda.Frame) {
	this := frame.GetThis()
	obj := frame.GetRefVar(1)

	class := this.GetGoClass()
	ok := obj.IsInstanceOf(class)

	frame.PushBoolean(ok)
}

// public native boolean isArray();
// ()Z
func isArray(frame *rtda.Frame) {
	class := frame.GetThis().GetGoClass()
	frame.PushBoolean(class.IsArray())
}

// public native boolean isInterface();
// ()Z
func isInterface(frame *rtda.Frame) {
	class := frame.GetThis().GetGoClass()
	frame.PushBoolean(class.IsInterface())
}

// public native boolean isPrimitive();
// ()Z
func isPrimitive(frame *rtda.Frame) {
	class := frame.GetThis().GetGoClass()
	frame.PushBoolean(class.IsPrimitive())
}

// private native String getGenericSignature0();
// ()Ljava/lang/String;
func getGenericSignature0(frame *rtda.Frame) {
	class := frame.GetThis().GetGoClass()
	if class == nil {
		panic("illegal class")
	}

	// Return null for arrays and primatives
	if !class.IsPrimitive() {
		signature := class.Signature
		if signature == "" {
			frame.PushNull()
		} else {
			frame.PushRef(frame.GetRuntime().JSFromGoStr(signature))
		}
		return
	}

	frame.PushNull()
}
