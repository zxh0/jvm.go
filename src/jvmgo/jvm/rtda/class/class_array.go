package class

func (self *Class) IsArray() bool {
	return self.name[0] == '['
}

func (self *Class) ComponentClass() *Class {
	componentClassName := getComponentClassName(self.name)
	return self.classLoader.LoadClass(componentClassName)
}

func (self *Class) ArrayClass() *Class {
	if self.IsArray() {
		arrayClassName := "[" + self.name
		return self.classLoader.LoadClass(arrayClassName)
	} else if self.IsPrimitive() {
		return nil // todo
	} else {
		arrayClassName := "[L" + self.name + ";"
		return self.classLoader.LoadClass(arrayClassName)
	}
}
