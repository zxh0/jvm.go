package lang

import (
	"time"

	"github.com/zxh0/jvm.go/jvmgo/rtda"
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

	thread := frame.Thread()
	if millis < 0 {
		thread.ThrowIllegalArgumentException("timeout value is negative")
		return
	}

	m := millis * int64(time.Millisecond)
	d := time.Duration(m)
	interrupted := thread.Sleep(d)

	if interrupted {
		thread.ThrowInterruptedException("sleep interrupted")
	}
}
