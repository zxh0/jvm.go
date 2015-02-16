package class

var _emptyParameterTypes = []*Class{}

func (self *Method) ParameterTypes() []*Class {
	if self.argCount == 0 {
		return _emptyParameterTypes
	}

	classLoader := self.class.classLoader
	paramClasses := make([]*Class, self.argCount)
	for i, paramType := range self.MethodDescriptor().parameterTypes {
		paramClassName := getClassName(paramType.descriptor)
		paramClasses[i] = classLoader.LoadClass(paramClassName)
	}

	return paramClasses
}

func (self *Method) ReturnType() *Class {
	returnDescriptor := self.MethodDescriptor().returnType.descriptor
	returnClassName := getClassName(returnDescriptor)
	returnClass := self.class.classLoader.LoadClass(returnClassName)
	return returnClass
}
