package lang

import (
    "jvmgo/jvm/rtda"
    rtc "jvmgo/jvm/rtda/class"
)

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
    count := uint(len(goFields))
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
            jField := fieldClass.NewObjWithExtra(goField)
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
