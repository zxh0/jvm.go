package lang

import (
    //"fmt"
    "strings"
    . "jvmgo/any"
    "jvmgo/jvm/rtda"
    rtc "jvmgo/jvm/rtda/class"
)

func init() {
    _class(desiredAssertionStatus0,     "desiredAssertionStatus0",  "(Ljava/lang/Class;)Z")
    _class(forName0,                    "forName0",                 "(Ljava/lang/String;ZLjava/lang/ClassLoader;)Ljava/lang/Class;")
    _class(getClassLoader0,             "getClassLoader0",          "()Ljava/lang/ClassLoader;")
    _class(getDeclaredConstructors0,    "getDeclaredConstructors0", "(Z)[Ljava/lang/reflect/Constructor;")
    _class(getDeclaredFields0,          "getDeclaredFields0",       "(Z)[Ljava/lang/reflect/Field;")
    _class(getInterfaces,               "getInterfaces",            "()[Ljava/lang/Class;")
    _class(getName0,                    "getName0",                 "()Ljava/lang/String;")
    _class(getPrimitiveClass,           "getPrimitiveClass",        "(Ljava/lang/String;)Ljava/lang/Class;")
    _class(getSuperclass,               "getSuperclass",            "()Ljava/lang/Class;")
    _class(isInterface,                 "isInterface",              "()Z")
    _class(isPrimitive,                 "isPrimitive",              "()Z")
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

// private static native Class<?> forName0(String name, boolean initialize, ClassLoader loader) throws ClassNotFoundException;
// (Ljava/lang/String;ZLjava/lang/ClassLoader;)Ljava/lang/Class;
func forName0(frame *rtda.Frame) {
    stack := frame.OperandStack()
    jLoader := stack.PopRef()
    initialize := stack.PopBoolean()
    jName := stack.PopRef()

    goName := rtda.GoString(jName)
    goName = strings.Replace(goName, ".", "/", -1)
    goClass := frame.Method().Class().ClassLoader().LoadClass(goName)
    jClass := goClass.JClass()

    if initialize && goClass.InitializationNotStarted() {
        // undo forName0
        thread := frame.Thread()
        frame.SetNextPC(thread.PC())
        stack.PushRef(jName)
        stack.PushBoolean(initialize)
        stack.PushRef(jLoader)
        // init class
        rtda.InitClass(goClass, thread)
    } else {
        stack.PushRef(jClass)
    }
}

// native ClassLoader getClassLoader0();
// ()Ljava/lang/ClassLoader;
func getClassLoader0(frame *rtda.Frame) {
    // todo
    stack := frame.OperandStack()
    _ = stack.PopRef() // this
    stack.PushRef(nil)
}

// private native Field[] getDeclaredFields0(boolean publicOnly);
// (Z)[Ljava/lang/reflect/Field;
func getDeclaredFields0(frame *rtda.Frame) {
    stack := frame.OperandStack()
    publicOnly := stack.PopBoolean()
    jClass := stack.PopRef() // this
    goClass := jClass.Extra().(*rtc.Class)
    goFields := goClass.GetFields(publicOnly)

    classLoader := goClass.ClassLoader()
    fieldClass := classLoader.LoadClass("java/lang/reflect/Field")
    count := int32(len(goFields))
    fieldArr := rtc.NewRefArray(fieldClass, count)
    stack.PushRef(fieldArr)

    if count > 0 {
        /*
        Field(Class<?> declaringClass,
              String name,
              Class<?> type,
              int modifiers,
              int slot,
              String signature,
              byte[] annotations)
        */
        constructor := fieldClass.GetConstructor("(Ljava/lang/Class;Ljava/lang/String;Ljava/lang/Class;IILjava/lang/String;[B)V")
        jFields := fieldArr.Fields().([]*rtc.Obj)
        thread := frame.Thread()
        for i, goField := range goFields {
            jField := fieldClass.NewObj()
            jFields[i] = jField

            jName := rtda.NewJString(goField.Name(), frame)
            jType := goField.Type().JClass()

            newFrame := thread.NewFrame(constructor)
            vars := newFrame.LocalVars()
            vars.SetRef(0, jField) // this
            vars.SetRef(1, jClass) // declaringClass
            vars.SetRef(2, jName) // name
            vars.SetRef(3, jType) // type
            vars.SetInt(4, int32(goField.GetAccessFlags())) // modifiers
            vars.SetInt(5, int32(goField.Slot())) // slot
            vars.SetRef(6, nil) // todo signature
            vars.SetRef(7, nil) // todo annotations
            thread.PushFrame(newFrame)
        }

        //panic("todo getDeclaredFields0")
    }
}

// private native Method[]      getDeclaredMethods0(boolean publicOnly);
// private native Class<?>[]   getDeclaredClasses0();

// private native String getName0();
// ()Ljava/lang/String;
func getName0(frame *rtda.Frame) {
    stack := frame.OperandStack()
    jClass := stack.PopRef() // this
    goClass := jClass.Extra().(*rtc.Class)
    goName := goClass.Name()
    jName := rtda.NewJString(goName, frame)
    stack.PushRef(jName)
}

// static native Class<?> getPrimitiveClass(String name);
// (Ljava/lang/String;)Ljava/lang/Class;
func getPrimitiveClass(frame *rtda.Frame) {
    stack := frame.OperandStack()
    jName := stack.PopRef()
    chars := rtda.JStringChars(jName)

    classLoader := frame.Method().Class().ClassLoader()
    jClass := _getPrimitiveClass(chars, classLoader).JClass()
    stack.PushRef(jClass)
}

func _getPrimitiveClass(name []uint16, classLoader *rtc.ClassLoader) (*rtc.Class) {
    switch name[0] {
    case 'b':
        switch name[1] {
            case 'o': return classLoader.GetPrimitiveClass("boolean")
            case 'y': return classLoader.GetPrimitiveClass("byte")
        }
    case 'c': return classLoader.GetPrimitiveClass("char")
    case 'd': return classLoader.GetPrimitiveClass("double")
    case 'f': return classLoader.GetPrimitiveClass("float")
    case 'i': return classLoader.GetPrimitiveClass("int")
    case 'l': return classLoader.GetPrimitiveClass("long")
    case 's': return classLoader.GetPrimitiveClass("short")
    case 'v': return classLoader.GetPrimitiveClass("void")
    }
    panic("BAD primitive type!") // todo
}

// private native Class<?>[] getInterfaces();
// ()[Ljava/lang/Class;
func getInterfaces(frame *rtda.Frame) {
    stack := frame.OperandStack()
    jClass := stack.PopRef() // this
    goClass := jClass.Extra().(*rtc.Class)
    goInterfaces := goClass.Interfaces()
    jInterfaces := make([]*rtc.Obj, len(goInterfaces))

    for i, goInterface := range goInterfaces {
        jInterfaces[i] = goInterface.JClass()
    }

    classLoader := goClass.ClassLoader()
    classClass := classLoader.LoadClass("java/lang/Class")

    interfaceArr := rtc.NewRefArray2(classClass, jInterfaces)
    stack.PushRef(interfaceArr)
}

// public native Class<? super T> getSuperclass();
// ()Ljava/lang/Class;
func getSuperclass(frame *rtda.Frame) {
    stack := frame.OperandStack()
    jClass := stack.PopRef() // this
    goClass := jClass.Extra().(*rtc.Class)
    jSuperClass := goClass.SuperClass().JClass()
    stack.PushRef(jSuperClass)
}

// public native boolean isInterface();
// ()Z
func isInterface(frame *rtda.Frame) {
    stack := frame.OperandStack()
    jClass := stack.PopRef() // this
    goClass := jClass.Extra().(*rtc.Class)
    stack.PushBoolean(goClass.IsInterface())
}

// public native boolean isPrimitive();
// ()Z
func isPrimitive(frame *rtda.Frame) {
    stack := frame.OperandStack()
    jClass := stack.PopRef() // this
    goClass := jClass.Extra().(*rtc.Class)
    stack.PushBoolean(goClass.IsPrimitive())
}
