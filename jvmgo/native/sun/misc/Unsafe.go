package misc

import (
	//"unsafe"
	. "github.com/zxh0/jvm.go/jvmgo/any"
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
	"time"
)

func init() {
	_unsafe(arrayBaseOffset, "arrayBaseOffset", "(Ljava/lang/Class;)I")
	_unsafe(arrayIndexScale, "arrayIndexScale", "(Ljava/lang/Class;)I")
	_unsafe(objectFieldOffset, "objectFieldOffset", "(Ljava/lang/reflect/Field;)J")
	_unsafe(park, "park", "(ZJ)V")
}

func _unsafe(method Any, name, desc string) {
	rtc.RegisterNativeMethod("sun/misc/Unsafe", name, desc, method)
}

// public native int arrayBaseOffset(Class<?> type);
// (Ljava/lang/Class;)I
func arrayBaseOffset(frame *rtda.Frame) {
	vars := frame.LocalVars()
	vars.GetRef(0) // this
	vars.GetRef(1) // type

	stack := frame.OperandStack()
	stack.PushInt(0) // todo
}

// public native int arrayIndexScale(Class<?> type);
// (Ljava/lang/Class;)I
func arrayIndexScale(frame *rtda.Frame) {
	//stack.PopRef() // type
	//stack.PopRef() // this

	stack := frame.OperandStack()
	stack.PushInt(1) // todo
}

// public native long objectFieldOffset(Field field);
// (Ljava/lang/reflect/Field;)J
func objectFieldOffset(frame *rtda.Frame) {
	vars := frame.LocalVars()
	// this := vars.GetRef(0)
	jField := vars.GetRef(1)

	offset := jField.GetFieldValue("slot", "I").(int32)
	stack := frame.OperandStack()
	stack.PushLong(int64(offset))
}

// public native void park(boolean var1, long var2);
// (ZJ)V [
func park(frame *rtda.Frame) {
	vars := frame.LocalVars()
	absolute := vars.GetBoolean(1)
	var2 := vars.GetLong(2)
	var parkTime time.Duration

	//deadline the absolute time, in milliseconds from the Epoch,
	// *        to wait until
	if absolute {
		parkTime = time.Duration((time.Now().UnixNano() / 1000) - var2)
	} else {
		parkTime = time.Duration(var2)
	}
	frame.Thread().Sleep(parkTime)
}
