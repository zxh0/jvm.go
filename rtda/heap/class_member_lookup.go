package heap

func (class *Class) GetDeclaredField(name, descriptor string) *Field {
	for _, field := range class.Fields {
		if field.Name == name && field.Descriptor == descriptor {
			return field
		}
	}
	return nil
}
func (class *Class) GetDeclaredMethod(name, descriptor string) *Method {
	for _, method := range class.Methods {
		if method.Name == name && method.Descriptor == descriptor {
			return method
		}
	}
	return nil
}

func (class *Class) GetField(name, descriptor string) *Field {
	for k := class; k != nil; k = k.SuperClass {
		for _, field := range k.Fields {
			if field.Name == name &&
				(descriptor == "*" || field.Descriptor == descriptor) {

				return field
			}
		}
	}
	return nil
}
func (class *Class) GetMethod(name, descriptor string) *Method {
	for k := class; k != nil; k = k.SuperClass {
		for _, method := range k.Methods {
			if method.Name == name && method.Descriptor == descriptor {
				return method
			}
		}
	}
	return nil
}

func (class *Class) GetStaticField(name, descriptor string) *Field {
	field := class.GetField(name, descriptor)
	if field != nil && !field.IsStatic() {
		panic("not static")
	}
	return field
}
func (class *Class) GetInstanceField(name, descriptor string) *Field {
	field := class.GetField(name, descriptor)
	if field != nil && field.IsStatic() {
		panic("static")
	}
	return field
}

func (class *Class) GetStaticMethod(name, descriptor string) *Method {
	method := class.GetMethod(name, descriptor)
	if method != nil && !method.IsStatic() {
		println(class.Name, name, descriptor)
		panic("not static:")
	}
	return method
}
func (class *Class) GetInstanceMethod(name, descriptor string) *Method {
	method := class.GetMethod(name, descriptor)
	if method != nil && method.IsStatic() {
		panic("static")
	}
	return method
}

func (class *Class) GetMainMethod() *Method {
	return class.GetStaticMethod(mainMethodName, mainMethodDesc)
}
func (class *Class) GetClinitMethod() *Method {
	return class.GetDeclaredMethod(clinitMethodName, clinitMethodDesc)
}

// reflection
func (class *Class) GetStaticValue(fieldName, fieldDescriptor string) Slot {
	field := class.GetStaticField(fieldName, fieldDescriptor)
	return field.GetStaticValue()
}
func (class *Class) SetStaticValue(fieldName, fieldDescriptor string, value Slot) {
	field := class.GetStaticField(fieldName, fieldDescriptor)
	field.PutStaticValue(value)
}
