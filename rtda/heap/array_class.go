package heap

func (class *Class) IsArray() bool {
	return class.name[0] == '['
}

func (class *Class) IsPrimitiveArray() bool {
	return class.IsArray() && len(class.name) == 2
}

func (class *Class) ComponentClass() *Class {
	componentClassName := getComponentClassName(class.name)
	return bootLoader.LoadClass(componentClassName)
}

func (class *Class) arrayClass() *Class {
	arrayClassName := getArrayClassName(class.name)
	return bootLoader.LoadClass(arrayClassName)
}
