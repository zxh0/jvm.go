package lang

import (
    "time"
    "unsafe"
    . "jvmgo/any"
    "jvmgo/jvm/rtda"
    rtc "jvmgo/jvm/rtda/class"
)

func init() {
    _system(arraycopy,          "arraycopy",            "(Ljava/lang/Object;ILjava/lang/Object;II)V")
    _system(currentTimeMillis,  "currentTimeMillis",    "()J")
    _system(identityHashCode,   "identityHashCode",     "(Ljava/lang/Object;)I")
    _system(initProperties,     "initProperties",       "(Ljava/util/Properties;)Ljava/util/Properties;")
    _system(nanoTime,           "nanoTime",             "()J")
}

func _system(method Any, name, desc string) {
    rtc.RegisterNativeMethod("java/lang/System", name, desc, method)
}

// public static native void arraycopy(Object src, int srcPos, Object dest, int destPos, int length)
// (Ljava/lang/Object;ILjava/lang/Object;II)V
func arraycopy(frame *rtda.Frame) {
    stack := frame.OperandStack()
    length := stack.PopInt()
    destPos := stack.PopInt()
    dest := stack.PopRef()
    srcPos := stack.PopInt()
    src := stack.PopRef()

    // NullPointerException
    if src == nil || dest == nil {
        panic("NPE") // todo
    }
    // ArrayStoreException
    if !rtc.HaveSameArrayType(src, dest) {
        panic("ArrayStoreException")
    }
    // IndexOutOfBoundsException
    if srcPos < 0 || destPos < 0 || length < 0 ||
            srcPos + length > rtc.ArrayLength(src) ||
            destPos + length > rtc.ArrayLength(dest) {

        panic("IndexOutOfBoundsException") // todo
    }

    rtc.ArrayCopy(src, dest, srcPos, destPos, length)
}

// public static native long currentTimeMillis();
// ()J
func currentTimeMillis(frame *rtda.Frame) {
    stack := frame.OperandStack()
    millis := time.Now().UnixNano() / 1000
    stack.PushLong(millis)
}

// public static native int identityHashCode(Object x);
// (Ljava/lang/Object;)I
func identityHashCode(frame *rtda.Frame) {
    // todo
    stack := frame.OperandStack()
    ref := stack.PopRef()
    hashCode := int32(uintptr(unsafe.Pointer(ref)))
    stack.PushInt(hashCode)
}

// private static native Properties initProperties(Properties props);
// (Ljava/util/Properties;)Ljava/util/Properties;
func initProperties(frame *rtda.Frame) {
    stack := frame.OperandStack()
    props := stack.PopRef()
    stack.PushRef(props)
    // todo
}

// public static native long nanoTime();
// ()J
func nanoTime(frame *rtda.Frame) {
    stack := frame.OperandStack()
    nanoTime := time.Now().UnixNano()
    stack.PushLong(nanoTime)
}

// private static native void setIn0(InputStream in);
// private static native void setOut0(PrintStream out);
// private static native void setErr0(PrintStream err);
