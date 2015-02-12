package lang

import (
    rtc "jvmgo/jvm/rtda/class"
)

func getParameterTypeArr(method *rtc.Method) (*rtc.Obj) {
    goParamTypes := method.ParameterTypes()
    paramCount := uint(len(goParamTypes))

    classClass := method.Class().ClassLoader().JLClassClass()
    classArr := classClass.NewArray(paramCount)

    if paramCount > 0 {
        classObjs := classArr.Fields().([]*rtc.Obj)
        for i, goParamType := range goParamTypes {
            classObjs[i] = goParamType.JClass()
        }
    }

    return classArr
}

func getReturnType(method *rtc.Method) (*rtc.Obj) {
    goReturnType := method.ReturnType()
    return goReturnType.JClass()
}
