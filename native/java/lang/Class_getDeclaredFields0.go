package lang

import (
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
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
	classObj := frame.GetThis()
	publicOnly := frame.GetBooleanVar(1)

	class := classObj.GetGoClass()
	fields := class.GetFields(publicOnly)
	fieldCount := uint(len(fields))

	fieldClass := frame.GetBootLoader().LoadClass("java/lang/reflect/Field")
	fieldArr := fieldClass.NewArray(fieldCount)

	frame.PushRef(fieldArr)

	if fieldCount > 0 {
		rt := frame.GetRuntime()
		thread := frame.Thread
		fieldObjs := fieldArr.GetRefs()
		fieldConstructor := fieldClass.GetConstructor(_fieldConstructorDescriptor)
		for i, goField := range fields {
			fieldObj := fieldClass.NewObjWithExtra(goField)
			fieldObjs[i] = fieldObj

			// init fieldObj
			thread.InvokeMethodWithShim(fieldConstructor, []heap.Slot{
				heap.NewRefSlot(fieldObj),                                         // this
				heap.NewRefSlot(classObj),                                         // declaringClass
				heap.NewRefSlot(frame.GetRuntime().JSFromGoStr(goField.Name)),     // name
				heap.NewRefSlot(goField.Type().JClass),                            // type
				heap.NewIntSlot(int32(goField.AccessFlags)),                       // modifiers
				heap.NewIntSlot(int32(goField.SlotId)),                            // slot
				heap.NewRefSlot(getSignatureStr(rt, goField.Signature)),           // signature
				heap.NewRefSlot(getAnnotationByteArr(rt, goField.AnnotationData)), // annotations
			})
		}
	}
}
