package heap

import (
	"github.com/zxh0/jvm.go/classfile"
)

type Field struct {
	ClassMember
	IsLongOrDouble  bool
	ConstValueIndex uint16
	SlotId          uint
	_type           *Class
}

func newField(class *Class, cf *classfile.ClassFile, cfMember classfile.MemberInfo) *Field {
	field := &Field{}
	field.Class = class
	field.copyMemberData(cf, cfMember)
	field.IsLongOrDouble = field.Descriptor == "J" || field.Descriptor == "D"
	field.ConstValueIndex = cfMember.GetConstantValueIndex()
	return field
}

func (field *Field) GetValue(ref *Object) Slot {
	fields := ref.Fields.([]Slot)
	return fields[field.SlotId]
}
func (field *Field) PutValue(ref *Object, val Slot) {
	fields := ref.Fields.([]Slot)
	fields[field.SlotId] = val
}

func (field *Field) GetStaticValue() Slot {
	return field.Class.StaticFieldSlots[field.SlotId]
}
func (field *Field) PutStaticValue(val Slot) {
	field.Class.StaticFieldSlots[field.SlotId] = val
}

// reflection
func (field *Field) Type() *Class {
	if field._type == nil {
		field._type = field.resolveType()
	}
	return field._type
}
func (field *Field) resolveType() *Class {
	bootLoader := field.Class.bootLoader
	className := getClassName(field.Descriptor)
	return bootLoader.LoadClass(className)
}
