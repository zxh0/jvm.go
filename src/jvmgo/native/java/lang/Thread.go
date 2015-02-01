package lang

import (
    . "jvmgo/any"
    "jvmgo/rtda"
    rtc "jvmgo/rtda/class"
)

func init() {
    _thread(currentThread,  "currentThread",    "()Ljava/lang/Thread;")
    _thread(setPriority0,   "setPriority0",     "(I)V")
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

// private native void setPriority0(int newPriority);
func setPriority0(frame *rtda.Frame) {
    stack := frame.OperandStack()
    /*newPriority := */stack.PopInt()
    /*this := */stack.PopRef()
    // todo
}
