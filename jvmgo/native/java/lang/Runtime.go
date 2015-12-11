package lang

import (
	"runtime"

	"github.com/zxh0/jvm.go/jvmgo/rtda"
	"github.com/zxh0/jvm.go/jvmgo/rtda/heap"
)

func init() {
	_runtime(availableProcessors, "availableProcessors", "()I")
	_runtime(freeMemory, "freeMemory", "()J")
}

func _runtime(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("java/lang/Runtime", name, desc, method)
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

// public native long totalMemory();
// public native long maxMemory();
// public native void gc();
// private static native void runFinalization0();
// public native void traceInstructions(boolean on);
// public native void traceMethodCalls(boolean on);
