package heap

import (
	"github.com/zxh0/jvm.go/classfile"
)

type ConstantMethodRef struct {
	ConstantMemberRef
	ParamSlotCount uint
	method         *Method
	vslot          int
}

func newConstantMethodRef(cf *classfile.ClassFile,
	cfRef classfile.ConstantMethodRefInfo) *ConstantMethodRef {

	ref := &ConstantMethodRef{vslot: -1}
	ref.init(cf, cfRef.ClassIndex, cfRef.NameAndTypeIndex)
	ref.ParamSlotCount = calcParamSlotCount(ref.descriptor)
	return ref
}

func (mr *ConstantMethodRef) GetMethod(static bool) *Method {
	if mr.method == nil {
		if static {
			mr.resolveStaticMethod()
		} else {
			mr.resolveSpecialMethod()
		}
	}
	return mr.method
}

func (mr *ConstantMethodRef) resolveStaticMethod() {
	method := mr.findMethod(true)
	if method != nil {
		mr.method = method
	} else {
		// todo
		panic("static method not found!")
	}
}

func (mr *ConstantMethodRef) resolveSpecialMethod() {
	method := mr.findMethod(false)
	if method != nil {
		mr.method = method
		return
	}

	// todo
	// class := mr.cp.class.classLoader.LoadClass(mr.className)
	// if class.IsInterface() {
	// 	method = mr.findMethodInInterfaces(class)
	// 	if method != nil {
	// 		mr.method = method
	// 		return
	// 	}
	// }

	// todo
	panic("special method not found!")
}

func (mr *ConstantMethodRef) findMethod(isStatic bool) *Method {
	class := bootLoader.LoadClass(mr.className)
	return class.getMethod(mr.name, mr.descriptor, isStatic)
}

// todo
/*func (mr *ConstantMethodref) findMethodInInterfaces(iface *Class) *Method {
	for _, m := range iface.methods {
		if !m.IsAbstract() {
			if m.name == mr.name && m.descriptor == mr.descriptor {
				return m
			}
		}
	}

	for _, superIface := range iface.interfaces {
		if m := mr.findMethodInInterfaces(superIface); m != nil {
			return m
		}
	}

	return nil
}*/

func (mr *ConstantMethodRef) GetVirtualMethod(ref *Object) *Method {
	if mr.vslot < 0 {
		mr.vslot = getVslot(ref.Class, mr.name, mr.descriptor)
	}
	return ref.Class.vtable[mr.vslot]
}
