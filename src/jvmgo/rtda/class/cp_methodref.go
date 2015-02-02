package class

import (
    //"fmt"
    cf "jvmgo/classfile"
)

type ConstantMethodref struct {
    className       string
    name            string
    descriptor      string
    argCount        uint
    cp              *ConstantPool
    method          *Method
}

func newConstantMethodref(cp *ConstantPool, methodrefInfo *cf.ConstantMethodrefInfo) (*ConstantMethodref) {
    methodref := &ConstantMethodref{}
    methodref.className = methodrefInfo.ClassName()
    methodref.name = methodrefInfo.Name()
    methodref.descriptor = methodrefInfo.Descriptor()
    methodref.argCount = methodrefInfo.ArgCount()
    methodref.cp = cp
    return methodref
}

func (self *ConstantMethodref) ArgCount() (uint) {
    return self.argCount       
}

func (self *ConstantMethodref) StaticMethod() (*Method) {
    if self.method == nil {
        self.resolveStaticMethod()
    }
    return self.method
}
func (self *ConstantMethodref) resolveStaticMethod() {
    method := self.findMethod(self.className)
    if method != nil && method.IsStatic() {
        if method.IsNative() {
            method.nativeMethod = findNativeMethod(method)
        }
        self.method = method
    } else {
        // todo
        panic("static method not found!")
    }
}

func (self *ConstantMethodref) SpecialMethod() (*Method) {
    if self.method == nil {
        self.resolveSpecialMethod()
    }
    return self.method
}
func (self *ConstantMethodref) resolveSpecialMethod() {
    method := self.findMethod(self.className)
    if method != nil && !method.IsStatic() {
        if method.IsNative() {
            method.nativeMethod = findNativeMethod(method)
        }
        self.method = method
    } else {
        // todo
        panic("special method not found!")
    }
}

func (self *ConstantMethodref) findMethod(className string) (*Method) {
    class := self.cp.class.classLoader.LoadClass(className)
    return class.GetMethod(self.name, self.descriptor)
}

// todo
func (self *ConstantMethodref) VirtualMethod(ref *Obj) (*Method) {
    className := ref.class.name // todo
    return self.findVirtualMethod(className)
}
func (self *ConstantMethodref) findVirtualMethod(className string) (*Method) {
    classLoader := self.cp.class.classLoader
    for {
        if className != "" {
            class := classLoader.LoadClass(className)
            method := class.GetMethod(self.name, self.descriptor)
            if method != nil && !method.IsStatic() {
                if method.IsNative() && method.nativeMethod == nil {
                    method.nativeMethod = findNativeMethod(method)
                }
                return method
            } else {
                className = class.superClassName
            }
        } else {
            break
        }
    }
    // todo
    panic("virtual method not found!")
}


type ConstantInterfaceMethodref struct {
    ConstantMethodref
}

func newConstantInterfaceMethodref(cp *ConstantPool, methodrefInfo *cf.ConstantInterfaceMethodrefInfo) (*ConstantInterfaceMethodref) {
    methodref := &ConstantInterfaceMethodref{}
    methodref.cp = cp
    methodref.className = methodrefInfo.ClassName()
    methodref.name = methodrefInfo.Name()
    methodref.descriptor = methodrefInfo.Descriptor()
    methodref.argCount = methodrefInfo.ArgCount()
    return methodref
}
