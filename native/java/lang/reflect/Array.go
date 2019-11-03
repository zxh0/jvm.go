package reflect

import (
	"github.com/zxh0/jvm.go/native"
	"github.com/zxh0/jvm.go/native/box"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

func init() {
	_array(get, "get", "(Ljava/lang/Object;I)Ljava/lang/Object;")
	_array(getLength, "getLength", "(Ljava/lang/Object;)I")
	_array(newArray, "newArray", "(Ljava/lang/Class;I)Ljava/lang/Object;")
	_array(set, "set", "(Ljava/lang/Object;ILjava/lang/Object;)V")
}

func _array(method native.Method, name, desc string) {
	native.Register("java/lang/reflect/Array", name, desc, method)
}

// public static native Object get(Object array, int index)
//         throws IllegalArgumentException, ArrayIndexOutOfBoundsException;
// (Ljava/lang/Object;I)Ljava/lang/Object;
func get(frame *rtda.Frame) {
	arr := frame.GetRefVar(0)
	index := frame.GetIntVar(1)

	if arr == nil {
		frame.Thread.ThrowNPE()
		return
	}
	if !arr.IsArray() {
		frame.Thread.ThrowIllegalArgumentException("Argument is not an array")
		return
	}
	if index < 0 || index >= arr.ArrayLength() {
		frame.Thread.ThrowArrayIndexOutOfBoundsExceptionNoMsg()
		return
	}

	if !arr.IsPrimitiveArray() {
		obj := arr.GetRefs()[index]
		frame.PushRef(obj)
		return
	}

	// primitive array
	primitiveDescriptor := arr.Class.Name[1]
	switch primitiveDescriptor {
	case 'Z':
		frame.PushBoolean(arr.GetBooleans()[index] == 1)
	case 'B':
		frame.PushInt(int32(arr.GetBytes()[index]))
	case 'C':
		frame.PushInt(int32(arr.GetChars()[index]))
	case 'S':
		frame.PushInt(int32(arr.GetShorts()[index]))
	case 'I':
		frame.PushInt(arr.GetInts()[index])
	case 'J':
		frame.PushLong(arr.GetLongs()[index])
	case 'F':
		frame.PushFloat(arr.GetFloats()[index])
	case 'D':
		frame.PushDouble(arr.GetDoubles()[index])
	}

	// boxing
	box.Box(frame, primitiveDescriptor)
}

// public static native void set(Object array, int index, Object value)
//        throws IllegalArgumentException, ArrayIndexOutOfBoundsException;
// (Ljava/lang/Object;ILjava/lang/Object;)V
func set(frame *rtda.Frame) {
	arr := frame.GetRefVar(0)
	index := frame.GetIntVar(1)
	value := frame.GetRefVar(2)

	if arr == nil {
		frame.Thread.ThrowNPE()
		return
	}
	if !arr.IsArray() {
		frame.Thread.ThrowIllegalArgumentException("Argument is not an array")
		return
	}

	if index < 0 || index >= arr.ArrayLength() {
		frame.Thread.ThrowArrayIndexOutOfBoundsExceptionNoMsg()
		return
	}

	if !arr.IsPrimitiveArray() {
		arr.GetRefs()[index] = value
		return
	}

	//TODO Consistent with the current need to determine whether the type and source type
	// Such as:
	// [I
	// java/lang/Integer
	// frame.Thread.ThrowIllegalArgumentException("argument type mismatch")

	// primitive array
	primitiveDescriptorStr := arr.Class.Name[1:]
	if primitiveDescriptorStr != heap.GetPrimitiveDescriptor(value.Class.Name) {
		frame.Thread.ThrowIllegalArgumentException("argument type mismatch")
		return
	}

	unboxed := box.Unbox(value, primitiveDescriptorStr)

	primitiveDescriptor := arr.Class.Name[1]
	switch primitiveDescriptor {
	case 'Z':
		arr.GetBooleans()[index] = int8(unboxed.IntValue())
	case 'B':
		arr.GetBytes()[index] = int8(unboxed.IntValue())
	case 'C':
		arr.GetChars()[index] = uint16(unboxed.IntValue())
	case 'S':
		arr.GetShorts()[index] = int16(unboxed.IntValue())
	case 'I':
		arr.GetInts()[index] = unboxed.IntValue()
	case 'J':
		arr.GetLongs()[index] = unboxed.LongValue()
	case 'F':
		arr.GetFloats()[index] = unboxed.FloatValue()
	case 'D':
		arr.GetDoubles()[index] = unboxed.DoubleValue()
	}
}

// public static native int getLength(Object array) throws IllegalArgumentException;
// (Ljava/lang/Object;)I
func getLength(frame *rtda.Frame) {
	arr := frame.GetRefVar(0)

	// todo IllegalArgumentException
	_len := arr.ArrayLength()
	frame.PushInt(_len)
}

// private static native Object newArray(Class<?> componentType, int length)
// throws NegativeArraySizeException;
// (Ljava/lang/Class;I)Ljava/lang/Object;
func newArray(frame *rtda.Frame) {
	componentType := frame.GetRefVar(0)
	length := frame.GetIntVar(1)
	if length < 0 {
		// todo
		panic("NegativeArraySizeException")
	}

	componentClass := componentType.GetGoClass()
	arrObj := componentClass.NewArray(uint(length))

	frame.PushRef(arrObj)
}
