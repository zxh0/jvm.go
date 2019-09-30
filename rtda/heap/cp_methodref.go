package heap

type ConstantMethodRef struct {
	ConstantMemberRef
	ArgSlotCount uint
	method       *Method
	vslot        int
}

func (mr *ConstantMethodRef) StaticMethod() *Method {
	if mr.method == nil {
		mr.resolveStaticMethod()
	}
	return mr.method
}
func (mr *ConstantMethodRef) resolveStaticMethod() {
	method := mr.findMethod(true)
	if method != nil {
		mr.method = method
	} else {
		// todo
		panic("static method not found!")
	}
}

func (mr *ConstantMethodRef) SpecialMethod() *Method {
	if mr.method == nil {
		mr.resolveSpecialMethod()
	}
	return mr.method
}
func (mr *ConstantMethodRef) resolveSpecialMethod() {
	method := mr.findMethod(false)
	if method != nil {
		mr.method = method
		return
	}

	// todo
	// class := mr.cp.class.classLoader.LoadClass(mr.className)
	// if class.IsInterface() {
	// 	method = mr.findMethodInInterfaces(class)
	// 	if method != nil {
	// 		mr.method = method
	// 		return
	// 	}
	// }

	// todo
	panic("special method not found!")
}

func (mr *ConstantMethodRef) findMethod(isStatic bool) *Method {
	class := bootLoader.LoadClass(mr.className)
	return class.getMethod(mr.name, mr.descriptor, isStatic)
}

// todo
/*func (mr *ConstantMethodref) findMethodInInterfaces(iface *Class) *Method {
	for _, m := range iface.methods {
		if !m.IsAbstract() {
			if m.name == mr.name && m.descriptor == mr.descriptor {
				return m
			}
		}
	}

	for _, superIface := range iface.interfaces {
		if m := mr.findMethodInInterfaces(superIface); m != nil {
			return m
		}
	}

	return nil
}*/

func (mr *ConstantMethodRef) GetVirtualMethod(ref *Object) *Method {
	if mr.vslot < 0 {
		mr.vslot = getVslot(ref.Class, mr.name, mr.descriptor)
	}
	return ref.Class.vtable[mr.vslot]
}
