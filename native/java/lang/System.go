package lang

import (
	"time"
	"unsafe"

	"github.com/zxh0/jvm.go/native"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

func init() {
	native.ForClass("java/lang/System").
		Register(arraycopy, "(Ljava/lang/Object;ILjava/lang/Object;II)V").
		Register(currentTimeMillis, "()J").
		Register(identityHashCode, "(Ljava/lang/Object;)I").
		Register(nanoTime, "()J").
		Register(setErr0, "(Ljava/io/PrintStream;)V").
		Register(setIn0, "(Ljava/io/InputStream;)V").
		Register(setOut0, "(Ljava/io/PrintStream;)V")
}

// public static native void arraycopy(Object src, int srcPos, Object dest, int destPos, int length)
// (Ljava/lang/Object;ILjava/lang/Object;II)V
func arraycopy(frame *rtda.Frame) {
	src := frame.GetRefVar(0)
	srcPos := frame.GetIntVar(1)
	dest := frame.GetRefVar(2)
	destPos := frame.GetIntVar(3)
	length := frame.GetIntVar(4)

	// NullPointerException
	if src == nil || dest == nil {
		panic("NPE") // todo
	}
	// ArrayStoreException
	if !checkArrayCopy(src, dest) {
		panic("ArrayStoreException")
	}
	// IndexOutOfBoundsException
	if srcPos < 0 || destPos < 0 || length < 0 ||
		srcPos+length > src.ArrayLength() ||
		destPos+length > dest.ArrayLength() {

		panic("IndexOutOfBoundsException") // todo
	}

	heap.ArrayCopy(src, dest, srcPos, destPos, length)
}

func checkArrayCopy(src, dest *heap.Object) bool {
	srcClass := src.Class
	destClass := dest.Class

	if !srcClass.IsArray() || !destClass.IsArray() {
		return false
	}
	if srcClass.IsPrimitiveArray() || destClass.IsPrimitiveArray() {
		return srcClass == destClass
	}
	return true
}

// public static native long currentTimeMillis();
// ()J
func currentTimeMillis(frame *rtda.Frame) {
	millis := time.Now().UnixNano() / int64(time.Millisecond)
	frame.PushLong(millis)
}

// public static native int identityHashCode(Object x);
// (Ljava/lang/Object;)I
func identityHashCode(frame *rtda.Frame) {
	ref := frame.GetRefVar(0)

	// todo
	hashCode := int32(uintptr(unsafe.Pointer(ref)))
	frame.PushInt(hashCode)
}

// public static native long nanoTime();
// ()J
func nanoTime(frame *rtda.Frame) {
	nanoTime := time.Now().UnixNano()
	frame.PushLong(nanoTime)
}

// private static native void setErr0(PrintStream err);
// (Ljava/io/PrintStream;)V
func setErr0(frame *rtda.Frame) {
	err := frame.GetRefVar(0) // TODO

	sysClass := frame.GetClass()
	sysClass.SetStaticValue("err", "Ljava/io/PrintStream;", heap.NewRefSlot(err))
}

// private static native void setIn0(InputStream in);
// (Ljava/io/InputStream;)V
func setIn0(frame *rtda.Frame) {
	in := frame.GetRefVar(0) // TODO

	sysClass := frame.GetClass()
	sysClass.SetStaticValue("in", "Ljava/io/InputStream;", heap.NewRefSlot(in))
}

// private static native void setOut0(PrintStream out);
// (Ljava/io/PrintStream;)V
func setOut0(frame *rtda.Frame) {
	out := frame.GetRefVar(0) // TODO

	sysClass := frame.GetClass()
	sysClass.SetStaticValue("out", "Ljava/io/PrintStream;", heap.NewRefSlot(out))
}
