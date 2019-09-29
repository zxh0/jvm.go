package heap

type ConstantInterfaceMethodRef struct {
	ConstantMethodRef
}

// todo
func (imr *ConstantInterfaceMethodRef) FindInterfaceMethod(ref *Object) *Method {
	for class := ref.class; class != nil; class = class.SuperClass {
		method := class.getMethod(imr.name, imr.descriptor, false)
		if method != nil {
			return method
		}
	}

	if method := findInterfaceMethod(ref.class.Interfaces, imr.name, imr.descriptor); method != nil {
		return method
	} else {
		//TODO
		panic("virtual method not found!")
	}
}

func findInterfaceMethod(interfaces []*Class, name, descriptor string) *Method {
	for i := 0; i < len(interfaces); i++ {
		if method := findInterfaceMethod(interfaces[i].Interfaces, name, descriptor); method != nil {
			return method
		}
		method := interfaces[i].getMethod(name, descriptor, false)
		if method != nil {
			return method
		}
	}
	return nil
}
