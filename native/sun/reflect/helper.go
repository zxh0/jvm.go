package reflect

import (
	"github.com/zxh0/jvm.go/native/box"
	"github.com/zxh0/jvm.go/rtda/heap"
)

func getGoMethod(methodObj *heap.Object) *heap.Method {
	return _getGoMethod(methodObj, false)
}
func getGoConstructor(constructorObj *heap.Object) *heap.Method {
	return _getGoMethod(constructorObj, true)
}
func _getGoMethod(methodObj *heap.Object, isConstructor bool) *heap.Method {
	extra := methodObj.Extra
	if extra != nil {
		return extra.(*heap.Method)
	}

	if isConstructor {
		root := methodObj.GetFieldValue("root", "Ljava/lang/reflect/Constructor;").Ref
		return root.Extra.(*heap.Method)
	} else {
		root := methodObj.GetFieldValue("root", "Ljava/lang/reflect/Method;").Ref
		return root.Extra.(*heap.Method)
	}
}

// Object[] -> []Slot
func convertArgs(this, argArr *heap.Object, method *heap.Method) []heap.Slot {
	if method.ArgSlotCount == 0 {
		return nil
	}
	if method.ArgSlotCount == 1 && !method.IsStatic() {
		return []heap.Slot{heap.NewRefSlot(this)}
	}

	argObjs := argArr.Refs()
	argTypes := method.ParsedDescriptor.ParameterTypes

	args := make([]heap.Slot, method.ArgSlotCount)
	j := 0
	if !method.IsStatic() {
		args[0] = heap.NewRefSlot(this)
		j = 1
	}

	for i, argType := range argTypes {
		argObj := argObjs[i]

		if argType.IsBaseType() {
			// todo
			unboxed := box.Unbox(argObj, string(argType))
			args[i+j] = unboxed
			if argType.IsLongOrDouble() {
				j++
			}
		} else {
			args[i+j] = heap.NewRefSlot(argObj)
		}
	}

	return args
}
