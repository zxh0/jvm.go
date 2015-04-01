package lang

import (
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
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

	class := classObj.Extra().(*rtc.Class)
	constructors := class.GetConstructors(publicOnly)
	constructorCount := uint(len(constructors))

	constructorClass := rtc.BootLoader().LoadClass("java/lang/reflect/Constructor")
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

			// call <init>
			newFrame := thread.NewFrame(constructorInitMethod)
			vars := newFrame.LocalVars()
			vars.SetRef(0, constructorObj)                                              // this
			vars.SetRef(1, classObj)                                                    // declaringClass
			vars.SetRef(2, getParameterTypeArr(constructor))                            // parameterTypes
			vars.SetRef(3, getExceptionTypeArr(constructor))                            // checkedExceptions
			vars.SetInt(4, int32(constructor.GetAccessFlags()))                         // modifiers
			vars.SetInt(5, int32(0))                                                    // todo slot
			vars.SetRef(6, getSignatureStr(constructor.Signature()))                    // signature
			vars.SetRef(7, getAnnotationByteArr(constructor.AnnotationData()))          // annotations
			vars.SetRef(8, getAnnotationByteArr(constructor.ParameterAnnotationData())) // parameterAnnotations
			thread.PushFrame(newFrame)
		}
	}
}
