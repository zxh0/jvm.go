package lang

import (
	. "jvmgo/any"
	"jvmgo/jvm/rtda"
	rtc "jvmgo/jvm/rtda/class"
)

/*
Method(Class<?> declaringClass,
       String name,
       Class<?>[] parameterTypes,
       Class<?> returnType,
       Class<?>[] checkedExceptions,
       int modifiers,
       int slot,
       String signature,
       byte[] annotations,
       byte[] parameterAnnotations,
       byte[] annotationDefault)
*/
const _methodConstructorDescriptor = "" +
	"(Ljava/lang/Class;" +
	"Ljava/lang/String;" +
	"[Ljava/lang/Class;" +
	"Ljava/lang/Class;" +
	"[Ljava/lang/Class;" +
	"II" +
	"Ljava/lang/String;" +
	"[B[B[B)V"

func init() {
	_class(getDeclaredMethods0, "getDeclaredMethods0", "(Z)[Ljava/lang/reflect/Method;")
}

// private native Method[] getDeclaredMethods0(boolean publicOnly);
// (Z)[Ljava/lang/reflect/Method;
func getDeclaredMethods0(frame *rtda.Frame) {
	vars := frame.LocalVars()
	classObj := vars.GetThis()
	publicOnly := vars.GetBoolean(1)

	class := classObj.Extra().(*rtc.Class)
	methods := class.GetMethods(publicOnly)
	methodCount := uint(len(methods))
	methodClass := class.ClassLoader().LoadClass("java/lang/reflect/Method")
	methodConstructor := methodClass.GetConstructor(_methodConstructorDescriptor)
	methodArrObj := methodClass.NewArray(methodCount)

	stack := frame.OperandStack()
	stack.PushRef(methodArrObj)

	// create method objs
	if methodCount > 0 {
		thread := frame.Thread()
		methodObjs := methodArrObj.Fields().([]*rtc.Obj)
		for i, method := range methods {
			methodObj := methodClass.NewObjWithExtra(method)
			methodObjs[i] = methodObj

			// init method obj
			args := _methodConstructorArgs(classObj, methodObj, method)
			thread.InvokeMethodWithShim(methodConstructor, args)
		}
	}
}

func _methodConstructorArgs(classObj, methodObj *rtc.Obj, method *rtc.Method) []Any {
	nameObj := rtda.NewJString(method.Name(), method)

	return []Any{
		methodObj,                      // this
		classObj,                       // declaringClass
		nameObj,                        // name
		getParameterTypeArr(method),    // parameterTypes
		getReturnType(method),          // returnType
		getExceptionTypeArr(method),    // checkedExceptions
		int32(method.GetAccessFlags()), // modifiers
		int32(0),                       // todo slot
		getSignature(&method.ClassMember),         // signature
		getAnnotationByteArr(&method.ClassMember), // annotations
		getParameterAnnotationDyteArr(method),     // parameterAnnotations
		getAnnotationDefaultData(method),          // annotationDefault
	}
}
