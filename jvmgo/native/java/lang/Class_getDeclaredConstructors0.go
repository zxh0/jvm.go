package lang

import (
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	"github.com/zxh0/jvm.go/jvmgo/rtda/heap"
)

/*
Constructor(Class<T> declaringClass,
            Class<?>[] parameterTypes,
            Class<?>[] checkedExceptions,
            int modifiers,
            int slot,
            String signature,
            byte[] annotations,
            byte[] parameterAnnotations)
}
*/
const _constructorConstructorDescriptor = "" +
	"(Ljava/lang/Class;" +
	"[Ljava/lang/Class;" +
	"[Ljava/lang/Class;" +
	"II" +
	"Ljava/lang/String;" +
	"[B[B)V"

func init() {
	_class(getDeclaredConstructors0, "getDeclaredConstructors0", "(Z)[Ljava/lang/reflect/Constructor;")
}

// private native Constructor<T>[] getDeclaredConstructors0(boolean publicOnly);
// (Z)[Ljava/lang/reflect/Constructor;
func getDeclaredConstructors0(frame *rtda.Frame) {
	vars := frame.LocalVars()
	classObj := vars.GetThis()
	publicOnly := vars.GetBoolean(1)

	class := classObj.Extra().(*heap.Class)
	constructors := class.GetConstructors(publicOnly)
	constructorCount := uint(len(constructors))

	constructorClass := heap.BootLoader().LoadClass("java/lang/reflect/Constructor")
	constructorArr := constructorClass.NewArray(constructorCount)

	stack := frame.OperandStack()
	stack.PushRef(constructorArr)

	if constructorCount > 0 {
		thread := frame.Thread()
		constructorObjs := constructorArr.Refs()
		constructorInitMethod := constructorClass.GetConstructor(_constructorConstructorDescriptor)
		for i, constructor := range constructors {
			constructorObj := constructorClass.NewObjWithExtra(constructor)
			constructorObjs[i] = constructorObj

			// init constructorObj
			thread.InvokeMethodWithShim(constructorInitMethod, []interface{}{
				constructorObj,                                     // this
				classObj,                                           // declaringClass
				getParameterTypeArr(constructor),                   // parameterTypes
				getExceptionTypeArr(constructor),                   // checkedExceptions
				int32(constructor.GetAccessFlags()),                // modifiers
				int32(0),                                           // todo slot
				getSignatureStr(constructor.Signature()),           // signature
				getAnnotationByteArr(constructor.AnnotationData()), // annotations
				getAnnotationByteArr(constructor.ParameterAnnotationData()), // parameterAnnotations
			})
		}
	}
}
