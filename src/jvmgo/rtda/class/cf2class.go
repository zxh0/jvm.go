package class

import (
    //"fmt"
    . "jvmgo/any"
    "jvmgo/classfile"
    //"jvmgo/rtda"
)

func newClass(cf *classfile.ClassFile) (*Class) {
    class := &Class{}
    class.obj = &Obj{} // todo
    class.copyConstantPool(cf)
    class.name = cf.ClassName()
    class.superClassName = cf.SuperClassName()
    class.interfaceNames = cf.InterfaceNames()
    class.copyFields(cf)
    class.copyMethods(cf)
    class.initClassFields()
    return class
}

func (self *Class) copyConstantPool(cf *classfile.ClassFile) {
    self.constantPool = newConstantPool(self, cf.ConstantPool())
}

func (self *Class) copyFields(cf *classfile.ClassFile) {
    self.fields = make([]*Field, len(cf.Fields()))
    for i, fieldInfo := range cf.Fields() {
        self.fields[i] = newField(self, fieldInfo)
    }
}

func (self *Class) copyMethods(cf *classfile.ClassFile) {
    self.methods = make([]*Method, len(cf.Methods()))
    for i, methodInfo := range cf.Methods() {
        self.methods[i] = newMethod(self, methodInfo)
    }
}

func (self *Class) initClassFields() {
    fields := make([]Any, len(self.fields))
    self.obj.fields = fields
    for i, f := range self.fields {
        f.slot = uint(i)
    }
}
