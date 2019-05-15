package reflect

import (
	"github.com/zxh0/jvm.go/jvmgo/native/box"
	"github.com/zxh0/jvm.go/jvmgo/rtda/heap"
)

func getGoMethod(methodObj *heap.Object) *heap.Method {
	return _getGoMethod(methodObj, false)
}
func getGoConstructor(constructorObj *heap.Object) *heap.Method {
	return _getGoMethod(constructorObj, true)
}
func _getGoMethod(methodObj *heap.Object, isConstructor bool) *heap.Method {
	extra := methodObj.Extra()
	if extra != nil {
		return extra.(*heap.Method)
	}

	if isConstructor {
		root := methodObj.GetFieldValue("root", "Ljava/lang/reflect/Constructor;").(*heap.Object)
		return root.Extra().(*heap.Method)
	} else {
		root := methodObj.GetFieldValue("root", "Ljava/lang/reflect/Method;").(*heap.Object)
		return root.Extra().(*heap.Method)
	}
}

// Object[] -> []interface{}
func convertArgs(this, argArr *heap.Object, method *heap.Method) []interface{} {
	if method.ArgSlotCount() == 0 {
		return nil
	}
	if method.ArgSlotCount() == 1 && !method.IsStatic() {
		return []interface{}{this}
	}

	argObjs := argArr.Refs()
	argTypes := method.ParsedDescriptor().ParameterTypes()

	args := make([]interface{}, method.ArgSlotCount())
	j := 0
	if !method.IsStatic() {
		args[0] = this
		j = 1
	}

	for i, argType := range argTypes {
		argObj := argObjs[i]

		if argType.IsBaseType() {
			// todo
			unboxed := box.Unbox(argObj, argType.Descriptor())
			args[i+j] = unboxed
			if argType.IsLongOrDouble() {
				j++
			}
		} else {
			args[i+j] = argObj
		}
	}

	return args
}
