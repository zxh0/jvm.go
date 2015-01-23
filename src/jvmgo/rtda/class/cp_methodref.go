package class

import cf "jvmgo/classfile"

type ConstantMethodref struct {
    className   string
    name        string
    descriptor  string
    cp          *ConstantPool
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
    self.cp.class.classLoader.getClass(self.className)
    // todo
    panic("cp_methodref.go!!")
}

func newConstantMethodref(cp *ConstantPool, methodrefInfo *cf.ConstantMethodrefInfo) (*ConstantMethodref) {
    methodref := &ConstantMethodref{}
    methodref.className = methodrefInfo.ClassName()
    methodref.name = methodrefInfo.Name()
    methodref.descriptor = methodrefInfo.Descriptor()
    methodref.cp = cp
    return methodref
}

func newConstantInterfaceMethodref(cp *ConstantPool, methodrefInfo *cf.ConstantInterfaceMethodrefInfo) (*ConstantInterfaceMethodref) {
    methodref := &ConstantInterfaceMethodref{}
    methodref.className = methodrefInfo.ClassName()
    methodref.name = methodrefInfo.Name()
    methodref.descriptor = methodrefInfo.Descriptor()
    methodref.cp = cp
    return methodref
}
