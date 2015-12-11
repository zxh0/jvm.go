package lang

import (
	"unsafe"

	"github.com/zxh0/jvm.go/jvmgo/rtda"
	"github.com/zxh0/jvm.go/jvmgo/rtda/heap"
)

func init() {
	_object(clone, "clone", "()Ljava/lang/Object;")
	_object(getClass, "getClass", "()Ljava/lang/Class;")
	_object(hashCode, "hashCode", "()I")
	_object(notifyAll, "notifyAll", "()V")
	_object(wait, "wait", "(J)V")
}

func _object(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("java/lang/Object", name, desc, method)
}

// protected native Object clone() throws CloneNotSupportedException;
// ()Ljava/lang/Object;
func clone(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()

	// todo
	stack := frame.OperandStack()
	stack.PushRef(this.Clone())
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
// ()V
func notifyAll(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()

	thread := frame.Thread()
	monitor := this.Monitor()
	if !monitor.HasOwner(thread) {
		// todo
		panic("IllegalMonitorStateException")
	}

	monitor.NotifyAll()
}

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
