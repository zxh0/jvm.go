package lang

import (
	. "jvmgo/any"
	"jvmgo/jvm/rtda"
	rtc "jvmgo/jvm/rtda/class"
	"runtime"
)

func init() {
	_runtime(availableProcessors, "availableProcessors", "()I")
	_runtime(freeMemory, "freeMemory", "()J")
}

func _runtime(method Any, name, desc string) {
	rtc.RegisterNativeMethod("java/lang/Runtime", name, desc, method)
}

// public native int availableProcessors();
// ()I
func availableProcessors(frame *rtda.Frame) {
	numCPU := runtime.NumCPU()

	stack := frame.OperandStack()
	stack.PushInt(int32(numCPU))
}

// public native long freeMemory();
// ()J
func freeMemory(frame *rtda.Frame) {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	frees := memStats.Frees

	stack := frame.OperandStack()
	stack.PushLong(int64(frees))
}
