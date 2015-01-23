package class

import cf "jvmgo/classfile"

type ConstantFieldref struct {
    className   string
    name        string
    descriptor  string
    cp          *ConstantPool
    field       *Field
}

func (self *ConstantFieldref) Field() (*Field) {
    if self.field == nil {
        self.resolve()
    }
    return self.field
}

func (self *ConstantFieldref) resolve() {
    // todo
}

func newConstantFieldref(cp *ConstantPool, fieldrefInfo *cf.ConstantFieldrefInfo) (*ConstantFieldref) {
    fieldref := &ConstantFieldref{cp:cp}
    fieldref.className = fieldrefInfo.ClassName()
    fieldref.name = fieldrefInfo.Name()
    fieldref.descriptor = fieldrefInfo.Descriptor()
    return fieldref
}
