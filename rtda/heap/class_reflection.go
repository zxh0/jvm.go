package heap

func (class *Class) IsPrimitive() bool {
	return isPrimitiveType(class.Name)
}

func (class *Class) GetFields(publicOnly bool) []*Field {
	if publicOnly {
		publicFields := make([]*Field, 0, len(class.Fields))
		for _, field := range class.Fields {
			if field.IsPublic() {
				n := len(publicFields)
				publicFields = publicFields[:n+1]
				publicFields[n] = field
			}
		}
		return publicFields
	} else {
		return class.Fields
	}
}

func (class *Class) GetMethods(publicOnly bool) []*Method {
	result := make([]*Method, 0, len(class.Methods))
	for _, method := range class.Methods {
		if !method.IsClinit() && !method.IsConstructor() {
			if !publicOnly || method.IsPublic() {
				n := len(result)
				result = result[:n+1]
				result[n] = method
			}
		}
	}
	return result
}

func (class *Class) GetConstructors(publicOnly bool) []*Method {
	constructors := make([]*Method, 0, len(class.Methods))
	for _, method := range class.Methods {
		if method.IsConstructor() {
			if !publicOnly || method.IsPublic() {
				n := len(constructors)
				constructors = constructors[:n+1]
				constructors[n] = method
			}
		}
	}
	return constructors
}

func (class *Class) GetConstructor(descriptor string) *Method {
	return class.getDeclaredMethod(constructorName, descriptor, false)
}

func (class *Class) GetDefaultConstructor() *Method {
	return class.GetConstructor("()V")
}
