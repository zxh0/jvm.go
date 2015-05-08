package class

import (
	cf "github.com/zxh0/jvm.go/jvmgo/classfile"
)

type ConstantMethodref struct {
	isIface      bool
	className    string
	name         string
	descriptor   string
	argSlotCount uint
	cp           *ConstantPool
	method       *Method
	vslot        int
}

func newConstantMethodref(cp *ConstantPool, methodrefInfo *cf.ConstantMethodrefInfo) *ConstantMethodref {
	return &ConstantMethodref{
		isIface:      false,
		className:    methodrefInfo.ClassName(),
		name:         methodrefInfo.Name(),
		descriptor:   methodrefInfo.Descriptor(),
		argSlotCount: calcArgSlotCount(methodrefInfo.Descriptor()),
		cp:           cp,
		vslot:        -1,
	}
}

// todo
func newConstantInterfaceMethodref(cp *ConstantPool, methodrefInfo *cf.ConstantInterfaceMethodrefInfo) *ConstantMethodref {
	return &ConstantMethodref{
		isIface:      true,
		className:    methodrefInfo.ClassName(),
		name:         methodrefInfo.Name(),
		descriptor:   methodrefInfo.Descriptor(),
		argSlotCount: calcArgSlotCount(methodrefInfo.Descriptor()),
		cp:           cp,
		vslot:        -1,
	}
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

func (self *ConstantMethodref) GetVirtualMethod(ref *Obj) *Method {
	if self.vslot < 0 {
		self.vslot = getVslot(ref.class, self.name, self.descriptor)
	}
	return ref.class.vtable[self.vslot]
}

// todo
func (self *ConstantMethodref) FindInterfaceMethod(ref *Obj) *Method {
	for class := ref.class; class != nil; class = class.superClass {
		method := class.getMethod(self.name, self.descriptor, false)
		if method != nil {
			return method
		}
	}

	if method := findInterfaceMethod(ref.class.interfaces, self.name, self.descriptor); method != nil {
		return method
	} else {
		//TODO
		panic("virtual method not found!")
	}
}

func findInterfaceMethod(interfaces []*Class, name, descriptor string) *Method {
	for i := 0; i < len(interfaces); i++ {
		if method := findInterfaceMethod(interfaces[i].interfaces, name, descriptor); method != nil {
			return method
		}
		method := interfaces[i].getMethod(name, descriptor, false)
		if method != nil {
			return method
		}
	}
	return nil
}
