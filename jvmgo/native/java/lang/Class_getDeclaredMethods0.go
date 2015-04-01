package lang

import (
	. "github.com/zxh0/jvm.go/jvmgo/any"
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
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

	methodClass := rtc.BootLoader().LoadClass("java/lang/reflect/Method")
	methodArr := methodClass.NewArray(methodCount)

	stack := frame.OperandStack()
	stack.PushRef(methodArr)

	// create method objs
	if methodCount > 0 {
		thread := frame.Thread()
		methodObjs := methodArr.Refs()
		methodConstructor := methodClass.GetConstructor(_methodConstructorDescriptor)
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
	return []Any{
		methodObj,                      // this
		classObj,                       // declaringClass
		rtda.JString(method.Name()),    // name
		getParameterTypeArr(method),    // parameterTypes
		getReturnType(method),          // returnType
		getExceptionTypeArr(method),    // checkedExceptions
		int32(method.GetAccessFlags()), // modifiers
		int32(0),                       // todo slot
		getSignatureStr(method.Signature()),                    // signature
		getAnnotationByteArr(method.AnnotationData()),          // annotations
		getAnnotationByteArr(method.ParameterAnnotationData()), // parameterAnnotations
		getAnnotationByteArr(method.AnnotationDefaultData()),   // annotationDefault
	}
}
