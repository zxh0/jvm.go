package lang

import (
	"github.com/zxh0/jvm.go/jvmgo/jutil"
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	"github.com/zxh0/jvm.go/jvmgo/rtda/heap"
)

func getParameterTypeArr(method *heap.Method) *heap.Object {
	paramTypes := method.ParameterTypes()
	paramCount := len(paramTypes)

	classClass := heap.BootLoader().JLClassClass()
	classArr := classClass.NewArray(uint(paramCount))

	if paramCount > 0 {
		classObjs := classArr.Refs()
		for i, paramType := range paramTypes {
			classObjs[i] = paramType.JClass()
		}
	}

	return classArr
}

func getReturnType(method *heap.Method) *heap.Object {
	goReturnType := method.ReturnType()
	return goReturnType.JClass()
}

func getExceptionTypeArr(method *heap.Method) *heap.Object {
	exTypes := method.ExceptionTypes()
	exCount := len(exTypes)

	classClass := heap.BootLoader().JLClassClass()
	classArr := classClass.NewArray(uint(exCount))

	if exCount > 0 {
		classObjs := classArr.Refs()
		for i, exType := range exTypes {
			classObjs[i] = exType.JClass()
		}
	}

	return classArr
}

func getAnnotationByteArr(goBytes []byte) *heap.Object {
	if goBytes != nil {
		jBytes := jutil.CastUint8sToInt8s(goBytes)
		return heap.NewByteArray(jBytes)
	}
	return nil
}

func getSignatureStr(signature string) *heap.Object {
	if signature != "" {
		return rtda.JString(signature)
	}
	return nil
}
