package heap

func (method *Method) GetParameterTypes() []*Class {
	if method.ArgSlotCount == 0 {
		return nil
	}

	paramClasses := make([]*Class, 0, method.ArgSlotCount)
	for _, paramType := range method.ParameterTypes {
		paramClassName := getClassName(string(paramType))
		paramClasses = append(paramClasses, bootLoader.LoadClass(paramClassName))
	}

	return paramClasses
}

func (method *Method) GetReturnType() *Class {
	returnDescriptor := method.ReturnType
	returnClassName := getClassName(string(returnDescriptor))
	returnClass := bootLoader.LoadClass(returnClassName)
	return returnClass
}

func (method *Method) GetExceptionTypes() []*Class {
	if method.exIndexTable == nil {
		return nil
	}

	exClasses := make([]*Class, len(method.exIndexTable))
	cp := method.Class.ConstantPool

	for i, exIndex := range method.exIndexTable {
		kClass := cp.GetConstantClass(uint(exIndex))
		exClasses[i] = kClass.GetClass()
	}

	return exClasses
}
