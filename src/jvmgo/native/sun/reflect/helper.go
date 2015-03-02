package reflect

import (
	. "jvmgo/any"
	rtc "jvmgo/jvm/rtda/class"
)

// Object[] -> []Any
func convertArgs(this, argArr *rtc.Obj, method *rtc.Method) []Any {
	if method.ArgCount() == 0 {
		if method.IsStatic() {
			return nil
		}
		return []Any{this}
	}

	argObjs := argArr.Fields().([]*rtc.Obj)
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
