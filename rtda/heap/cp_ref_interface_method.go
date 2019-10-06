package heap

import (
	"github.com/zxh0/jvm.go/classfile"
)

type ConstantInterfaceMethodRef struct {
	ConstantMethodRef
}

func newConstantInterfaceMethodRef(cf *classfile.ClassFile,
	cfRef classfile.ConstantInterfaceMethodRefInfo) *ConstantInterfaceMethodRef {

	ref := &ConstantInterfaceMethodRef{}
	ref.init(cf, cfRef.ClassIndex, cfRef.NameAndTypeIndex)
	ref.ArgSlotCount = calcArgSlotCount(ref.descriptor)
	return ref
}

// todo
func (imr *ConstantInterfaceMethodRef) FindInterfaceMethod(ref *Object) *Method {
	for class := ref.Class; class != nil; class = class.SuperClass {
		method := class.getMethod(imr.name, imr.descriptor, false)
		if method != nil {
			return method
		}
	}

	if method := findInterfaceMethod(ref.Class.Interfaces, imr.name, imr.descriptor); method != nil {
		return method
	} else {
		//TODO
		panic("virtual method not found!")
	}
}

func findInterfaceMethod(interfaces []*Class, name, descriptor string) *Method {
	for i := 0; i < len(interfaces); i++ {
		if method := findInterfaceMethod(interfaces[i].Interfaces, name, descriptor); method != nil {
			return method
		}
		method := interfaces[i].getMethod(name, descriptor, false)
		if method != nil {
			return method
		}
	}
	return nil
}
