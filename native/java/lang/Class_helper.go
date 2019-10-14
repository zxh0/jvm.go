package lang

import (
	"github.com/zxh0/jvm.go/rtda/heap"
	"github.com/zxh0/jvm.go/vmutils"
)

func getParameterTypeArr(method *heap.Method) *heap.Object {
	paramTypes := method.GetParameterTypes()
	paramCount := len(paramTypes)

	classClass := heap.BootLoader().JLClassClass()
	classArr := classClass.NewArray(uint(paramCount))

	if paramCount > 0 {
		classObjs := classArr.Refs()
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

func getExceptionTypeArr(method *heap.Method) *heap.Object {
	exTypes := method.GetExceptionTypes()
	exCount := len(exTypes)

	classClass := heap.BootLoader().JLClassClass()
	classArr := classClass.NewArray(uint(exCount))

	if exCount > 0 {
		classObjs := classArr.Refs()
		for i, exType := range exTypes {
			classObjs[i] = exType.JClass
		}
	}

	return classArr
}

func getAnnotationByteArr(goBytes []byte) *heap.Object {
	if goBytes != nil {
		jBytes := vmutils.CastBytesToInt8s(goBytes)
		return heap.NewByteArray(jBytes)
	}
	return nil
}

func getSignatureStr(signature string) *heap.Object {
	if signature != "" {
		return heap.JSFromGoStr(signature)
	}
	return nil
}
