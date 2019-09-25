package heap

func (self *Class) IsPrimitive() bool {
	return isPrimitiveType(self.name)
}

func (self *Class) GetFields(publicOnly bool) []*Field {
	if publicOnly {
		publicFields := make([]*Field, 0, len(self.fields))
		for _, field := range self.fields {
			if field.IsPublic() {
				n := len(publicFields)
				publicFields = publicFields[:n+1]
				publicFields[n] = field
			}
		}
		return publicFields
	} else {
		return self.fields
	}
}

func (self *Class) GetMethods(publicOnly bool) []*Method {
	result := make([]*Method, 0, len(self.methods))
	for _, method := range self.methods {
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

func (self *Class) GetConstructors(publicOnly bool) []*Method {
	constructors := make([]*Method, 0, len(self.methods))
	for _, method := range self.methods {
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

func (self *Class) GetConstructor(descriptor string) *Method {
	return self._getMethod(constructorName, descriptor, false)
}

func (self *Class) GetDefaultConstructor() *Method {
	return self.GetConstructor("()V")
}
