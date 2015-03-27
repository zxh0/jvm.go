package misc

import (
	. "github.com/zxh0/jvm.go/jvmgo/any"
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
)

func init() {
	_unsafe(putObject, "putObject", "(Ljava/lang/Object;JLjava/lang/Object;)V")
	_unsafe(getObject, "getObject", "(Ljava/lang/Object;J)Ljava/lang/Object;")
	_unsafe(putObjectVolatile, "putObjectVolatile", "(Ljava/lang/Object;JLjava/lang/Object;)V")
	_unsafe(getObjectVolatile, "getObjectVolatile", "(Ljava/lang/Object;J)Ljava/lang/Object;")
	_unsafe(putOrderedObject, "putOrderedObject", "(Ljava/lang/Object;JLjava/lang/Object;)V")
	_unsafe(getOrderedObject, "getOrderedObject", "(Ljava/lang/Object;J)Ljava/lang/Object;")
	_unsafe(getIntVolatile, "getIntVolatile", "(Ljava/lang/Object;J)I")
	_unsafe(getLongVolatile, "getLongVolatile", "(Ljava/lang/Object;J)J")
}

// public native void putObject(Object o, long offset, Object x);
// (Ljava/lang/Object;JLjava/lang/Object;)V
func putObject(frame *rtda.Frame) {
	vars := frame.LocalVars()
	fields := vars.GetRef(1).Fields()
	offset := vars.GetLong(2)
	x := vars.GetRef(4)

	if anys, ok := fields.([]Any); ok {
		// object
		anys[offset] = x
	} else if objs, ok := fields.([]*rtc.Obj); ok {
		// ref[]
		objs[offset] = x
	} else {
		panic("putObject!")
	}
}

// public native Object getObject(Object o, long offset);
// (Ljava/lang/Object;J)Ljava/lang/Object;
func getObject(frame *rtda.Frame) {
	vars := frame.LocalVars()
	fields := vars.GetRef(1).Fields()
	offset := vars.GetLong(2)

	if anys, ok := fields.([]Any); ok {
		// object
		x := _getObj(anys, offset)
		frame.OperandStack().PushRef(x)
	} else if objs, ok := fields.([]*rtc.Obj); ok {
		// ref[]
		x := objs[offset]
		frame.OperandStack().PushRef(x)
	} else {
		panic("getObject!")
	}
}
func _getObj(fields []Any, offset int64) *rtc.Obj {
	f := fields[offset]
	if f != nil {
		return f.(*rtc.Obj)
	} else {
		return nil
	}
}

// public native void putObjectVolatile(Object o, long offset, Object x);
// (Ljava/lang/Object;JLjava/lang/Object;)V
func putObjectVolatile(frame *rtda.Frame) {
	putObject(frame) // todo
}

// public native Object getObjectVolatile(Object o, long offset);
//(Ljava/lang/Object;J)Ljava/lang/Object;
func getObjectVolatile(frame *rtda.Frame) {
	getObject(frame) // todo
}

// public native void putOrderedObject(Object o, long offset, Object x);
// (Ljava/lang/Object;JLjava/lang/Object;)V
func putOrderedObject(frame *rtda.Frame) {
	putObjectVolatile(frame) // todo
}

// public native Object getOrderedObject(Object o, long offset);
//(Ljava/lang/Object;J)Ljava/lang/Object;
func getOrderedObject(frame *rtda.Frame) {
	getObjectVolatile(frame) // todo
}

// public native int getIntVolatile(Object o, long offset);
// (Ljava/lang/Object;J)I
func getIntVolatile(frame *rtda.Frame) {
	vars := frame.LocalVars()
	obj := vars.GetRef(1)
	offset := vars.GetLong(2)

	// todo
	value := obj.Fields().([]Any)[offset].(int32)
	frame.OperandStack().PushInt(value)
}

// public native long getLongVolatile(Object o, long offset);
// (Ljava/lang/Object;J)J
func getLongVolatile(frame *rtda.Frame) {
	vars := frame.LocalVars()
	obj := vars.GetRef(1)
	offset := vars.GetLong(2)

	// todo
	value := obj.Fields().([]Any)[offset].(int64)
	frame.OperandStack().PushLong(value)
}
