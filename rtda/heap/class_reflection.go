package heap

func (class *Class) IsPrimitive() bool {
	return isPrimitiveType(class.name)
}

func (class *Class) GetFields(publicOnly bool) []*Field {
	if publicOnly {
		publicFields := make([]*Field, 0, len(class.fields))
		for _, field := range class.fields {
			if field.IsPublic() {
				n := len(publicFields)
				publicFields = publicFields[:n+1]
				publicFields[n] = field
			}
		}
		return publicFields
	} else {
		return class.fields
	}
}

func (class *Class) GetMethods(publicOnly bool) []*Method {
	result := make([]*Method, 0, len(class.methods))
	for _, method := range class.methods {
		if !method.IsClinit() && !method.isConstructor() {
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
	constructors := make([]*Method, 0, len(class.methods))
	for _, method := range class.methods {
		if method.isConstructor() {
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
	return class._getMethod(constructorName, descriptor, false)
}

func (class *Class) GetDefaultConstructor() *Method {
	return class.GetConstructor("()V")
}
