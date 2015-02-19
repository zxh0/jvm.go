package class

func createVtable(class *Class) {
	superVtable := getSuperVtable(class)
	newVirtualMethodCount := countNewVirtualMethod(class)

	newCap := len(superVtable) + newVirtualMethodCount
	newVtable := make([]*Method, len(superVtable), newCap)
	copy(newVtable, superVtable)

	for _, m := range class.methods {
		if i := search(superVtable, m); i > -1 {
			newVtable[i] = m // override
		} else {
			newVtable = append(newVtable, m)
		}
	}
}

func getSuperVtable(class *Class) []*Method {
	if class.superClass != nil {
		return class.superClass.vtable
	} else {
		return nil
	}
}

func countNewVirtualMethod(class *Class) int {
	superVtable := getSuperVtable(class)

	count := 0
	for _, m := range class.methods {
		if isVirtualMethod(m) && search(superVtable, m) < 0 {
			count++
		}
	}
	return count
}

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
