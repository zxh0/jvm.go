package misc

import (
	. "jvmgo/any"
	"jvmgo/jvm/rtda"
	rtc "jvmgo/jvm/rtda/class"
	"jvmgo/util"
)

func init() {
	_unsafe(compareAndSwapInt, "compareAndSwapInt", "(Ljava/lang/Object;JII)Z")
	_unsafe(compareAndSwapLong, "compareAndSwapLong", "(Ljava/lang/Object;JJJ)Z")
	_unsafe(putObject, "putObject", "(Ljava/lang/Object;JLjava/lang/Object;)V")
	_unsafe(getObject, "getObject", "(Ljava/lang/Object;J)Ljava/lang/Object;")
	_unsafe(putObjectVolatile, "putObjectVolatile", "(Ljava/lang/Object;JLjava/lang/Object;)V")
	_unsafe(getObjectVolatile, "getObjectVolatile", "(Ljava/lang/Object;J)Ljava/lang/Object;")
	_unsafe(putOrderedObject, "putOrderedObject", "(Ljava/lang/Object;JLjava/lang/Object;)V")
	_unsafe(getOrderedObject, "getOrderedObject", "(Ljava/lang/Object;J)Ljava/lang/Object;")
}

// public final native boolean compareAndSwapInt(Object o, long offset, int expected, int x);
// (Ljava/lang/Object;JII)Z
func compareAndSwapInt(frame *rtda.Frame) {
	vars := frame.LocalVars()
	// vars.GetRef(0) // this
	o := vars.GetRef(1)
	offset := vars.GetLong(2)
	expected := vars.GetInt(4)
	x := vars.GetInt(5)

	// todo
	fields := o.Fields().([]Any)
	actual := fields[offset].(int32)
	stack := frame.OperandStack()
	if actual == expected {
		fields[offset] = x
		stack.PushBoolean(true)
	} else {
		stack.PushBoolean(false)
	}
}

func compareAndSwapInt2(frame *rtda.Frame) {
	vars := frame.LocalVars()
	fields := vars.GetRef(1).Fields()
	offset := vars.GetLong(2)
	expected := vars.GetInt(4)
	_new := vars.GetInt(5)

	if anys, ok := fields.([]Any); ok {
		swapped := util.CasInt32(anys[offset], expected, _new)
		frame.OperandStack().PushBoolean(swapped)
	} else {
		panic("123")
	}
}

// public final native boolean compareAndSwapLong(Object o, long offset, long expected, long x);
// (Ljava/lang/Object;JJJ)Z
func compareAndSwapLong(frame *rtda.Frame) {
	vars := frame.LocalVars()
	// vars.GetRef(0) // this
	o := vars.GetRef(1)
	offset := vars.GetLong(2)
	expected := vars.GetLong(4)
	x := vars.GetLong(6)

	// todo
	fields := o.Fields().([]Any)
	actual := fields[offset].(int64)
	stack := frame.OperandStack()
	if actual == expected {
		fields[offset] = x
		stack.PushBoolean(true)
	} else {
		stack.PushBoolean(false)
	}
}

// public native void putObject(Object o, long offset, Object x);
// (Ljava/lang/Object;JLjava/lang/Object;)V
func putObject(frame *rtda.Frame) {
	vars := frame.LocalVars()
	obj := vars.GetRef(1)
	offset := vars.GetLong(2)
	x := vars.GetRef(4)

	fields := obj.Fields()
	switch fields.(type) {
	case []Any: // object
		fields.([]Any)[offset] = x
	case []*rtc.Obj: // array of ref
		fields.([]*rtc.Obj)[offset] = x
	}
}

// public native Object getObject(Object o, long offset);
// (Ljava/lang/Object;J)Ljava/lang/Object;
func getObject(frame *rtda.Frame) {
	vars := frame.LocalVars()
	obj := vars.GetRef(1)
	offset := vars.GetLong(2)

	var x *rtc.Obj
	fields := obj.Fields()
	switch fields.(type) {
	case []Any: // object
		x = fields.([]Any)[offset].(*rtc.Obj)
	case []*rtc.Obj: // array of ref
		x = fields.([]*rtc.Obj)[offset]
	}

	stack := frame.OperandStack()
	stack.PushRef(x)
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
