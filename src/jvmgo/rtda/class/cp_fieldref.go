package class

import cf "jvmgo/classfile"

type ConstantFieldref struct {
    className   string
    name        string
    descriptor  string
    cp          *ConstantPool
    field       *Field
}

func (self *ConstantFieldref) InstanceField() (*Field) {
    if self.field == nil {
        self.resolveInstanceField()
    }
    return self.field
}
func (self *ConstantFieldref) resolveInstanceField() {
    classLoader := self.cp.class.classLoader
    className := self.className
    for {
        if className != "" {
            class := classLoader.getClass(className)
            field := class.GetField(self.name, self.descriptor)
            if field != nil && !field.IsStatic() {
                self.field = field
                return
            } else {
                className = class.superClassName
            }
        } else {
            break
        }
    }
    panic("field not found!") // todo
}

func (self *ConstantFieldref) Field() (*Field) {
    if self.field == nil {
        self.resolve()
    }
    return self.field
}

func (self *ConstantFieldref) resolve() {
    class := self.cp.class.classLoader.LoadClass(self.className)
    self.field = class.GetField(self.name, self.descriptor)
}

func newConstantFieldref(cp *ConstantPool, fieldrefInfo *cf.ConstantFieldrefInfo) (*ConstantFieldref) {
    fieldref := &ConstantFieldref{}
    fieldref.cp = cp
    fieldref.className = fieldrefInfo.ClassName()
    fieldref.name = fieldrefInfo.Name()
    fieldref.descriptor = fieldrefInfo.Descriptor()
    return fieldref
}
