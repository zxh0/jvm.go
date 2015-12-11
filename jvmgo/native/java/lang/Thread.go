package lang

import (
	"github.com/zxh0/jvm.go/jvmgo/interpreter"
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	"github.com/zxh0/jvm.go/jvmgo/rtda/heap"
)

func init() {
	_thread(interrupt0, "interrupt0", "()V")
	_thread(isInterrupted, "isInterrupted", "(Z)Z")
	_thread(isAlive, "isAlive", "()Z")
	_thread(setPriority0, "setPriority0", "(I)V")
	_thread(start0, "start0", "()V")
}

func _thread(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("java/lang/Thread", name, desc, method)
}

// @Deprecated public native int countStackFrames();
// private native void stop0(Object o);
// private native void suspend0();
// private native void resume0();
// private native void setNativeName(String name);

// private native void interrupt0();
// ()V
func interrupt0(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()

	thread := _extraThread(this)
	thread.Interrupt()
}

// private native boolean isInterrupted(boolean ClearInterrupted);
// (Z)Z
func isInterrupted(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()
	clearInterrupted := vars.GetBoolean(1)

	// todo
	thread := _extraThread(this)
	interrupted := thread.IsInterrupted(clearInterrupted)

	stack := frame.OperandStack()
	stack.PushBoolean(interrupted)
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

func _extraThread(threadObj *heap.Object) *rtda.Thread {
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
