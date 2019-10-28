package heap

import (
	"fmt"
	"unicode/utf8"

	"github.com/zxh0/jvm.go/module"
	"github.com/zxh0/jvm.go/vmutils"
)

type Runtime struct {
	bootLoader *ClassLoader
	stringPool map[string]*Object // interned strings
}

func NewRuntime(mp module.Path, verboseClass bool) *Runtime {
	bl := newBootLoader(mp, verboseClass)
	rt := &Runtime{
		bootLoader: bl,
		stringPool: map[string]*Object{},
	}
	bl.rt = rt
	bl.init()
	return rt
}

func (rt *Runtime) BootLoader() *ClassLoader {
	return rt.bootLoader
}

/* string pool */

func (rt *Runtime) JSIntern(goStr string, jStr *Object) *Object {
	if internedStr, ok := rt.stringPool[goStr]; ok {
		return internedStr
	}

	rt.stringPool[goStr] = jStr
	return jStr
}

// go string -> java.lang.String
func (rt *Runtime) JSFromGoStr(goStr string) *Object {
	if internedStr, found := rt.stringPool[goStr]; found {
		return internedStr
	}

	jStr := rt.jsFromGoStr(goStr)
	jStr = rt.JSIntern(goStr, jStr) // TODO
	return jStr
}

func (rt *Runtime) jsFromGoStr(goStr string) *Object {
	var jBytes []int8
	var coder int32

	bytes := []byte(goStr)
	if len(bytes) == utf8.RuneCount(bytes) { // latin
		jBytes = vmutils.CastBytesToInt8s(bytes)
		coder = 0
	} else {
		uint16s := vmutils.UTF8ToUTF16(goStr)
		jBytes = vmutils.CastUint16sToInt8s(uint16s)
		coder = 1
	}

	jByteArr := rt.NewByteArray(jBytes)
	jStr := rt.bootLoader.JLStringClass().NewObj()
	jStr.SetFieldValue("value", "[B", NewRefSlot(jByteArr))
	jStr.SetFieldValue("coder", "B", NewIntSlot(coder))
	return jStr
}

/* array factory */

func (rt *Runtime) NewByteArray(bytes []int8) *Object {
	return newObj(rt.bootLoader.getClass("[B"), bytes, nil)
}
func (rt *Runtime) NewCharArray(chars []uint16) *Object {
	return newObj(rt.bootLoader.getClass("[C"), chars, nil)
}
func (rt *Runtime) NewObjectArray(objs []*Object) *Object {
	return newObj(rt.bootLoader.jlObjectClass.getArrayClass(), objs, nil)
}
func (rt *Runtime) NewStringArray(objs []*Object) *Object {
	return newObj(rt.bootLoader.jlStringClass.getArrayClass(), objs, nil)
}
func (rt *Runtime) NewClassArray(objs []*Object) *Object {
	return newObj(rt.bootLoader.jlClassClass.getArrayClass(), objs, nil)
}

func (rt *Runtime) NewPrimitiveArray(atype uint8, count uint) *Object {
	switch atype {
	case ATBoolean:
		return newPrimitiveArray(rt.bootLoader.getClass("[Z"), count)
	case ATByte:
		return newPrimitiveArray(rt.bootLoader.getClass("[B"), count)
	case ATChar:
		return newPrimitiveArray(rt.bootLoader.getClass("[C"), count)
	case ATShort:
		return newPrimitiveArray(rt.bootLoader.getClass("[S"), count)
	case ATInt:
		return newPrimitiveArray(rt.bootLoader.getClass("[I"), count)
	case ATLong:
		return newPrimitiveArray(rt.bootLoader.getClass("[J"), count)
	case ATFloat:
		return newPrimitiveArray(rt.bootLoader.getClass("[F"), count)
	case ATDouble:
		return newPrimitiveArray(rt.bootLoader.getClass("[D"), count)
	default:
		panic(fmt.Errorf("invalid atype: %v", atype))
	}
}
