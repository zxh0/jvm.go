package class

import (
    "fmt"
    cf "jvmgo/classfile"
    "jvmgo/util"
)

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

func (self *ConstantFieldref) String() string {
    return fmt.Sprintf("{ConstantFieldref className:%v name:%v descriptor:%v}",
            self.className, self.name, self.descriptor)
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
    util.Panicf("field not found! %v", self)
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
        return
    }

    // todo
    util.Panicf("static field not found! %v", self)
}
