package heap

func (class *Class) IsAssignableFrom(cls *Class) bool {
	return class == cls ||
		class.isSuperClassOf(cls) ||
		class.isSuperInterfaceOf(cls)
}

// class implements iface
func (class *Class) isImplements(iface *Class) bool {
	for k := class; k != nil; k = k.superClass {
		for _, i := range k.interfaces {
			if i == iface || i.isSubInterfaceOf(iface) {
				return true
			}
		}
	}
	return false
}

// iface extends class
func (class *Class) isSuperInterfaceOf(iface *Class) bool {
	return iface.isSubInterfaceOf(class)
}

// class extends iface
func (class *Class) isSubInterfaceOf(iface *Class) bool {
	for _, superInterface := range class.interfaces {
		if superInterface == iface || superInterface.isSubInterfaceOf(iface) {
			return true
		}
	}
	return false
}

// c extends class
func (class *Class) isSuperClassOf(c *Class) bool {
	return c.isSubClassOf(class)
}

// class extends c
func (class *Class) isSubClassOf(c *Class) bool {
	for k := class.superClass; k != nil; k = k.superClass {
		if k == c {
			return true
		}
	}
	return false
}
