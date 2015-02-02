package lang

import (
    . "jvmgo/any"
    "jvmgo/rtda"
    rtc "jvmgo/rtda/class"
)

func init() {
    _class(desiredAssertionStatus0, "desiredAssertionStatus0",  "(Ljava/lang/Class;)Z")
    _class(getClassLoader0,         "getClassLoader0",          "()Ljava/lang/ClassLoader;")
    _class(getName0,                "getName0",                 "()Ljava/lang/String;")
    _class(getPrimitiveClass,       "getPrimitiveClass",        "(Ljava/lang/String;)Ljava/lang/Class;")
}

func _class(method Any, name, desc string) {
    rtc.RegisterNativeMethod("java/lang/Class", name, desc, method)
}

// private static native boolean desiredAssertionStatus0(Class<?> clazz);
// (Ljava/lang/Class;)Z
func desiredAssertionStatus0(frame *rtda.Frame) {
    // todo
    stack := frame.OperandStack()
    _ = stack.PopRef() // this
    stack.PushBoolean(false)
}

// native ClassLoader getClassLoader0();
// ()Ljava/lang/ClassLoader;
func getClassLoader0(frame *rtda.Frame) {
    // todo
    stack := frame.OperandStack()
    _ = stack.PopRef() // this
    stack.PushRef(nil)
}

// private native String getName0();
// ()Ljava/lang/String;
func getName0(frame *rtda.Frame) {
    stack := frame.OperandStack()
    this := stack.PopRef()
    goClass := this.Extra().(*rtc.Class)
    goName := goClass.Name()
    jName := rtda.NewJString(goName, frame.Thread())
    stack.PushRef(jName)
}

// static native Class<?> getPrimitiveClass(String name);
// (Ljava/lang/String;)Ljava/lang/Class;
func getPrimitiveClass(frame *rtda.Frame) {
    // todo
    stack := frame.OperandStack()
    _ = stack.PopRef() // name
    stack.PushRef(nil) // todo
}
