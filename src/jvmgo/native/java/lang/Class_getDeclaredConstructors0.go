package lang

import (
    "jvmgo/jvm/rtda"
    rtc "jvmgo/jvm/rtda/class"
)

/*
Constructor(Class<T> declaringClass,
            Class<?>[] parameterTypes,
            Class<?>[] checkedExceptions,
            int modifiers,
            int slot,
            String signature,
            byte[] annotations,
            byte[] parameterAnnotations)
}
*/
const _constructorConstructorDescriptor = 
        "(Ljava/lang/Class;" +
        "[Ljava/lang/Class;" +
        "[Ljava/lang/Class;" +
        "II" +
        "Ljava/lang/String;" +
        "[B[B)V"

// private native Constructor<T>[] getDeclaredConstructors0(boolean publicOnly);
// (Z)[Ljava/lang/reflect/Constructor;
func getDeclaredConstructors0(frame *rtda.Frame) {
    vars := frame.LocalVars()
    jClass := vars.GetRef(0) // this
    publicOnly := vars.GetBoolean(1)

    goClass := jClass.Extra().(*rtc.Class)
    goConstructors := goClass.GetConstructors(publicOnly)
    constructorCount := uint(len(goConstructors))
    
    constructorClass := goClass.ClassLoader().LoadClass("java/lang/reflect/Constructor")
    constructorInitMethod := constructorClass.GetConstructor(_constructorConstructorDescriptor)
    constructorArr := constructorClass.NewArray(constructorCount)
    stack := frame.OperandStack()
    stack.PushRef(constructorArr)

    if constructorCount > 0 {
        constructorObjs := constructorArr.Fields().([]*rtc.Obj)
        thread := frame.Thread()
        for i, goConstructor := range goConstructors {
            constructorObj := constructorClass.NewObjWithExtra(goConstructor)
            constructorObjs[i] = constructorObj
            // call <init>
            newFrame := thread.NewFrame(constructorInitMethod)
            vars := newFrame.LocalVars()
            vars.SetRef(0, constructorObj) // this
            vars.SetRef(1, jClass) // declaringClass
            vars.SetRef(2, getParameterTypeArr(goConstructor)) // parameterTypes
            vars.SetRef(3, nil) // todo checkedExceptions
            vars.SetInt(4, int32(goConstructor.GetAccessFlags())) // modifiers
            vars.SetInt(5, int32(0)) // todo slot
            vars.SetRef(6, nil) // todo signature
            vars.SetRef(7, nil) // todo annotations
            vars.SetRef(8, nil) // todo parameterAnnotations
            thread.PushFrame(newFrame)
        }
    }
}
