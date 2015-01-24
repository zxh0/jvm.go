package class

import (
    . "jvmgo/any"
    cf "jvmgo/classfile"
)

type ConstantMethodref struct {
    className       string
    name            string
    descriptor      string
    cp              *ConstantPool
    method          *Method
    nativeMethod    Any // cannot use package 'native' because of cycle import!
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
func (self *ConstantMethodref) NativeMethod() (Any) {
    return self.nativeMethod
}

func (self *ConstantMethodref) resolve() {
    class := self.cp.class.classLoader.LoadClass(self.className)
    self.method = class.getMethod(self.name, self.descriptor)
    if self.method.IsNative() && !self.method.IsRegisterNatives() {
        self.nativeMethod = findNativeMethod(self.method)
    }
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
