package lang

import (
    "unsafe"
    . "jvmgo/any"
    "jvmgo/jvm/rtda"
    rtc "jvmgo/jvm/rtda/class"
)

func init() {
    _object(clone,      "clone",    "()Ljava/lang/Object;")
    _object(getClass,   "getClass", "()Ljava/lang/Class;")
    _object(hashCode,   "hashCode", "()I")
    _object(wait,       "wait",     "(J)V")
}

func _object(method Any, name, desc string) {
    rtc.RegisterNativeMethod("java/lang/Object", name, desc, method)
}

// protected native Object clone() throws CloneNotSupportedException;
// ()Ljava/lang/Object;
func clone(frame *rtda.Frame) {
    vars := frame.LocalVars()
    this := vars.GetThis()
    // todo
    stack := frame.OperandStack()
    stack.PushRef(this)
}

// public final native Class<?> getClass();
// ()Ljava/lang/Class;
func getClass(frame *rtda.Frame) {
    vars := frame.LocalVars()
    this := vars.GetThis()

    class := this.Class().JClass()
    stack := frame.OperandStack()
    stack.PushRef(class)
}

// public native int hashCode();
// ()I
func hashCode(frame *rtda.Frame) {
    vars := frame.LocalVars()
    this := vars.GetThis()

    hash := int32(uintptr(unsafe.Pointer(this)))
    stack := frame.OperandStack()
    stack.PushInt(hash)
}

// public final native void notify();
// public final native void notifyAll();

// public final native void wait(long timeout) throws InterruptedException;
// (J)V
func wait(frame *rtda.Frame) {
    vars := frame.LocalVars()
    this := vars.GetThis()
    // timeout := vars.GetLong(1) // todo

    thread := frame.Thread()
    monitor := this.Monitor()
    if !monitor.HasOwner(thread) {
        // todo
        panic("IllegalMonitorStateException")
    }
    
    monitor.Wait()
}
