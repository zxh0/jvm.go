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

func newConstantMethodref(methodrefInfo *cf.ConstantMethodrefInfo) (*ConstantMethodref) {
    methodref := &ConstantMethodref{}
    methodref.className = methodrefInfo.ClassName()
    methodref.name = methodrefInfo.Name()
    methodref.descriptor = methodrefInfo.Descriptor()
    return methodref
}

func newConstantInterfaceMethodref(methodrefInfo *cf.ConstantInterfaceMethodrefInfo) (*ConstantInterfaceMethodref) {
    methodref := &ConstantInterfaceMethodref{}
    methodref.className = methodrefInfo.ClassName()
    methodref.name = methodrefInfo.Name()
    methodref.descriptor = methodrefInfo.Descriptor()
    return methodref
}
