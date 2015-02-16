package misc

import (
	//"unsafe"
	//. "jvmgo/any"
	"jvmgo/jvm/rtda"
	rtc "jvmgo/jvm/rtda/class"
)

func init() {
	_unsafe(ensureClassInitialized, "ensureClassInitialized", "(Ljava/lang/Class;)V")
	_unsafe(staticFieldOffset, "staticFieldOffset", "(Ljava/lang/reflect/Field;)J")
	_unsafe(staticFieldBase, "staticFieldBase", "(Ljava/lang/reflect/Field;)Ljava/lang/Object;")
}

// public native void ensureClassInitialized(Class<?> c);
// (Ljava/lang/Class;)V
func ensureClassInitialized(frame *rtda.Frame) {
	vars := frame.LocalVars()
	// this := vars.GetRef(0)
	classObj := vars.GetRef(1)

	goClass := classObj.Extra().(*rtc.Class)
	if goClass.InitializationNotStarted() {
		// undo ensureClassInitialized()
		frame.RevertNextPC()
		// init
		rtda.InitClass(goClass, frame.Thread())
	}
}

// public native long staticFieldOffset(Field f);
// (Ljava/lang/reflect/Field;)J
func staticFieldOffset(frame *rtda.Frame) {
	vars := frame.LocalVars()
	// vars.GetRef(0) // this
	fieldObj := vars.GetRef(1)

	offset := fieldObj.GetFieldValue("slot", "I").(int32)
	stack := frame.OperandStack()
	stack.PushLong(int64(offset))
}

// public native Object staticFieldBase(Field f);
// (Ljava/lang/reflect/Field;)Ljava/lang/Object;
func staticFieldBase(frame *rtda.Frame) {
	vars := frame.LocalVars()
	// vars.GetRef(0) // this
	fieldObj := vars.GetRef(1)

	goField := getExtra(fieldObj)
	goClass := goField.Class()
	obj := goClass.AsObj()

	stack := frame.OperandStack()
	stack.PushRef(obj)
}

func getExtra(fieldObj *rtc.Obj) *rtc.Field {
	extra := fieldObj.Extra()
	if extra != nil {
		return extra.(*rtc.Field)
	}

	root := fieldObj.GetFieldValue("root", "Ljava/lang/reflect/Field;").(*rtc.Obj)
	return root.Extra().(*rtc.Field)
}
