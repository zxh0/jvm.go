package heap

import (
	"fmt"

	"github.com/zxh0/jvm.go/classpath"
	"github.com/zxh0/jvm.go/vmutils"
)

type Runtime struct {
	bootLoader *ClassLoader
	stringPool map[string]*Object // interned strings
}

func NewRuntime(cp *classpath.ClassPath, verboseClass bool) *Runtime {
	bl := newBootLoader(cp, verboseClass)
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

	chars := vmutils.UTF8ToUTF16(goStr)
	jStr := rt.JSFromChars(chars)
	jStr = rt.JSIntern(goStr, jStr) // TODO
	return jStr
}

// java char[] -> java.lang.String
func (rt *Runtime) JSFromChars(chars []uint16) *Object {
	charArr := rt.NewCharArray(chars)
	jStr := rt.bootLoader.JLStringClass().NewObj()
	jStr.SetFieldValue("value", "[C", NewRefSlot(charArr))
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
