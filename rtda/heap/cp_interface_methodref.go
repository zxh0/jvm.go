package heap

type ConstantInterfaceMethodref struct {
	ConstantMethodref
}

// todo
func (self *ConstantInterfaceMethodref) FindInterfaceMethod(ref *Object) *Method {
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
