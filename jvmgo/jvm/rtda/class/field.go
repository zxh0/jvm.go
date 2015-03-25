package class

import (
	. "github.com/zxh0/jvm.go/jvmgo/any"
	cf "github.com/zxh0/jvm.go/jvmgo/classfile"
)

type Field struct {
	ClassMember
	IsLongOrDouble bool
	slot           uint
	_type          *Class
}

func newField(class *Class, fieldInfo *cf.FieldInfo) *Field {
	field := &Field{}
	field.class = class
	field.accessFlags = fieldInfo.AccessFlags()
	field.name = fieldInfo.Name()
	field.descriptor = fieldInfo.Descriptor()
	field.signature = fieldInfo.Signature()
	field.IsLongOrDouble = (field.descriptor == "J" || field.descriptor == "D")
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
	descriptor := self.descriptor
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
