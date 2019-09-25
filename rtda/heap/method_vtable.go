package heap

func getVslot(class *Class, name, descriptor string) int {
	for i, m := range class.vtable {
		if m.name == name && m.descriptor == descriptor {
			return i
		}
	}
	// todo
	return -1
}

func createVtable(class *Class) {
	class.vtable = copySuperVtable(class)

	for _, m := range class.methods {
		if isVirtualMethod(m) {
			if i := indexOf(class.vtable, m); i > -1 {
				class.vtable[i] = m // override
			} else {
				addVmethod(class, m)
			}
		}
	}

	_eachInterfaceMethod(class, func(m *Method) {
		if i := indexOf(class.vtable, m); i < 0 {
			addVmethod(class, m)
		}
	})
}

func copySuperVtable(class *Class) []*Method {
	if class.superClass != nil {
		superVtable := class.superClass.vtable
		newVtable := make([]*Method, len(superVtable))
		copy(newVtable, superVtable)
		return newVtable
	} else {
		return nil
	}
}

func isVirtualMethod(method *Method) bool {
	return !method.IsStatic() &&
		//!method.IsFinal() &&
		!method.IsPrivate() &&
		method.Name() != constructorName
}

func indexOf(vtable []*Method, m *Method) int {
	for i, vm := range vtable {
		if vm.name == m.name && vm.descriptor == m.descriptor {
			return i
		}
	}
	return -1
}

func addVmethod(class *Class, m *Method) {
	_len := len(class.vtable)
	if _len == cap(class.vtable) {
		newVtable := make([]*Method, _len, _len+8)
		copy(newVtable, class.vtable)
		class.vtable = newVtable
	}

	class.vtable = append(class.vtable, m)
}

// visit all interface methods
func _eachInterfaceMethod(class *Class, f func(*Method)) {
	for _, iface := range class.interfaces {
		_eachInterfaceMethod(iface, f)
		for _, m := range iface.methods {
			f(m)
		}
	}
}
