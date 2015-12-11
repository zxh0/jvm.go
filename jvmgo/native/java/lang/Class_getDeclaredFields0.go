package lang

import (
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	"github.com/zxh0/jvm.go/jvmgo/rtda/heap"
)

/*
Field(Class<?> declaringClass,
      String name,
      Class<?> type,
      int modifiers,
      int slot,
      String signature,
      byte[] annotations)
*/
const _fieldConstructorDescriptor = "" +
	"(Ljava/lang/Class;" +
	"Ljava/lang/String;" +
	"Ljava/lang/Class;" +
	"II" +
	"Ljava/lang/String;" +
	"[B)V"

func init() {
	_class(getDeclaredFields0, "getDeclaredFields0", "(Z)[Ljava/lang/reflect/Field;")
}

// private native Field[] getDeclaredFields0(boolean publicOnly);
// (Z)[Ljava/lang/reflect/Field;
func getDeclaredFields0(frame *rtda.Frame) {
	vars := frame.LocalVars()
	classObj := vars.GetThis()
	publicOnly := vars.GetBoolean(1)

	class := classObj.Extra().(*heap.Class)
	fields := class.GetFields(publicOnly)
	fieldCount := uint(len(fields))

	fieldClass := heap.BootLoader().LoadClass("java/lang/reflect/Field")
	fieldArr := heap.NewRefArray(fieldClass, fieldCount)

	stack := frame.OperandStack()
	stack.PushRef(fieldArr)

	if fieldCount > 0 {
		thread := frame.Thread()
		fieldObjs := fieldArr.Refs()
		fieldConstructor := fieldClass.GetConstructor(_fieldConstructorDescriptor)
		for i, goField := range fields {
			fieldObj := fieldClass.NewObjWithExtra(goField)
			fieldObjs[i] = fieldObj

			// init fieldObj
			thread.InvokeMethodWithShim(fieldConstructor, []interface{}{
				fieldObj,                                       // this
				classObj,                                       // declaringClass
				rtda.JString(goField.Name()),                   // name
				goField.Type().JClass(),                        // type
				int32(goField.GetAccessFlags()),                // modifiers
				int32(goField.SlotId()),                        // slot
				getSignatureStr(goField.Signature()),           // signature
				getAnnotationByteArr(goField.AnnotationData()), // annotations
			})
		}
	}
}
