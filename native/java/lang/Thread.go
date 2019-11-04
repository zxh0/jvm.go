package lang

import (
	"time"

	"github.com/zxh0/jvm.go/cpu"
	"github.com/zxh0/jvm.go/native"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

func init() {
	native.ForClass("java/lang/Thread").
		Register(interrupt0, "()V").
		Register(isInterrupted, "(Z)Z").
		Register(isAlive, "()Z").
		Register(setPriority0, "(I)V").
		Register(start0, "()V").
		Register(currentThread, "()Ljava/lang/Thread;").
		Register(sleep, "(J)V")
}

// @Deprecated public native int countStackFrames();
// private native void stop0(Object o);
// private native void suspend0();
// private native void resume0();
// private native void setNativeName(String name);

// private native void interrupt0();
// ()V
func interrupt0(frame *rtda.Frame) {
	this := frame.GetThis()

	thread := _extraThread(this)
	thread.Interrupt()
}

// private native boolean isInterrupted(boolean ClearInterrupted);
// (Z)Z
func isInterrupted(frame *rtda.Frame) {
	this := frame.GetThis()
	clearInterrupted := frame.GetBooleanVar(1)

	// todo
	thread := _extraThread(this)
	interrupted := thread.IsInterrupted(clearInterrupted)

	frame.PushBoolean(interrupted)
}

// public final native boolean isAlive();
// ()Z
func isAlive(frame *rtda.Frame) {
	this := frame.GetThis()

	thread := _extraThread(this)
	alive := thread != nil && !thread.IsStackEmpty()

	frame.PushBoolean(alive)
}

func _extraThread(threadObj *heap.Object) *rtda.Thread {
	threadObj.RLockState()
	defer threadObj.RUnlockState()

	extra := threadObj.Extra
	if extra == nil {
		return nil
	} else {
		return extra.(*rtda.Thread)
	}
}

// private native void setPriority0(int newPriority);
// (I)V
func setPriority0(frame *rtda.Frame) {
	// vars := frame.
	// this := frame.GetThis()
	// newPriority := frame.GetIntVar(1))
	// todo
}

// private native void start0();
// ()V
func start0(frame *rtda.Frame) {
	this := frame.GetThis()

	newThread := rtda.NewThread(this, frame.Thread.VMOptions, frame.Thread.Runtime)
	runMethod := this.Class.GetInstanceMethod("run", "()V")
	newFrame := newThread.NewFrame(runMethod)
	newFrame.SetRefVar(0, this)
	newThread.PushFrame(newFrame)

	this.LockState()
	this.Extra = newThread
	this.UnlockState()

	go cpu.Loop(newThread)
}

// public static native boolean holdsLock(Object obj);
// public static native void yield();
// private native static StackTraceElement[][] dumpThreads(Thread[] threads);
// private native static Thread[] getThreads();

// public static native Thread currentThread();
// ()Ljava/lang/Thread;
func currentThread(frame *rtda.Frame) {
	jThread := frame.Thread.JThread()
	frame.PushRef(jThread)
}

// public static native void sleep(long millis) throws InterruptedException;
// (J)V
func sleep(frame *rtda.Frame) {
	millis := frame.GetLongVar(0)

	thread := frame.Thread
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
