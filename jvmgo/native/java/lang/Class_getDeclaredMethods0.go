package lang

import (
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	"github.com/zxh0/jvm.go/jvmgo/rtda/heap"
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

	class := classObj.Extra().(*heap.Class)
	methods := class.GetMethods(publicOnly)
	methodCount := uint(len(methods))

	methodClass := heap.BootLoader().LoadClass("java/lang/reflect/Method")
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

			// init methodObj
			thread.InvokeMethodWithShim(methodConstructor, []interface{}{
				methodObj,                                              // this
				classObj,                                               // declaringClass
				rtda.JString(method.Name()),                            // name
				getParameterTypeArr(method),                            // parameterTypes
				getReturnType(method),                                  // returnType
				getExceptionTypeArr(method),                            // checkedExceptions
				int32(method.GetAccessFlags()),                         // modifiers
				int32(method.Slot()),                                   // slot
				getSignatureStr(method.Signature()),                    // signature
				getAnnotationByteArr(method.AnnotationData()),          // annotations
				getAnnotationByteArr(method.ParameterAnnotationData()), // parameterAnnotations
				getAnnotationByteArr(method.AnnotationDefaultData()),   // annotationDefault
			})
		}
	}
}
