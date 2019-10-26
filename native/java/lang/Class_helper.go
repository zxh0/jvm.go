package lang

import (
	"github.com/zxh0/jvm.go/rtda/heap"
	"github.com/zxh0/jvm.go/vmutils"
)

func getParameterTypeArr(rt *heap.Runtime, method *heap.Method) *heap.Object {
	paramTypes := method.GetParameterTypes()
	paramCount := len(paramTypes)

	classClass := rt.BootLoader().JLClassClass()
	classArr := classClass.NewArray(uint(paramCount))

	if paramCount > 0 {
		classObjs := classArr.GetRefs()
		for i, paramType := range paramTypes {
			classObjs[i] = paramType.JClass
		}
	}

	return classArr
}

func getReturnType(method *heap.Method) *heap.Object {
	goReturnType := method.GetReturnType()
	return goReturnType.JClass
}

func getExceptionTypeArr(rt *heap.Runtime, method *heap.Method) *heap.Object {
	exTypes := method.GetExceptionTypes()
	exCount := len(exTypes)

	classClass := rt.BootLoader().JLClassClass()
	classArr := classClass.NewArray(uint(exCount))

	if exCount > 0 {
		classObjs := classArr.GetRefs()
		for i, exType := range exTypes {
			classObjs[i] = exType.JClass
		}
	}

	return classArr
}

func getAnnotationByteArr(rt *heap.Runtime, goBytes []byte) *heap.Object {
	if goBytes != nil {
		jBytes := vmutils.CastBytesToInt8s(goBytes)
		return rt.NewByteArray(jBytes)
	}
	return nil
}

func getSignatureStr(rt *heap.Runtime, signature string) *heap.Object {
	if signature != "" {
		return rt.JSFromGoStr(signature)
	}
	return nil
}
