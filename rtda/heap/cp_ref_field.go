package heap

import (
	"fmt"

	"github.com/zxh0/jvm.go/classfile"
)

type ConstantFieldRef struct {
	ConstantMemberRef
	field *Field
}

func newConstantFieldRef(cf *classfile.ClassFile,
	cfRef classfile.ConstantFieldRefInfo) *ConstantFieldRef {

	ref := &ConstantFieldRef{}
	ref.init(cf, cfRef.ClassIndex, cfRef.NameAndTypeIndex)
	return ref
}

func (fr *ConstantFieldRef) String() string {
	return fmt.Sprintf("{ConstantFieldref className:%v name:%v descriptor:%v}",
		fr.className, fr.name, fr.descriptor)
}

func (fr *ConstantFieldRef) GetField(static bool) *Field {
	if fr.field == nil {
		if static {
			fr.resolveStaticField()
		} else {
			fr.resolveInstanceField()
		}
	}
	return fr.field
}

func (fr *ConstantFieldRef) resolveInstanceField() {
	fromClass := bootLoader.LoadClass(fr.className)

	for class := fromClass; class != nil; class = class.SuperClass {
		field := class.getField(fr.name, fr.descriptor, false)
		if field != nil {
			fr.field = field
			return
		}
	}

	// todo
	panic(fmt.Errorf("instance field not found! %v", fr))
}

func (fr *ConstantFieldRef) resolveStaticField() {
	fromClass := bootLoader.LoadClass(fr.className)

	for class := fromClass; class != nil; class = class.SuperClass {
		field := class.getField(fr.name, fr.descriptor, true)
		if field != nil {
			fr.field = field
			return
		}
		if fr._findInterfaceField(class) {
			return
		}
	}

	// todo
	panic(fmt.Errorf("static field not found! %v", fr))
}

func (fr *ConstantFieldRef) _findInterfaceField(class *Class) bool {
	for _, iface := range class.Interfaces {
		for _, f := range iface.Fields {
			if f.Name == fr.name && f.Descriptor == fr.descriptor {
				fr.field = f
				return true
			}
		}
		if fr._findInterfaceField(iface) {
			return true
		}
	}
	return false
}
