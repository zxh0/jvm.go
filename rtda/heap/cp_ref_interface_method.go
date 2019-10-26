package heap

import (
	"github.com/zxh0/jvm.go/classfile"
)

type ConstantInterfaceMethodRef struct {
	ConstantMethodRef
}

func newConstantInterfaceMethodRef(class *Class, cf *classfile.ClassFile,
	cfRef classfile.ConstantInterfaceMethodRefInfo) *ConstantInterfaceMethodRef {

	ref := &ConstantInterfaceMethodRef{}
	ref.ConstantMemberRef = newConstantMemberRef(class, cf, cfRef.ClassIndex, cfRef.NameAndTypeIndex)
	ref.ParamSlotCount = calcParamSlotCount(ref.descriptor)
	return ref
}

// todo
func (ref *ConstantInterfaceMethodRef) FindInterfaceMethod(obj *Object) *Method {
	for class := obj.Class; class != nil; class = class.SuperClass {
		method := class.getMethod(ref.name, ref.descriptor, false)
		if method != nil {
			return method
		}
	}

	if method := findInterfaceMethod(obj.Class.Interfaces, ref.name, ref.descriptor); method != nil {
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
