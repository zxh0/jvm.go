package reflect

import (
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
	"github.com/zxh0/jvm.go/jvmgo/native/box"
)

func getGoMethod(methodObj *rtc.Obj) *rtc.Method {
	return _getGoMethod(methodObj, false)
}
func getGoConstructor(constructorObj *rtc.Obj) *rtc.Method {
	return _getGoMethod(constructorObj, true)
}
func _getGoMethod(methodObj *rtc.Obj, isConstructor bool) *rtc.Method {
	extra := methodObj.Extra()
	if extra != nil {
		return extra.(*rtc.Method)
	}

	if isConstructor {
		root := methodObj.GetFieldValue("root", "Ljava/lang/reflect/Constructor;").(*rtc.Obj)
		return root.Extra().(*rtc.Method)
	} else {
		root := methodObj.GetFieldValue("root", "Ljava/lang/reflect/Method;").(*rtc.Obj)
		return root.Extra().(*rtc.Method)
	}
}

// Object[] -> []interface{}
func convertArgs(this, argArr *rtc.Obj, method *rtc.Method) []interface{} {
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
