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
	if method.exIndexTable == nil {
		return nil
	}

	exClasses := make([]*Class, len(method.exIndexTable))
	cp := method.class.constantPool

	for i, exIndex := range method.exIndexTable {
		kClass := cp.GetConstant(uint(exIndex)).(*ConstantClass)
		exClasses[i] = kClass.Class()
	}

	return exClasses
}
