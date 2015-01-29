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
    fields := self.class.obj.fields.([]Any)
    return fields[self.slot]
}
func (self *Field) PutStaticValue(val Any) {
    fields := self.class.obj.fields.([]Any)
    fields[self.slot] = val
}

func newField(class *Class, fieldInfo *cf.FieldInfo) (*Field) {
    field := &Field{}
    field.class = class
    field.accessFlags = fieldInfo.AccessFlags()
    field.name = fieldInfo.Name()
    field.descriptor = fieldInfo.Descriptor()
    return field
}
