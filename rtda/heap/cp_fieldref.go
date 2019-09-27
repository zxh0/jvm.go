package heap

import (
	"fmt"
)

type ConstantFieldref struct {
	ConstantMemberref
	field *Field
}

func (fr *ConstantFieldref) String() string {
	return fmt.Sprintf("{ConstantFieldref className:%v name:%v descriptor:%v}",
		fr.className, fr.name, fr.descriptor)
}

func (fr *ConstantFieldref) InstanceField() *Field {
	if fr.field == nil {
		fr.resolveInstanceField()
	}
	return fr.field
}
func (fr *ConstantFieldref) resolveInstanceField() {
	fromClass := bootLoader.LoadClass(fr.className)

	for class := fromClass; class != nil; class = class.superClass {
		field := class.getField(fr.name, fr.descriptor, false)
		if field != nil {
			fr.field = field
			return
		}
	}

	// todo
	panic(fmt.Errorf("instance field not found! %v", fr))
}

func (fr *ConstantFieldref) StaticField() *Field {
	if fr.field == nil {
		fr.resolveStaticField()
	}
	return fr.field
}
func (fr *ConstantFieldref) resolveStaticField() {
	fromClass := bootLoader.LoadClass(fr.className)

	for class := fromClass; class != nil; class = class.superClass {
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

func (fr *ConstantFieldref) _findInterfaceField(class *Class) bool {
	for _, iface := range class.interfaces {
		for _, f := range iface.fields {
			if f.name == fr.name && f.descriptor == fr.descriptor {
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
