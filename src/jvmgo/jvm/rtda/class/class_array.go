package class

func (self *Class) ComponentClass() *Class {
	componentDescriptor := getComponentDescriptor(self.name)
	componentClassName := getClassName(componentDescriptor)
	return self.classLoader.LoadClass(componentClassName)
}
