package misc

import (
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	"github.com/zxh0/jvm.go/jvmgo/rtda/heap"
)

func init() {
	_unsafe(arrayBaseOffset, "arrayBaseOffset", "(Ljava/lang/Class;)I")
	_unsafe(arrayIndexScale, "arrayIndexScale", "(Ljava/lang/Class;)I")
	_unsafe(objectFieldOffset, "objectFieldOffset", "(Ljava/lang/reflect/Field;)J")
	_unsafe(getBoolean, "getBoolean", "(Ljava/lang/Object;J)Z")
	_unsafe(putBoolean, "putBoolean", "(Ljava/lang/Object;JZ)V")
	_unsafe(getByte, "getByte", "(Ljava/lang/Object;J)B")
	_unsafe(putByte, "putByte", "(Ljava/lang/Object;JB)V")
	_unsafe(getChar, "getChar", "(Ljava/lang/Object;J)C")
	_unsafe(putChar, "putChar", "(Ljava/lang/Object;JC)V")
	_unsafe(getShort, "getShort", "(Ljava/lang/Object;J)S")
	_unsafe(putShort, "putShort", "(Ljava/lang/Object;JS)V")
	_unsafe(getInt, "getInt", "(Ljava/lang/Object;J)I")
	_unsafe(putInt, "putInt", "(Ljava/lang/Object;JI)V")
	_unsafe(getLong, "getLong", "(Ljava/lang/Object;J)J")
	_unsafe(putLong, "putLong", "(Ljava/lang/Object;JJ)V")
	_unsafe(getFloat, "getFloat", "(Ljava/lang/Object;J)F")
	_unsafe(putFloat, "putFloat", "(Ljava/lang/Object;JF)V")
	_unsafe(getDouble, "getDouble", "(Ljava/lang/Object;J)D")
	_unsafe(putDouble, "putDouble", "(Ljava/lang/Object;JD)V")
	_unsafe(getObject, "getObject", "(Ljava/lang/Object;J)Ljava/lang/Object;")
	_unsafe(putObject, "putObject", "(Ljava/lang/Object;JLjava/lang/Object;)V")
	// todo
	_unsafe(getObject, "getObjectVolatile", "(Ljava/lang/Object;J)Ljava/lang/Object;")
	_unsafe(putObject, "putObjectVolatile", "(Ljava/lang/Object;JLjava/lang/Object;)V")
	_unsafe(getObject, "getOrderedObject", "(Ljava/lang/Object;J)Ljava/lang/Object;")
	_unsafe(putObject, "putOrderedObject", "(Ljava/lang/Object;JLjava/lang/Object;)V")
	_unsafe(getInt, "getIntVolatile", "(Ljava/lang/Object;J)I")
	_unsafe(getLong, "getLongVolatile", "(Ljava/lang/Object;J)J")
	_unsafe(getBoolean, "getBooleanVolatile", "(Ljava/lang/Object;J)Z")
	_unsafe(getByte, "getByteVolatile", "(Ljava/lang/Object;J)B")
	_unsafe(getChar, "getCharVolatile", "(Ljava/lang/Object;J)C")
	_unsafe(getShort, "getShortVolatile", "(Ljava/lang/Object;J)S")
	_unsafe(getFloat, "getFloatVolatile", "(Ljava/lang/Object;J)F")
	_unsafe(getDouble, "getDoubleVolatile", "(Ljava/lang/Object;J)D")
}

// public native int arrayBaseOffset(Class<?> type);
// (Ljava/lang/Class;)I
func arrayBaseOffset(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PushInt(0) // todo
}

// public native int arrayIndexScale(Class<?> type);
// (Ljava/lang/Class;)I
func arrayIndexScale(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PushInt(1) // todo
}

// public native long objectFieldOffset(Field field);
// (Ljava/lang/reflect/Field;)J
func objectFieldOffset(frame *rtda.Frame) {
	vars := frame.LocalVars()
	jField := vars.GetRef(1)

	offset := jField.GetFieldValue("slot", "I").(int32)

	stack := frame.OperandStack()
	stack.PushLong(int64(offset))
}

// public native boolean getBoolean(Object o, long offset);
// (Ljava/lang/Object;J)Z
func getBoolean(frame *rtda.Frame) {
	vars := frame.LocalVars()
	fields := vars.GetRef(1).Fields()
	offset := vars.GetLong(2)

	stack := frame.OperandStack()
	if anys, ok := fields.([]interface{}); ok {
		// object
		stack.PushBoolean(anys[offset].(int32) == 1)
	} else if bytes, ok := fields.([]int8); ok {
		// byte[]
		stack.PushBoolean(bytes[offset] == 1)
	} else {
		panic("getBoolean!")
	}
}

