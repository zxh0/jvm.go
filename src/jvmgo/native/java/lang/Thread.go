package lang

import (
    . "jvmgo/any"
    "jvmgo/jvm/interpreter"
    "jvmgo/jvm/rtda"
    rtc "jvmgo/jvm/rtda/class"
)

func init() {
    _thread(currentThread,  "currentThread",    "()Ljava/lang/Thread;")
    _thread(isAlive,        "isAlive",          "()Z")
    _thread(setPriority0,   "setPriority0",     "(I)V")
    _thread(start0,         "start0",           "()V")
}

func _thread(method Any, name, desc string) {
    rtc.RegisterNativeMethod("java/lang/Thread", name, desc, method)
}

// @Deprecated public native int countStackFrames();
// private native boolean isInterrupted(boolean ClearInterrupted);
// public static native boolean holdsLock(Object obj);
// private native static StackTraceElement[][] dumpThreads(Thread[] threads);
// private native static Thread[] getThreads();
// private native void setPriority0(int newPriority);
// private native void stop0(Object o);
// private native void suspend0();
// private native void resume0();
// private native void interrupt0();
// private native void setNativeName(String name);

// public static native Thread currentThread();
// ()Ljava/lang/Thread;
func currentThread(frame *rtda.Frame) {
    jThread := frame.Thread().JThread()
    frame.OperandStack().PushRef(jThread)
}

// public final native boolean isAlive();
// ()Z
func isAlive(frame *rtda.Frame) {
    // todo
    stack := frame.OperandStack()
    stack.PushBoolean(false)
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

    runMethod := this.Class().GetInstanceMethod("run", "()V")
    newThread := rtda.NewThread(this)
    newFrame := newThread.NewFrame(runMethod)
    newThread.PushFrame(newFrame)
    interpreter.Loop(newThread)
}
