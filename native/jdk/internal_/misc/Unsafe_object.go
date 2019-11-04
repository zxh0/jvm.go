package misc

import (
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
	"github.com/zxh0/jvm.go/vmutils"
)

func init() {
	_unsafe(arrayBaseOffset0, "arrayBaseOffset0", "(Ljava/lang/Class;)I")
	_unsafe(arrayIndexScale0, "arrayIndexScale0", "(Ljava/lang/Class;)I")
	_unsafe(objectFieldOffset0, "objectFieldOffset0", "(Ljava/lang/reflect/Field;)J")
	_unsafe(objectFieldOffset1, "objectFieldOffset1", "(Ljava/lang/Class;Ljava/lang/String;)J")
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
	_unsafe(getReference, "getReference", "(Ljava/lang/Object;J)Ljava/lang/Object;")
	_unsafe(putReference, "putReference", "(Ljava/lang/Object;JLjava/lang/Object;)V")
	_unsafe(getReference, "getReferenceVolatile", "(Ljava/lang/Object;J)Ljava/lang/Object;")
	_unsafe(putReference, "putReferenceVolatile", "(Ljava/lang/Object;JLjava/lang/Object;)V")
	_unsafe(getReference, "getOrderedObject", "(Ljava/lang/Object;J)Ljava/lang/Object;")
	_unsafe(putReference, "putOrderedObject", "(Ljava/lang/Object;JLjava/lang/Object;)V")
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
func arrayBaseOffset0(frame *rtda.Frame) {
	frame.PushInt(0) // todo
}

// public native int arrayIndexScale(Class<?> type);
// (Ljava/lang/Class;)I
func arrayIndexScale0(frame *rtda.Frame) {
	frame.PushInt(1) // todo
}

// public native long objectFieldOffset0(Field field);
// (Ljava/lang/reflect/Field;)J
func objectFieldOffset0(frame *rtda.Frame) {
	jField := frame.GetRefVar(1)

	offset := jField.GetFieldValue("slot", "I").IntValue()

	frame.PushLong(int64(offset))
}

// private native long objectFieldOffset1(Class<?> c, String name);
// (Ljava/lang/Class;Ljava/lang/String;)J
func objectFieldOffset1(frame *rtda.Frame) {
	class := frame.GetRefVar(1).GetGoClass()
	name := frame.GetRefVar(2).JSToGoStr()
	field := class.GetInstanceField(name, "*")
	frame.PushLong(int64(field.SlotId))
}

// public native boolean getBoolean(Object o, long offset);
// (Ljava/lang/Object;J)Z
func getBoolean(frame *rtda.Frame) {
	obj, offset := _getArgs(frame)
	switch s := obj.Fields.(type) {
	case []heap.Slot: // object
		frame.PushBoolean(s[offset].IntValue() == 1) // TODO
	case []int8: // bytes
		frame.PushBoolean(s[offset] == 1)
	default:
		panic("getBoolean!")
	}
}

// public native void putBoolean(Object o, long offset, boolean x);
// (Ljava/lang/Object;JZ)V
func putBoolean(frame *rtda.Frame) {
	obj, offset := _getArgs(frame)
	x := frame.GetIntVar(4)
	switch s := obj.Fields.(type) {
	case []heap.Slot: // object
		s[offset] = heap.NewIntSlot(x) // TODO
	case []int8: // byte[]
		s[offset] = int8(x)
	default:
		panic("putBoolean!")
	}
}

// public native byte getByte(Object o, long offset);
// (Ljava/lang/Object;J)B
func getByte(frame *rtda.Frame) {
	obj, offset := _getArgs(frame)
	switch s := obj.Fields.(type) {
	case []heap.Slot: // object
		frame.PushInt(s[offset].IntValue())
	case []int8: // byte[]
		frame.PushInt(int32(s[offset]))
	default:
		panic("getByte!")
	}
}

// public native void putByte(Object o, long offset, byte x);
// (Ljava/lang/Object;JB)V
func putByte(frame *rtda.Frame) {
	obj, offset := _getArgs(frame)
	x := frame.GetIntVar(4)
	switch s := obj.Fields.(type) {
	case []heap.Slot: // object
		s[offset] = heap.NewIntSlot(x)
	case []int8: // byte[]
		s[offset] = int8(x)
	default:
		panic("putByte!")
	}
}

// public native char getChar(Object o, long offset);
// (Ljava/lang/Object;J)C
func getChar(frame *rtda.Frame) {
	obj, offset := _getArgs(frame)
	switch s := obj.Fields.(type) {
	case []heap.Slot: // object
		frame.PushInt(s[offset].IntValue())
	case []uint16: // char[]
		frame.PushInt(int32(s[offset]))
	default:
		panic("getChar!")
	}
}

// public native void putChar(Object o, long offset, char x);
// (Ljava/lang/Object;JC)V
func putChar(frame *rtda.Frame) {
	obj, offset := _getArgs(frame)
	x := frame.GetIntVar(4)
	switch s := obj.Fields.(type) {
	case []heap.Slot: // object
		s[offset] = heap.NewIntSlot(x)
	case []uint16: // char[]
		s[offset] = uint16(x)
	default:
		panic("putChar!")
	}
}

// public native short getShort(Object o, long offset);
// (Ljava/lang/Object;J)S
func getShort(frame *rtda.Frame) {
	obj, offset := _getArgs(frame)
	switch s := obj.Fields.(type) {
	case []heap.Slot: // object
		frame.PushInt(s[offset].IntValue()) // TODO
	case []int16: // short[]
		frame.PushInt(int32(s[offset]))
	case []int8: // byte[]
		shorts := vmutils.CastInt8sToUint16s(s[offset:])
		frame.PushInt(int32(shorts[0]))
	default:
		panic("getShort!")
	}
}

// public native void putShort(Object o, long offset, short x);
// (Ljava/lang/Object;JS)V
func putShort(frame *rtda.Frame) {
	obj, offset := _getArgs(frame)
	x := frame.GetIntVar(4)
	switch s := obj.Fields.(type) {
	case []heap.Slot: // object
		s[offset] = heap.NewIntSlot(x)
	case []int16: // short[]
		s[offset] = int16(x)
	default:
		panic("putShort!")
	}
}

// public native int getInt(Object o, long offset);
// (Ljava/lang/Object;J)I
func getInt(frame *rtda.Frame) {
	obj, offset := _getArgs(frame)
	switch s := obj.Fields.(type) {
	case []heap.Slot: // object
		frame.PushInt(s[offset].IntValue())
	case []int32: // int[]
		frame.PushInt(int32(s[offset]))
	default:
		panic("getInt!")
	}
}

// public native void putInt(Object o, long offset, int x);
// (Ljava/lang/Object;JI)V
func putInt(frame *rtda.Frame) {
	obj, offset := _getArgs(frame)
	x := frame.GetIntVar(4)
	switch s := obj.Fields.(type) {
	case []heap.Slot: // object
		s[offset] = heap.NewIntSlot(x)
	case []int32: // int[]
		s[offset] = x
	default:
		panic("putInt!")
	}
}

// public native long getLong(Object o, long offset);
// (Ljava/lang/Object;J)J
func getLong(frame *rtda.Frame) {
	obj, offset := _getArgs(frame)
	switch s := obj.Fields.(type) {
	case []heap.Slot: // object
		frame.PushLong(s[offset].LongValue())
	case []int64: // long[]
		frame.PushLong(s[offset])
	default:
		panic("getLong!")
	}
}

// public native void putLong(Object o, long offset, long x);
// (Ljava/lang/Object;JJ)V
func putLong(frame *rtda.Frame) {
	obj, offset := _getArgs(frame)
	x := frame.GetLongVar(4)
	switch s := obj.Fields.(type) {
	case []heap.Slot: // object
		s[offset] = heap.NewLongSlot(x)
	case []int64: // long[]
		s[offset] = x
	default:
		panic("putLong!")
	}
}

// public native float getFloat(Object o, long offset);
// (Ljava/lang/Object;J)F
func getFloat(frame *rtda.Frame) {
	obj, offset := _getArgs(frame)
	switch s := obj.Fields.(type) {
	case []heap.Slot: // object
		frame.PushFloat(s[offset].FloatValue())
	case []float32: // float[]
		frame.PushFloat(s[offset])
	default:
		panic("getFloat!")
	}
}

// public native void putFloat(Object o, long offset, float x);
// (Ljava/lang/Object;JF)V
func putFloat(frame *rtda.Frame) {
	obj, offset := _getArgs(frame)
	x := frame.GetFloatVar(4)
	switch s := obj.Fields.(type) {
	case []heap.Slot: // object
		s[offset] = heap.NewFloatSlot(x)
	case []float32: // float[]
		s[offset] = x
	default:
		panic("putFloat!")
	}
}

// public native double getDouble(Object o, long offset);
// (Ljava/lang/Object;J)D
func getDouble(frame *rtda.Frame) {
	obj, offset := _getArgs(frame)
	switch s := obj.Fields.(type) {
	case []heap.Slot: // object
		frame.PushDouble(s[offset].DoubleValue())
	case []float64: // double[]
		frame.PushDouble(s[offset])
	default:
		panic("getDouble!")
	}
}

// public native void putDouble(Object o, long offset, double x);
// (Ljava/lang/Object;JD)V
func putDouble(frame *rtda.Frame) {
	obj, offset := _getArgs(frame)
	x := frame.GetDoubleVar(4)
	switch s := obj.Fields.(type) {
	case []heap.Slot: // object
		s[offset] = heap.NewDoubleSlot(x)
	case []float64: // double[]
		s[offset] = x
	default:
		panic("putDouble!")
	}
}

// public native Object getReference(Object o, long offset);
// (Ljava/lang/Object;J)Ljava/lang/Object;
func getReference(frame *rtda.Frame) {
	obj, offset := _getArgs(frame)
	switch s := obj.Fields.(type) {
	case []heap.Slot: // object
		frame.PushRef(s[offset].Ref)
	case []*heap.Object: // ref[]
		frame.PushRef(s[offset])
	default:
		panic("getReference!")
	}
}

// public native void putReference(Object o, long offset, Object x);
// (Ljava/lang/Object;JLjava/lang/Object;)V
func putReference(frame *rtda.Frame) {
	obj, offset := _getArgs(frame)
	x := frame.GetRefVar(4)
	switch s := obj.Fields.(type) {
	case []heap.Slot: // object
		s[offset] = heap.NewRefSlot(x)
	case []*heap.Object: // ref[]
		s[offset] = x
	default:
		panic("putReference!")
	}
}

func _getArgs(frame *rtda.Frame) (obj *heap.Object, offset int64) {
	obj = frame.GetRefVar(1)
	offset = frame.GetLongVar(2)
	return
}
