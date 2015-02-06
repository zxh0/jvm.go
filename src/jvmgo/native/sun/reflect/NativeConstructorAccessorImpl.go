package reflect

import (
    //"unsafe"
    . "jvmgo/any"
    "jvmgo/jvm/rtda"
    rtc "jvmgo/jvm/rtda/class"
)

func init() {
    _ncai(newInstance0, "newInstance0", "(Ljava/lang/reflect/Constructor;[Ljava/lang/Object;)Ljava/lang/Object;")
}

func _ncai(method Any, name, desc string) {
    rtc.RegisterNativeMethod("sun/reflect/NativeConstructorAccessorImpl", name, desc, method)
}

// private static native Object newInstance0(Constructor<?> c, Object[] os)
// throws InstantiationException, IllegalArgumentException, InvocationTargetException;
// (Ljava/lang/reflect/Constructor;[Ljava/lang/Object;)Ljava/lang/Object;
func newInstance0(frame *rtda.Frame) {
    stack := frame.OperandStack()
    argArrObj := stack.PopRef()
    constructorObj := stack.PopRef()

    goConstructor := getExtra(constructorObj)
    goClass := goConstructor.Class()
    obj := goClass.NewObj()
    stack.PushRef(obj)

    // call <init>
    vars := frame.Thread().InvokeMethod2(goConstructor)
    vars.SetRef(0, obj) // this
    if goConstructor.ArgCount() > 0 {
        paramTypes := goConstructor.MethodDescriptor().ParameterTypes()
        argObjs := argArrObj.Fields().([]*rtc.Obj)
        for i, paramType := range paramTypes {
            argObj := argObjs[i]
            if paramType.IsBaseType() {
                // todo
                unboxed := unbox(argObj, paramType.Descriptor())
                vars.Set(uint(i + 1), unboxed)
            } else {
                vars.Set(uint(i + 1), argObj)
            }
        }
    }
}

func getExtra(constructorObj *rtc.Obj) (*rtc.Method) {
    extra := constructorObj.Extra()
    if extra != nil {
        return extra.(*rtc.Method)
    }

    root := constructorObj.GetFieldValue("root", "Ljava/lang/reflect/Constructor;").(*rtc.Obj)
    return root.Extra().(*rtc.Method)
}

func unbox(obj *rtc.Obj, descriptor string) Any {
    // todo
    return obj.GetFieldValue("value", descriptor)
}
