package heap

func (method *Method) ParameterTypes() []*Class {
	if method.argSlotCount == 0 {
		return nil
	}

	paramClasses := make([]*Class, 0, method.argSlotCount)
	for _, paramType := range method.md.parameterTypes {
		paramClassName := getClassName(paramType.descriptor)
		paramClasses = append(paramClasses, bootLoader.LoadClass(paramClassName))
	}

	return paramClasses
}

func (method *Method) ReturnType() *Class {
	returnDescriptor := method.md.returnType.descriptor
	returnClassName := getClassName(returnDescriptor)
	returnClass := bootLoader.LoadClass(returnClassName)
	return returnClass
}

func (method *Method) ExceptionTypes() []*Class {
	if method.exceptions == nil {
		return nil
	}

	exIndexTable := method.exceptions.ExceptionIndexTable
	exClasses := make([]*Class, len(exIndexTable))
	cp := method.class.constantPool

	for i, exIndex := range exIndexTable {
		kClass := cp.GetConstant(uint(exIndex)).(*ConstantClass)
		exClasses[i] = kClass.Class()
	}

	return exClasses
}
