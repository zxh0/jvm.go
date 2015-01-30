package class

import cf "jvmgo/classfile"

type ClassMember struct {
    cf.AccessFlags
    name        string
    descriptor  string
    class       *Class
}

func (self *ClassMember) Name() (string) {
    return self.name
}
func (self *ClassMember) Class() (*Class) {
    return self.class
}
