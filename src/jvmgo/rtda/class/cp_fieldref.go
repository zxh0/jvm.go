package class

import cf "jvmgo/classfile"

type ConstantFieldref struct {
    className   string
    name        string
    descriptor  string
    cp          *ConstantPool
    field       *Field
}

func newConstantFieldref(cp *ConstantPool, fieldrefInfo *cf.ConstantFieldrefInfo) (*ConstantFieldref) {
    fieldref := &ConstantFieldref{}
    fieldref.className = fieldrefInfo.ClassName()
    fieldref.name = fieldrefInfo.Name()
    fieldref.descriptor = fieldrefInfo.Descriptor()
    fieldref.cp = cp
    return fieldref
}

func (self *ConstantFieldref) InstanceField() (*Field) {
    if self.field == nil {
        self.resolveInstanceField()
    }
    return self.field
}
func (self *ConstantFieldref) resolveInstanceField() {
    classLoader := self.cp.class.classLoader
    fromClass := classLoader.getClass(self.className)

    for class := fromClass; class != nil; class = class.superClass {
        field := class.GetField(self.name, self.descriptor)
        if field != nil && !field.IsStatic() {
            self.field = field
            return
        }
    }

    // todo
    panic("field not found!")
}

func (self *ConstantFieldref) StaticField() (*Field) {
    if self.field == nil {
        self.resolveStaticField()
    }
    return self.field
}

func (self *ConstantFieldref) resolveStaticField() {
    classLoader := self.cp.class.classLoader
    class := classLoader.LoadClass(self.className)
    field := class.GetField(self.name, self.descriptor)
    if field != nil && field.IsStatic() {
        self.field = field
    } else {
        panic("static field not found!") // todo
    }
}
