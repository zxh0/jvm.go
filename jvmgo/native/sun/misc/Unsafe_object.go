package misc

import (
	. "jvmgo/any"
	"jvmgo/jvm/rtda"
	rtc "jvmgo/jvm/rtda/class"
	"jvmgo/util"
	"sync/atomic"
	"unsafe"
)

func init() {
	_unsafe(compareAndSwapInt, "compareAndSwapInt", "(Ljava/lang/Object;JII)Z")
	_unsafe(compareAndSwapLong, "compareAndSwapLong", "(Ljava/lang/Object;JJJ)Z")
	_unsafe(compareAndSwapObject, "compareAndSwapObject", "(Ljava/lang/Object;JLjava/lang/Object;Ljava/lang/Object;)Z")
	_unsafe(putObject, "putObject", "(Ljava/lang/Object;JLjava/lang/Object;)V")
	_unsafe(getObject, "getObject", "(Ljava/lang/Object;J)Ljava/lang/Object;")
	_unsafe(putObjectVolatile, "putObjectVolatile", "(Ljava/lang/Object;JLjava/lang/Object;)V")
	_unsafe(getObjectVolatile, "getObjectVolatile", "(Ljava/lang/Object;J)Ljava/lang/Object;")
	_unsafe(putOrderedObject, "putOrderedObject", "(Ljava/lang/Object;JLjava/lang/Object;)V")
	_unsafe(getOrderedObject, "getOrderedObject", "(Ljava/lang/Object;J)Ljava/lang/Object;")
	_unsafe(getIntVolatile, "getIntVolatile", "(Ljava/lang/Object;J)I")
	_unsafe(getLongVolatile, "getLongVolatile", "(Ljava/lang/Object;J)J")
}

// public final native boolean compareAndSwapInt(Object o, long offset, int expected, int x);
// (Ljava/lang/Object;JII)Z
func compareAndSwapInt(frame *rtda.Frame) {
	vars := frame.LocalVars()
	fields := vars.GetRef(1).Fields()
	offset := vars.GetLong(2)
	expected := vars.GetInt(4)
	newVal := vars.GetInt(5)

	if anys, ok := fields.([]Any); ok {
		// object
		swapped := util.CasInt32(anys[offset], expected, newVal)
		frame.OperandStack().PushBoolean(swapped)
	} else if ints, ok := fields.([]int32); ok {
		// int[]
		swapped := atomic.CompareAndSwapInt32(&ints[offset], expected, newVal)
		frame.OperandStack().PushBoolean(swapped)
	} else {
		// todo
		panic("todo: compareAndSwapInt!")
	}
}

// public final native boolean compareAndSwapLong(Object o, long offset, long expected, long x);
// (Ljava/lang/Object;JJJ)Z
func compareAndSwapLong(frame *rtda.Frame) {
	vars := frame.LocalVars()
	fields := vars.GetRef(1).Fields()
	offset := vars.GetLong(2)
	expected := vars.GetLong(4)
	newVal := vars.GetLong(6)

	if anys, ok := fields.([]Any); ok {
		// object
		swapped := util.CasInt64(anys[offset], expected, newVal)
		frame.OperandStack().PushBoolean(swapped)
	} else if ints, ok := fields.([]int64); ok {
		// long[]
		swapped := atomic.CompareAndSwapInt64(&ints[offset], expected, newVal)
		frame.OperandStack().PushBoolean(swapped)
	} else {
		// todo
		panic("todo: compareAndSwapLong!")
	}
}

// public final native boolean compareAndSwapObject(Object o, long offset, Object expected, Object x)
// (Ljava/lang/Object;JLjava/lang/Object;Ljava/lang/Object;)Z
func compareAndSwapObject(frame *rtda.Frame) {
	vars := frame.LocalVars()
	obj := vars.GetRef(1)
	fields := obj.Fields()
	offset := vars.GetLong(2)
	expected := vars.GetRef(4)
	newVal := vars.GetRef(5)

	// todo
	if anys, ok := fields.([]Any); ok {
		// object
		swapped := _casObj(obj, anys, offset, expected, newVal)
		frame.OperandStack().PushBoolean(swapped)
	} else if objs, ok := fields.([]*rtc.Obj); ok {
		// ref[]
		swapped := _casArr(objs, offset, expected, newVal)
		frame.OperandStack().PushBoolean(swapped)
	} else {
		// todo
		panic("todo: compareAndSwapObject!")
	}
}
func _casObj(obj *rtc.Obj, fields []Any, offset int64, expected, newVal *rtc.Obj) bool {
	// todo
	obj.LockState()
	defer obj.UnlockState()

	current := _getObj(fields, offset)
	if current == expected {
		fields[offset] = newVal
		return true
	} else {
		return false
	}
}
func _casArr(objs []*rtc.Obj, offset int64, expected, newVal *rtc.Obj) bool {
	// cast to []unsafe.Pointer
	ps := *((*[]unsafe.Pointer)(unsafe.Pointer(&objs)))

	addr := &ps[offset]
	old := unsafe.Pointer(expected)
	_new := unsafe.Pointer(newVal)
	return atomic.CompareAndSwapPointer(addr, old, _new)
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
