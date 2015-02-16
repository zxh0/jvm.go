package lang

import (
	. "jvmgo/any"
	"jvmgo/jvm/interpreter"
	"jvmgo/jvm/rtda"
	rtc "jvmgo/jvm/rtda/class"
	"time"
)

func init() {
	_thread(currentThread, "currentThread", "()Ljava/lang/Thread;")
	_thread(isAlive, "isAlive", "()Z")
	_thread(sleep, "sleep", "(J)V")
	_thread(setPriority0, "setPriority0", "(I)V")
	_thread(start0, "start0", "()V")
}

func _thread(method Any, name, desc string) {
	rtc.RegisterNativeMethod("java/lang/Thread", name, desc, method)
}

// @Deprecated public native int countStackFrames();
// private native boolean isInterrupted(boolean ClearInterrupted);
// public static native boolean holdsLock(Object obj);
// private native static StackTraceElement[][] dumpThreads(Thread[] threads);
// private native static Thread[] getThreads();
// private native void stop0(Object o);
// private native void suspend0();
// private native void resume0();
// private native void interrupt0();
// private native void setNativeName(String name);
// public static native void yield();

// public static native Thread currentThread();
// ()Ljava/lang/Thread;
func currentThread(frame *rtda.Frame) {
	jThread := frame.Thread().JThread()
	frame.OperandStack().PushRef(jThread)
}

// public final native boolean isAlive();
// ()Z
func isAlive(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()

	thread := _extraThread(this)
	alive := thread != nil && !thread.IsStackEmpty()

	stack := frame.OperandStack()
	stack.PushBoolean(alive)
}

func _extraThread(threadObj *rtc.Obj) *rtda.Thread {
	threadObj.RLockState()
	defer threadObj.RUnlockState()

	extra := threadObj.Extra()
	if extra == nil {
		return nil
	} else {
		return extra.(*rtda.Thread)
	}
}

// private native void setPriority0(int newPriority);
// (I)V
func setPriority0(frame *rtda.Frame) {
	// vars := frame.LocalVars()
	// this := vars.GetThis()
	// newPriority := vars.GetInt(1))
	// todo
}

// public static native void sleep(long millis) throws InterruptedException;
// (J)V
func sleep(frame *rtda.Frame) {
	vars := frame.LocalVars()
	millis := vars.GetLong(0)

	m := millis * int64(time.Millisecond)
	d := time.Duration(m)
	time.Sleep(d)
}

// private native void start0();
// ()V
func start0(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()

	newThread := rtda.NewThread(this)
	runMethod := this.Class().GetInstanceMethod("run", "()V")
	newFrame := newThread.NewFrame(runMethod)
	newFrame.LocalVars().SetRef(0, this)
	newThread.PushFrame(newFrame)

	this.LockState()
	this.SetExtra(newThread)
	this.UnlockState()

	go interpreter.Loop(newThread)
}
