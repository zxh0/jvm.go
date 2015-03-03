package class

func (self *Method) ParameterTypes() []*Class {
	if self.argCount == 0 {
		return nil
	}

	classLoader := self.class.classLoader
	paramClasses := make([]*Class, self.argCount)
	for i, paramType := range self.md.parameterTypes {
		paramClassName := getClassName(paramType.descriptor)
		paramClasses[i] = classLoader.LoadClass(paramClassName)
	}

	return paramClasses
}

func (self *Method) ReturnType() *Class {
	returnDescriptor := self.md.returnType.descriptor
	returnClassName := getClassName(returnDescriptor)
	returnClass := self.class.classLoader.LoadClass(returnClassName)
	return returnClass
}

func (self *Method) ExceptionTypes() []*Class {
	if self.exceptions == nil {
		return nil
	}

	exIndexTable := self.exceptions.ExceptionIndexTable()
	exClasses := make([]*Class, len(exIndexTable))
	cp := self.class.constantPool

	for i, exIndex := range exIndexTable {
		kClass := cp.GetConstant(uint(exIndex)).(*ConstantClass)
		exClasses[i] = kClass.Class()
	}

	return exClasses
}
