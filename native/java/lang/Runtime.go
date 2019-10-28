package lang

import (
	"math"
	"runtime"

	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

func init() {
	_runtime(availableProcessors, "availableProcessors", "()I")
	_runtime(freeMemory, "freeMemory", "()J")
	_runtime(maxMemory, "maxMemory", "()J")
}

func _runtime(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("java/lang/Runtime", name, desc, method)
}

// public native int availableProcessors();
// ()I
func availableProcessors(frame *rtda.Frame) {
	numCPU := runtime.NumCPU()

	frame.PushInt(int32(numCPU))
}

// public native long freeMemory();
// ()J
func freeMemory(frame *rtda.Frame) {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	frees := memStats.Frees

	frame.PushLong(int64(frees))
}

// public native long maxMemory();
// ()J
func maxMemory(frame *rtda.Frame) {
	frame.PushLong(math.MaxInt64)
}

// public native long totalMemory();
// public native void gc();
// private static native void runFinalization0();
// public native void traceInstructions(boolean on);
// public native void traceMethodCalls(boolean on);
