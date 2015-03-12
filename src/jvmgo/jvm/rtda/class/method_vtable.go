package class

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

	// class._eachMethod(func(m *Method) {

	// })
	for _, m := range class.methods {
		if !m.IsStatic() {
			if i := search(class.vtable, m); i > -1 {
				class.vtable[i] = m // override
			} else {
				addVmethod(class, m)
			}
		}
	}
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

func addVmethod(class *Class, m *Method) {
	_len := len(class.vtable)
	if _len == cap(class.vtable) {
		newVtable := make([]*Method, _len, _len+8)
		copy(newVtable, class.vtable)
		class.vtable = newVtable
	}

	class.vtable = append(class.vtable, m)
}

// func countNewVirtualMethod(class *Class) int {
// 	superVtable := getSuperVtable(class)

// 	count := 0
// 	class._eachMethod(func(m *Method) {
// 		if isVirtualMethod(m) && search(superVtable, m) < 0 {
// 			count++
// 		}
// 	})

// 	return count
// }

func isVirtualMethod(method *Method) bool {
	return !method.IsStatic() &&
		!method.IsFinal() &&
		!method.IsPrivate() &&
		method.Name() != constructorName
}

func search(vtable []*Method, m *Method) int {
	for i, vm := range vtable {
		if vm.name == m.name && vm.descriptor == m.descriptor {
			return i
		}
	}
	return -1
}

// visit all class and interface methods
func (self *Class) _eachMethod(f func(*Method)) {
	for _, m := range self.methods {
		f(m)
	}
	for _, iface := range self.interfaces {
		iface._eachMethod(f)
	}
}
