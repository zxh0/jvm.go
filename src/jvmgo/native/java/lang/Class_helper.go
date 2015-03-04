package lang

import (
	"jvmgo/jvm/rtda"
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

func getAnnotationByteArr(method *rtc.Method) *rtc.Obj {
	if bytes := method.AnnotationData(); bytes != nil {
		return rtc.NewByteArray(bytes, method.ClassLoader())
	}
	return nil
}

func getMethodSignature(method *rtc.Method) *rtc.Obj {
	if signature := method.Signature(); signature != "" {
		return rtda.NewJString(signature, method)
	}
	return nil
}
func getFieldSignature(field *rtc.Field) *rtc.Obj {
	if signature := field.Signature(); signature != "" {
		return rtda.NewJString(signature, field)
	}
	return nil
}
