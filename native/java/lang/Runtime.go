package lang

import (
	"math"
	"runtime"

	"github.com/zxh0/jvm.go/native"
	"github.com/zxh0/jvm.go/rtda"
)

func init() {
	native.ForClass("java/lang/Runtime").
		Register(availableProcessors, "()I").
		Register(freeMemory, "()J").
		Register(maxMemory, "()J")
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
