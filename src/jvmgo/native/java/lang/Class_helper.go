package lang

import (
    rtc "jvmgo/jvm/rtda/class"
)

func getParameterTypeArr(method *rtc.Method) (*rtc.Obj) {
    goParamTypes := method.ParameterTypes()
    paramCount := uint(len(goParamTypes))

    classClass := method.Class().ClassLoader().LoadClass("java/lang/Class")
    classArr := classClass.NewArray(paramCount)

    if paramCount > 0 {
        classObjs := classArr.Fields().([]*rtc.Obj)
        for i, goParamType := range goParamTypes {
            classObjs[i] = goParamType.JClass()
        }
    }

    return classArr
}
