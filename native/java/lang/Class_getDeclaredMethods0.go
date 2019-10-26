package lang

import (
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
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
	classObj := frame.GetThis()
	publicOnly := frame.GetBooleanVar(1)

	class := classObj.GetGoClass()
	methods := class.GetMethods(publicOnly)
	methodCount := uint(len(methods))

	methodClass := frame.GetBootLoader().LoadClass("java/lang/reflect/Method")
	methodArr := methodClass.NewArray(methodCount)

	frame.PushRef(methodArr)

	// create method objs
	if methodCount > 0 {
		rt := frame.GetRuntime()
		thread := frame.Thread
		methodObjs := methodArr.GetRefs()
		methodConstructor := methodClass.GetConstructor(_methodConstructorDescriptor)
		for i, method := range methods {
			methodObj := methodClass.NewObjWithExtra(method)
			methodObjs[i] = methodObj

			// init methodObj
			thread.InvokeMethodWithShim(methodConstructor, []heap.Slot{
				heap.NewRefSlot(methodObj),                                                // this
				heap.NewRefSlot(classObj),                                                 // declaringClass
				heap.NewRefSlot(rt.JSFromGoStr(method.Name)),                              // name
				heap.NewRefSlot(getParameterTypeArr(rt, method)),                          // parameterTypes
				heap.NewRefSlot(getReturnType(method)),                                    // returnType
				heap.NewRefSlot(getExceptionTypeArr(rt, method)),                          // checkedExceptions
				heap.NewIntSlot(int32(method.AccessFlags)),                                // modifiers
				heap.NewIntSlot(int32(method.Slot)),                                       // slot
				heap.NewRefSlot(getSignatureStr(rt, method.Signature)),                    // signature
				heap.NewRefSlot(getAnnotationByteArr(rt, method.AnnotationData)),          // annotations
				heap.NewRefSlot(getAnnotationByteArr(rt, method.ParameterAnnotationData)), // parameterAnnotations
				heap.NewRefSlot(getAnnotationByteArr(rt, method.AnnotationDefaultData)),   // annotationDefault
			})
		}
	}
}
