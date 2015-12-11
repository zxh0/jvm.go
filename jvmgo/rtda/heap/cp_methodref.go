package heap

import (
	cf "github.com/zxh0/jvm.go/jvmgo/classfile"
)

type ConstantMethodref struct {
	ConstantMemberref
	argSlotCount uint
	method       *Method
	vslot        int
}

func newConstantMethodref(refInfo *cf.ConstantMethodrefInfo) *ConstantMethodref {
	ref := &ConstantMethodref{vslot: -1}
	ref.copy(&refInfo.ConstantMemberrefInfo)
	ref.argSlotCount = calcArgSlotCount(ref.descriptor)
	return ref
}

func (self *ConstantMethodref) ArgSlotCount() uint {
	return self.argSlotCount
}

func (self *ConstantMethodref) StaticMethod() *Method {
	if self.method == nil {
		self.resolveStaticMethod()
	}
	return self.method
}
func (self *ConstantMethodref) resolveStaticMethod() {
	method := self.findMethod(true)
	if method != nil {
		self.method = method
	} else {
		// todo
		panic("static method not found!")
	}
}

func (self *ConstantMethodref) SpecialMethod() *Method {
	if self.method == nil {
		self.resolveSpecialMethod()
	}
	return self.method
}
func (self *ConstantMethodref) resolveSpecialMethod() {
	method := self.findMethod(false)
	if method != nil {
		self.method = method
		return
	}

	// todo
	// class := self.cp.class.classLoader.LoadClass(self.className)
	// if class.IsInterface() {
	// 	method = self.findMethodInInterfaces(class)
	// 	if method != nil {
	// 		self.method = method
	// 		return
	// 	}
	// }

	// todo
	panic("special method not found!")
}

func (self *ConstantMethodref) findMethod(isStatic bool) *Method {
	class := bootLoader.LoadClass(self.className)
	return class.getMethod(self.name, self.descriptor, isStatic)
}

// todo
/*func (self *ConstantMethodref) findMethodInInterfaces(iface *Class) *Method {
	for _, m := range iface.methods {
		if !m.IsAbstract() {
			if m.name == self.name && m.descriptor == self.descriptor {
				return m
			}
		}
	}

	for _, superIface := range iface.interfaces {
		if m := self.findMethodInInterfaces(superIface); m != nil {
			return m
		}
	}

	return nil
}*/

func (self *ConstantMethodref) GetVirtualMethod(ref *Object) *Method {
	if self.vslot < 0 {
		self.vslot = getVslot(ref.class, self.name, self.descriptor)
	}
	return ref.class.vtable[self.vslot]
}
