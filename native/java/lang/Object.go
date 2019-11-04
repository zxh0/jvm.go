package lang

import (
	"unsafe"

	"github.com/zxh0/jvm.go/native"
	"github.com/zxh0/jvm.go/rtda"
)

func init() {
	native.ForClass("java/lang/Object").
		Register(clone, "()Ljava/lang/Object;").
		Register(getClass, "()Ljava/lang/Class;").
		Register(hashCode, "()I").
		Register(notifyAll, "()V").
		Register(wait, "(J)V")
}

// protected native Object clone() throws CloneNotSupportedException;
// ()Ljava/lang/Object;
func clone(frame *rtda.Frame) {
	this := frame.GetThis()

	// todo
	frame.PushRef(this.Clone())
}

// public final native Class<?> getClass();
// ()Ljava/lang/Class;
func getClass(frame *rtda.Frame) {
	this := frame.GetThis()

	class := this.Class.JClass
	frame.PushRef(class)
}

// public native int hashCode();
// ()I
func hashCode(frame *rtda.Frame) {
	this := frame.GetThis()

	hash := int32(uintptr(unsafe.Pointer(this)))
	frame.PushInt(hash)
}

// public final native void notify();

// public final native void notifyAll();
// ()V
func notifyAll(frame *rtda.Frame) {
	this := frame.GetThis()

	thread := frame.Thread
	monitor := this.Monitor
	if !monitor.HasOwner(thread) {
		// todo
		panic("IllegalMonitorStateException")
	}

	monitor.NotifyAll()
}

// public final native void wait(long timeout) throws InterruptedException;
// (J)V
func wait(frame *rtda.Frame) {
	this := frame.GetThis()
	// timeout := frame.GetLongVar(1) // todo

	thread := frame.Thread
	monitor := this.Monitor
	if !monitor.HasOwner(thread) {
		// todo
		panic("IllegalMonitorStateException")
	}

	monitor.Wait()
}
