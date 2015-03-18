package lang

import (
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	"time"
)

func init() {
	_thread(currentThread, "currentThread", "()Ljava/lang/Thread;")
	_thread(sleep, "sleep", "(J)V")
}

// public static native boolean holdsLock(Object obj);
// public static native void yield();
// private native static StackTraceElement[][] dumpThreads(Thread[] threads);
// private native static Thread[] getThreads();

// public static native Thread currentThread();
// ()Ljava/lang/Thread;
func currentThread(frame *rtda.Frame) {
	jThread := frame.Thread().JThread()
	frame.OperandStack().PushRef(jThread)
}

// public static native void sleep(long millis) throws InterruptedException;
// (J)V
func sleep(frame *rtda.Frame) {
	vars := frame.LocalVars()
	millis := vars.GetLong(0)

	if millis < 0 {
		// IllegalArgumentException
	}

	m := millis * int64(time.Millisecond)
	d := time.Duration(m)
	time.Sleep(d)
}
