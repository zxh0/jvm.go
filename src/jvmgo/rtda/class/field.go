package class

import (
    . "jvmgo/any"
    cf "jvmgo/classfile"
)

type Field struct {
    ClassMember
    slot uint
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

func (self *Field) zeroValue() (Any) {
    return zeroValue(self.descriptor)
}

func newField(class *Class, fieldInfo *cf.FieldInfo) (*Field) {
    field := &Field{}
    field.class = class
    field.SetAccessFlags(fieldInfo.GetAccessFlags())
    field.name = fieldInfo.Name()
    field.descriptor = fieldInfo.Descriptor()
    return field
}
