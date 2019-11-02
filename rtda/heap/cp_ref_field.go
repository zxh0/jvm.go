package heap

import (
	"fmt"

	"github.com/zxh0/jvm.go/classfile"
)

type ConstantFieldRef struct {
	ConstantMemberRef
	resolved *Field
}

func newConstantFieldRef(class *Class, cf *classfile.ClassFile,
	cfRef classfile.ConstantFieldRefInfo) *ConstantFieldRef {

	ref := &ConstantFieldRef{}
	ref.ConstantMemberRef = newConstantMemberRef(class, cf, cfRef.ClassIndex, cfRef.NameAndTypeIndex)
	return ref
}

func (ref *ConstantFieldRef) String() string {
	return fmt.Sprintf("{ConstantFieldref className:%v name:%v descriptor:%v}",
		ref.className, ref.name, ref.descriptor)
}

func (ref *ConstantFieldRef) GetField(static bool) *Field {
	if ref.resolved == nil {
		if static {
			ref.resolveStaticField()
		} else {
			ref.resolveInstanceField()
		}
	}
	return ref.resolved
}

func (ref *ConstantFieldRef) resolveInstanceField() {
	fromClass := ref.getBootLoader().LoadClass(ref.className)

	for class := fromClass; class != nil; class = class.SuperClass {
		field := class.getField(ref.name, ref.descriptor, false)
		if field != nil {
			ref.resolved = field
			return
		}
	}

	// todo
	panic(fmt.Errorf("instance field not found! %v", ref))
}

func (ref *ConstantFieldRef) resolveStaticField() {
	fromClass := ref.getBootLoader().LoadClass(ref.className)

	for class := fromClass; class != nil; class = class.SuperClass {
		field := class.getField(ref.name, ref.descriptor, true)
		if field != nil {
			ref.resolved = field
			return
		}
		if ref._findInterfaceField(class) {
			return
		}
	}

	// todo
	panic(fmt.Errorf("static field not found! %v", ref))
}

func (ref *ConstantFieldRef) _findInterfaceField(class *Class) bool {
	for _, iface := range class.Interfaces {
		for _, f := range iface.Fields {
			if f.Name == ref.name && f.Descriptor == ref.descriptor {
				ref.resolved = f
				return true
			}
		}
		if ref._findInterfaceField(iface) {
			return true
		}
	}
	return false
}
