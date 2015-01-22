package class

import (
    . "jvmgo/any"
    cf "jvmgo/classfile"
)

type Field struct {
    accessFlags uint16
    name        string
    descriptor  uint16
    class       *Class
    slot        uint
}

// getters
func (self *Field) Class() (*Class) {
    return self.class
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

func newField(fieldInfo cf.FieldInfo) (*Field) {
    // todo
    return nil
}
