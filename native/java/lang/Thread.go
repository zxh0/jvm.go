package lang

import (
	"github.com/zxh0/jvm.go/cpu"
	"github.com/zxh0/jvm.go/native"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

func init() {
	_thread(interrupt0, "interrupt0", "()V")
	_thread(isInterrupted, "isInterrupted", "(Z)Z")
	_thread(isAlive, "isAlive", "()Z")
	_thread(setPriority0, "setPriority0", "(I)V")
	_thread(start0, "start0", "()V")
}

func _thread(method native.Method, name, desc string) {
	native.Register("java/lang/Thread", name, desc, method)
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
