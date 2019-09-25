package heap

import (
	cf "github.com/zxh0/jvm.go/classfile"
)

type Field struct {
	ClassMember
	IsLongOrDouble  bool
	constValueIndex uint16
	slotId          uint
	_type           *Class
}

func newField(class *Class, fieldInfo cf.MemberInfo) *Field {
	field := &Field{}
	field.class = class
	field.accessFlags = fieldInfo.AccessFlags
	field.name = fieldInfo.Name()
	field.descriptor = fieldInfo.Descriptor()
	field.signature = fieldInfo.Signature()
	field.IsLongOrDouble = field.descriptor == "J" || field.descriptor == "D"
	if kValAttr := fieldInfo.ConstantValueAttribute(); kValAttr != nil {
		field.constValueIndex = kValAttr.ConstantValueIndex
	}
	return field
}

func (field *Field) ConstValueIndex() uint16 {
	return field.constValueIndex
}
func (field *Field) SlotId() uint {
	return field.slotId
}

func (field *Field) GetValue(ref *Object) interface{} {
	fields := ref.fields.([]interface{})
	return fields[field.slotId]
}
func (field *Field) PutValue(ref *Object, val interface{}) {
	fields := ref.fields.([]interface{})
	fields[field.slotId] = val
}

func (field *Field) GetStaticValue() interface{} {
	return field.class.staticFieldSlots[field.slotId]
}
func (field *Field) PutStaticValue(val interface{}) {
	field.class.staticFieldSlots[field.slotId] = val
}

func (field *Field) defaultValue() interface{} {
	switch field.descriptor[0] {
	case 'Z': // boolean
		return int32(0)
	case 'B': // byte
		return int32(0)
	case 'S': // short
		return int32(0)
	case 'C': // char
		return int32(0)
	case 'I': // int
		return int32(0)
	case 'J': // long
		return int64(0)
	case 'F': // float
		return float32(0)
	case 'D': // double
		return float64(0)
	case 'L': // Object
		return nil
	case '[': // Array
		return nil
	default:
		panic("BAD field descriptor: " + field.descriptor)
	}
}

// reflection
func (field *Field) Type() *Class {
	if field._type == nil {
		field._type = field.resolveType()
	}
	return field._type
}
func (field *Field) resolveType() *Class {
	descriptor := field.descriptor
	if len(descriptor) == 1 {
		switch descriptor[0] {
		case 'B':
			return bootLoader.GetPrimitiveClass("byte")
		case 'C':
			return bootLoader.GetPrimitiveClass("char")
		case 'D':
			return bootLoader.GetPrimitiveClass("double")
		case 'F':
			return bootLoader.GetPrimitiveClass("float")
		case 'I':
			return bootLoader.GetPrimitiveClass("int")
		case 'J':
			return bootLoader.GetPrimitiveClass("long")
		case 'S':
			return bootLoader.GetPrimitiveClass("short")
		case 'V':
			return bootLoader.GetPrimitiveClass("void")
		case 'Z':
			return bootLoader.GetPrimitiveClass("boolean")
		default:
			panic("BAD descriptor: " + descriptor)
		}
	}
	if descriptor[0] == 'L' {
		return bootLoader.LoadClass(descriptor[1 : len(descriptor)-1])
	}
	return bootLoader.LoadClass(descriptor)
}
