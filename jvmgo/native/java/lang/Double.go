package lang

import (
	"math"

	"github.com/zxh0/jvm.go/jvmgo/rtda"
	"github.com/zxh0/jvm.go/jvmgo/rtda/heap"
)

func init() {
	_double(doubleToRawLongBits, "doubleToRawLongBits", "(D)J")
	_double(longBitsToDouble, "longBitsToDouble", "(J)D")
}

func _double(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("java/lang/Double", name, desc, method)
}

// public static native long doubleToRawLongBits(double value);
// (D)J
func doubleToRawLongBits(frame *rtda.Frame) {
	vars := frame.LocalVars()
	value := vars.GetDouble(0)

	// todo
	bits := math.Float64bits(value)
	stack := frame.OperandStack()
	stack.PushLong(int64(bits))
}

// public static native double longBitsToDouble(long bits);
// (J)D
func longBitsToDouble(frame *rtda.Frame) {
	vars := frame.LocalVars()
	bits := vars.GetLong(0)

	// todo
	value := math.Float64frombits(uint64(bits))
	stack := frame.OperandStack()
	stack.PushDouble(value)
}