// public native void putBoolean(Object o, long offset, boolean x);
// (Ljava/lang/Object;JZ)V
func putBoolean(frame *rtda.Frame) {
	vars := frame.LocalVars()
	fields := vars.GetRef(1).Fields()
	offset := vars.GetLong(2)
	x := vars.GetInt(4)

	if anys, ok := fields.([]interface{}); ok {
		// object
		anys[offset] = x
	} else if bytes, ok := fields.([]int8); ok {
		// byte[]
		bytes[offset] = int8(x)
	} else {
		panic("putBoolean!")
	}
}

// public native boolean getByte(Object o, long offset);
// (Ljava/lang/Object;J)B
func getByte(frame *rtda.Frame) {
	vars := frame.LocalVars()
	fields := vars.GetRef(1).Fields()
	offset := vars.GetLong(2)

	stack := frame.OperandStack()
	if anys, ok := fields.([]interface{}); ok {
		// object
		stack.PushInt(anys[offset].(int32))
	} else if bytes, ok := fields.([]int8); ok {
		// byte[]
		stack.PushInt(int32(bytes[offset]))
	} else {
		panic("getByte!")
	}
}

// public native void putByte(Object o, long offset, byte x);
// (Ljava/lang/Object;JB)V
func putByte(frame *rtda.Frame) {
	vars := frame.LocalVars()
	fields := vars.GetRef(1).Fields()
	offset := vars.GetLong(2)
	x := vars.GetInt(4)

	if anys, ok := fields.([]interface{}); ok {
		// object
		anys[offset] = x
	} else if bytes, ok := fields.([]int8); ok {
		// byte[]
		bytes[offset] = int8(x)
	} else {
		panic("putByte!")
	}
}

// public native boolean getChar(Object o, long offset);
// (Ljava/lang/Object;J)C
func getChar(frame *rtda.Frame) {
	vars := frame.LocalVars()
	fields := vars.GetRef(1).Fields()
	offset := vars.GetLong(2)

	stack := frame.OperandStack()
	if anys, ok := fields.([]interface{}); ok {
		// object
		stack.PushInt(anys[offset].(int32))
	} else if chars, ok := fields.([]uint16); ok {
		// char[]
		stack.PushInt(int32(chars[offset]))
	} else {
		panic("getChar!")
	}
}

// public native void putChar(Object o, long offset, char x);
// (Ljava/lang/Object;JC)V
func putChar(frame *rtda.Frame) {
	vars := frame.LocalVars()
	fields := vars.GetRef(1).Fields()
	offset := vars.GetLong(2)
	x := vars.GetInt(4)

	if anys, ok := fields.([]interface{}); ok {
		// object
		anys[offset] = x
	} else if chars, ok := fields.([]uint16); ok {
		// char[]
		chars[offset] = uint16(x)
	} else {
		panic("putChar!")
	}
}

// public native boolean getShort(Object o, long offset);
// (Ljava/lang/Object;J)S
func getShort(frame *rtda.Frame) {
	vars := frame.LocalVars()
	fields := vars.GetRef(1).Fields()
	offset := vars.GetLong(2)

	stack := frame.OperandStack()
	if anys, ok := fields.([]interface{}); ok {
		// object
		stack.PushInt(anys[offset].(int32))
	} else if shorts, ok := fields.([]int16); ok {
		// short[]
		stack.PushInt(int32(shorts[offset]))
	} else {
		panic("getShort!")
	}
}

// public native void putShort(Object o, long offset, short x);
// (Ljava/lang/Object;JS)V
func putShort(frame *rtda.Frame) {
	vars := frame.LocalVars()
	fields := vars.GetRef(1).Fields()
	offset := vars.GetLong(2)
	x := vars.GetInt(4)

	if anys, ok := fields.([]interface{}); ok {
		// object
		anys[offset] = x
	} else if shorts, ok := fields.([]int16); ok {
		// short[]
		shorts[offset] = int16(x)
	} else {
		panic("putShort!")
	}
}

// public native boolean getInt(Object o, long offset);
// (Ljava/lang/Object;J)I
func getInt(frame *rtda.Frame) {
	vars := frame.LocalVars()
	fields := vars.GetRef(1).Fields()
	offset := vars.GetLong(2)

	stack := frame.OperandStack()
	if anys, ok := fields.([]interface{}); ok {
		// object
		stack.PushInt(anys[offset].(int32))
	} else if shorts, ok := fields.([]int32); ok {
		// int[]
		stack.PushInt(int32(shorts[offset]))
	} else {
		panic("getInt!")
	}
}

