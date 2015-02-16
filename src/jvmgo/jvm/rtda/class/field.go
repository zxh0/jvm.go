package class

import (
	. "jvmgo/any"
	cf "jvmgo/classfile"
)

type Field struct {
	ClassMember
	slot  uint
	_type *Class
}

func newField(class *Class, fieldInfo *cf.FieldInfo) *Field {
	field := &Field{}
	field.class = class
	field.accessFlags = fieldInfo.AccessFlags()
	field.name = fieldInfo.Name()
	field.descriptor = fieldInfo.Descriptor()
	return field
}

func (self *Field) Slot() uint {
	return self.slot
}

func (self *Field) GetValue(ref *Obj) Any {
	fields := ref.fields.([]Any)
	return fields[self.slot]
}
func (self *Field) PutValue(ref *Obj, val Any) {
	fields := ref.fields.([]Any)
	fields[self.slot] = val
}

func (self *Field) GetStaticValue() Any {
	return self.class.staticFieldValues[self.slot]
}
func (self *Field) PutStaticValue(val Any) {
	self.class.staticFieldValues[self.slot] = val
}

func (self *Field) defaultValue() Any {
	switch self.descriptor[0] {
	case 'Z':
		return int32(0) // boolean
	case 'B':
		return int32(0) // byte
	case 'S':
		return int32(0) // short
	case 'C':
		return int32(0) // char
	case 'I':
		return int32(0) // int
	case 'J':
		return int64(0) // long
	case 'F':
		return float32(0) // float
	case 'D':
		return float64(0) // double
	case 'L':
		return nil // Object
	case '[':
		return nil // Array
	default:
		panic("BAD field descriptor: " + self.descriptor)
	}
}

// reflection
func (self *Field) Type() *Class {
	if self._type == nil {
		self._type = self.resolveType()
	}
	return self._type
}
func (self *Field) resolveType() *Class {
	classLoader := self.class.classLoader
	descriptor := self.descriptor
	if len(descriptor) == 1 {
		switch descriptor[0] {
		case 'B':
			return classLoader.GetPrimitiveClass("byte")
		case 'C':
			return classLoader.GetPrimitiveClass("char")
		case 'D':
			return classLoader.GetPrimitiveClass("double")
		case 'F':
			return classLoader.GetPrimitiveClass("float")
		case 'I':
			return classLoader.GetPrimitiveClass("int")
		case 'J':
			return classLoader.GetPrimitiveClass("long")
		case 'S':
			return classLoader.GetPrimitiveClass("short")
		case 'V':
			return classLoader.GetPrimitiveClass("void")
		case 'Z':
			return classLoader.GetPrimitiveClass("boolean")
		default:
			panic("BAD descriptor: " + descriptor)
		}
	}
	if descriptor[0] == 'L' {
		return classLoader.LoadClass(descriptor[1 : len(descriptor)-1])
	}
	return classLoader.LoadClass(descriptor)
}
