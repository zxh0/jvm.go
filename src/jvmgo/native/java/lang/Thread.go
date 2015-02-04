package lang

import (
    . "jvmgo/any"
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

// public static native Thread currentThread();
// ()Ljava/lang/Thread;
func currentThread(frame *rtda.Frame) {
    jThread := frame.Thread().JThread()
    frame.OperandStack().PushRef(jThread)
}

// public final native boolean isAlive();
// ()Z
func isAlive(frame *rtda.Frame) {
    stack := frame.OperandStack()
    stack.PopRef() // this
    stack.PushBoolean(false) // todo
}

// private native void setPriority0(int newPriority);
func setPriority0(frame *rtda.Frame) {
    stack := frame.OperandStack()
    /*newPriority := */stack.PopInt()
    /*this := */stack.PopRef()
    // todo
}

// private native void start0();
// ()V
func start0(frame *rtda.Frame) {
    stack := frame.OperandStack()
    stack.PopRef() // this
    // todo
}
