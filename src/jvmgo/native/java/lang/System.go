package lang

import (
    //"fmt"
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
    _system(setErr0,            "setErr0",              "(Ljava/io/PrintStream;)V")
    _system(setIn0,             "setIn0",               "(Ljava/io/InputStream;)V")
    _system(setOut0,            "setOut0",              "(Ljava/io/PrintStream;)V")
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
    // public synchronized Object setProperty(String key, String value)
    setPropMethod := props.Class().GetMethod("setProperty", "(Ljava/lang/String;Ljava/lang/String;)Ljava/lang/Object;")
    thread := frame.Thread()
    for key, val := range _props() {
        jKey := rtda.NewJString(key, frame)
        jVal := rtda.NewJString(val, frame)
        vars := thread.InvokeMethod2(setPropMethod)
        vars.SetRef(0, props) // this
        vars.SetRef(1, jKey)
        vars.SetRef(2, jVal)
    }
}

func _props() map[string]string {
    return map[string]string{
        "file.encoding": "UTF-8",
        "sun.stdout.encoding": "UTF-8",
        "sun.stderr.encoding": "UTF-8",
    }
}

// public static native long nanoTime();
// ()J
func nanoTime(frame *rtda.Frame) {
    stack := frame.OperandStack()
    nanoTime := time.Now().UnixNano()
    stack.PushLong(nanoTime)
}

// private static native void setErr0(PrintStream err);
// (Ljava/io/PrintStream;)V
func setErr0(frame *rtda.Frame) {
    stack := frame.OperandStack()
    err := stack.PopRef()
    sysClass := frame.Method().Class()
    sysClass.SetStaticValue("err", "Ljava/io/PrintStream;", err)
}

// private static native void setIn0(InputStream in);
// (Ljava/io/InputStream;)V
func setIn0(frame *rtda.Frame) {
    stack := frame.OperandStack()
    in := stack.PopRef()
    sysClass := frame.Method().Class()
    sysClass.SetStaticValue("in", "Ljava/io/InputStream;", in)
}

// private static native void setOut0(PrintStream out);
// (Ljava/io/PrintStream;)V
func setOut0(frame *rtda.Frame) {
    stack := frame.OperandStack()
    out := stack.PopRef()
    sysClass := frame.Method().Class()
    sysClass.SetStaticValue("out", "Ljava/io/PrintStream;", out)
}
