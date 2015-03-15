package reflect

import (
	. "jvmgo/any"
	rtc "jvmgo/jvm/rtda/class"
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

// Object[] -> []Any
func convertArgs(this, argArr *rtc.Obj, method *rtc.Method) []Any {
	if method.ArgCount() == 0 {
		if method.IsStatic() {
			return nil
		} else {
			return []Any{this}
		}
	}

	argObjs := argArr.Refs()
	argTypes := method.ParsedDescriptor().ParameterTypes()

	args := make([]Any, len(argObjs)+1)
	args[0] = this
	for i, argType := range argTypes {
		argObj := argObjs[i]

		if argType.IsBaseType() {
			// todo
			unboxed := unbox(argObj, argType.Descriptor())
			args[i+1] = unboxed
		} else {
			args[i+1] = argObj
		}
	}

	if method.IsStatic() {
		return args[1:] // no this
	} else {
		return args
	}
}

func unbox(obj *rtc.Obj, descriptor string) Any {
	// todo
	return obj.GetFieldValue("value", descriptor)
}
