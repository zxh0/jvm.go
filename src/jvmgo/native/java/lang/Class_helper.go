package lang

import (
	rtc "jvmgo/jvm/rtda/class"
)

func getParameterTypeArr(method *rtc.Method) *rtc.Obj {
	paramTypes := method.ParameterTypes()
	paramCount := len(paramTypes)

	classClass := method.Class().ClassLoader().JLClassClass()
	classArr := classClass.NewArray(uint(paramCount))

	if paramCount > 0 {
		classObjs := classArr.Fields().([]*rtc.Obj)
		for i, paramType := range paramTypes {
			classObjs[i] = paramType.JClass()
		}
	}

	return classArr
}

func getReturnType(method *rtc.Method) *rtc.Obj {
	goReturnType := method.ReturnType()
	return goReturnType.JClass()
}

func getExceptionTypeArr(method *rtc.Method) *rtc.Obj {
	exTypes := method.ExceptionTypes()
	exCount := len(exTypes)

	classClass := method.Class().ClassLoader().JLClassClass()
	classArr := classClass.NewArray(uint(exCount))

	if exCount > 0 {
		classObjs := classArr.Fields().([]*rtc.Obj)
		for i, exType := range exTypes {
			classObjs[i] = exType.JClass()
		}
	}

	return classArr
}