// public native void putInt(Object o, long offset, int x);
// (Ljava/lang/Object;JI)V
func putInt(frame *rtda.Frame) {
	vars := frame.LocalVars()
	fields := vars.GetRef(1).Fields()
	offset := vars.GetLong(2)
	x := vars.GetInt(4)

	if anys, ok := fields.([]interface{}); ok {
		// object
		anys[offset] = x
	} else if shorts, ok := fields.([]int32); ok {
		// int[]
		shorts[offset] = x
	} else {
		panic("putInt!")
	}
}

// public native boolean getLong(Object o, long offset);
// (Ljava/lang/Object;J)J
func getLong(frame *rtda.Frame) {
	vars := frame.LocalVars()
	fields := vars.GetRef(1).Fields()
	offset := vars.GetLong(2)

	stack := frame.OperandStack()
	if anys, ok := fields.([]interface{}); ok {
		// object
		stack.PushLong(anys[offset].(int64))
	} else if longs, ok := fields.([]int64); ok {
		// long[]
		stack.PushLong(longs[offset])
	} else {
		panic("getLong!")
	}
}

// public native void putLong(Object o, long offset, long x);
// (Ljava/lang/Object;JJ)V
func putLong(frame *rtda.Frame) {
	vars := frame.LocalVars()
	fields := vars.GetRef(1).Fields()
	offset := vars.GetLong(2)
	x := vars.GetLong(4)

	if anys, ok := fields.([]interface{}); ok {
		// object
		anys[offset] = x
	} else if longs, ok := fields.([]int64); ok {
		// long[]
		longs[offset] = x
	} else {
		panic("putLong!")
	}
}

// public native boolean getFloat(Object o, long offset);
// (Ljava/lang/Object;J)F
func getFloat(frame *rtda.Frame) {
	vars := frame.LocalVars()
	fields := vars.GetRef(1).Fields()
	offset := vars.GetLong(2)

	stack := frame.OperandStack()
	if anys, ok := fields.([]interface{}); ok {
		// object
		stack.PushFloat(anys[offset].(float32))
	} else if floats, ok := fields.([]float32); ok {
		// float[]
		stack.PushFloat(floats[offset])
	} else {
		panic("getFloat!")
	}
}

// public native void putFloat(Object o, long offset, float x);
// (Ljava/lang/Object;JF)V
func putFloat(frame *rtda.Frame) {
	vars := frame.LocalVars()
	fields := vars.GetRef(1).Fields()
	offset := vars.GetLong(2)
	x := vars.GetFloat(4)

	if anys, ok := fields.([]interface{}); ok {
		// object
		anys[offset] = x
	} else if floats, ok := fields.([]float32); ok {
		// float[]
		floats[offset] = x
	} else {
		panic("putFloat!")
	}
}

// public native boolean getDouble(Object o, long offset);
// (Ljava/lang/Object;J)D
func getDouble(frame *rtda.Frame) {
	vars := frame.LocalVars()
	fields := vars.GetRef(1).Fields()
	offset := vars.GetLong(2)

	stack := frame.OperandStack()
	if anys, ok := fields.([]interface{}); ok {
		// object
		stack.PushDouble(anys[offset].(float64))
	} else if doubles, ok := fields.([]float64); ok {
		// double[]
		stack.PushDouble(doubles[offset])
	} else {
		panic("getDouble!")
	}
}

// public native void putDouble(Object o, long offset, double x);
// (Ljava/lang/Object;JD)V
func putDouble(frame *rtda.Frame) {
	vars := frame.LocalVars()
	fields := vars.GetRef(1).Fields()
	offset := vars.GetLong(2)
	x := vars.GetDouble(4)

	if anys, ok := fields.([]interface{}); ok {
		// object
		anys[offset] = x
	} else if doubles, ok := fields.([]float64); ok {
		// double[]
		doubles[offset] = x
	} else {
		panic("putDouble!")
	}
}

// public native void putObject(Object o, long offset, Object x);
// (Ljava/lang/Object;JLjava/lang/Object;)V
func putObject(frame *rtda.Frame) {
	vars := frame.LocalVars()
	fields := vars.GetRef(1).Fields()
	offset := vars.GetLong(2)
	x := vars.GetRef(4)

	if anys, ok := fields.([]interface{}); ok {
		// object
		anys[offset] = x
	} else if objs, ok := fields.([]*heap.Object); ok {
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

	if anys, ok := fields.([]interface{}); ok {
		// object
		x := _getObj(anys, offset)
		frame.OperandStack().PushRef(x)
	} else if objs, ok := fields.([]*heap.Object); ok {
		// ref[]
		x := objs[offset]
		frame.OperandStack().PushRef(x)
	} else {
		panic("getObject!")
	}
}
func _getObj(fields []interface{}, offset int64) *heap.Object {
	f := fields[offset]
	if f != nil {
		return f.(*heap.Object)
	} else {
		return nil
	}
}
