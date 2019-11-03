package lang

import (
	"runtime"

	"github.com/zxh0/jvm.go/native"
	"github.com/zxh0/jvm.go/rtda"
)

func init() {
	_runtime(availableProcessors, "availableProcessors", "()I")
	_runtime(freeMemory, "freeMemory", "()J")
}

func _runtime(method native.Method, name, desc string) {
	native.Register("java/lang/Runtime", name, desc, method)
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

// public native long totalMemory();
// public native long maxMemory();
// public native void gc();
// private static native void runFinalization0();
// public native void traceInstructions(boolean on);
// public native void traceMethodCalls(boolean on);
