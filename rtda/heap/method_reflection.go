package heap

func (method *Method) ParameterTypes() []*Class {
	if method.ArgSlotCount == 0 {
		return nil
	}

	paramClasses := make([]*Class, 0, method.ArgSlotCount)
	for _, paramType := range method.ParsedDescriptor.ParameterTypes {
		paramClassName := getClassName(string(paramType))
		paramClasses = append(paramClasses, bootLoader.LoadClass(paramClassName))
	}

	return paramClasses
}

func (method *Method) ReturnType() *Class {
	returnDescriptor := method.ParsedDescriptor.ReturnType
	returnClassName := getClassName(string(returnDescriptor))
	returnClass := bootLoader.LoadClass(returnClassName)
	return returnClass
}

func (method *Method) ExceptionTypes() []*Class {
	if method.exIndexTable == nil {
		return nil
	}

	exClasses := make([]*Class, len(method.exIndexTable))
	cp := method.Class.ConstantPool

	for i, exIndex := range method.exIndexTable {
		kClass := cp.GetConstant(uint(exIndex)).(*ConstantClass)
		exClasses[i] = kClass.Class()
	}

	return exClasses
}
