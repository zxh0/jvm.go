package lang

import (
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
const _methodConstructorDescriptor = "(Ljava/lang/Class;" +
	"Ljava/lang/String;" +
	"[Ljava/lang/Class;" +
	"Ljava/lang/Class;" +
	"[Ljava/lang/Class;" +
	"II" +
	"Ljava/lang/String;" +
	"[B[B[B)V"

// private native Method[] getDeclaredMethods0(boolean publicOnly);
// (Z)[Ljava/lang/reflect/Method;
func getDeclaredMethods0(frame *rtda.Frame) {
	vars := frame.LocalVars()
	classObj := vars.GetRef(0) // this
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
			// call <init>
			newFrame := thread.NewFrame(methodConstructor)
			vars := newFrame.LocalVars()
			vars.SetRef(0, methodObj)                             // this
			vars.SetRef(1, classObj)                              // declaringClass
			vars.SetRef(2, rtda.NewJString(method.Name(), frame)) // name
			vars.SetRef(3, getParameterTypeArr(method))           // parameterTypes
			vars.SetRef(4, getReturnType(method))                 // returnType
			vars.SetRef(5, nil)                                   // todo checkedExceptions
			vars.SetInt(6, int32(method.GetAccessFlags()))        // modifiers
			vars.SetInt(7, int32(0))                              // todo slot
			vars.SetRef(8, nil)                                   // todo signature
			vars.SetRef(9, nil)                                   // todo annotations
			vars.SetRef(10, nil)                                  // todo parameterAnnotations
			vars.SetRef(11, nil)                                  // todo annotationDefault
			thread.PushFrame(newFrame)
		}
	}
}
