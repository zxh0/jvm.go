package class

import cf "jvmgo/classfile"

type ConstantMethodref struct {
    className   string
    name        string
    descriptor  string
    method      *Method
}

type ConstantInterfaceMethodref struct {
    ConstantMethodref
}

func (self *ConstantMethodref) Method() (*Method) {
    if self.method == nil {
        self.resolve()
    }
    return self.method
}

func (self *ConstantMethodref) resolve() {
    // todo
}

func (self *ConstantMethodref) init(methodrefInfo *cf.ConstantMethodrefInfo) {
    self.className = methodrefInfo.ClassName()
    self.name = methodrefInfo.Name()
    self.descriptor = methodrefInfo.Descriptor()
}

func newConstantMethodref(methodrefInfo *cf.ConstantMethodrefInfo) (*ConstantMethodref) {
    methodref := &ConstantMethodref{}
    methodref.init(methodrefInfo)
    return methodref
}
