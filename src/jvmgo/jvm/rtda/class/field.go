package class

import (
    . "jvmgo/any"
    cf "jvmgo/classfile"
)

type Field struct {
    ClassMember
    slot uint
}

func newField(class *Class, fieldInfo *cf.FieldInfo) (*Field) {
    field := &Field{}
    field.class = class
    field.SetAccessFlags(fieldInfo.GetAccessFlags())
    field.name = fieldInfo.Name()
    field.descriptor = fieldInfo.Descriptor()
    return field
}

func (self *Field) Slot() uint {
    return self.slot
}

func (self *Field) GetValue(ref *Obj) (Any) {
    fields := ref.fields.([]Any)
    return fields[self.slot]
}
func (self *Field) PutValue(ref *Obj, val Any) {
    fields := ref.fields.([]Any)
    fields[self.slot] = val
}

func (self *Field) GetStaticValue() (Any) {
    return self.class.staticFieldValues[self.slot]
}
func (self *Field) PutStaticValue(val Any) {
    self.class.staticFieldValues[self.slot] = val
}

func (self *Field) defaultValue() (Any) {
    switch self.descriptor[0] {
    case 'Z': return int32(0)   // boolean
    case 'B': return int32(0)   // byte
    case 'S': return int32(0)   // short
    case 'C': return int32(0)   // char
    case 'I': return int32(0)   // int
    case 'J': return int64(0)   // long
    case 'F': return float32(0) // float
    case 'D': return float64(0) // double
    case 'L': return nil        // Object
    case '[': return nil        // Array
    default: panic("BAD field descriptor: " + self.descriptor)
    }
}
