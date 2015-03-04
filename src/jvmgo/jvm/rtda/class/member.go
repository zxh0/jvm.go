package class

type ClassMember struct {
	AccessFlags
	name           string
	descriptor     string
	signature      string
	annotationData []int8 // RuntimeVisibleAnnotations_attribute
	class          *Class
}

func (self *ClassMember) Name() string {
	return self.name
}
func (self *ClassMember) Descriptor() string {
	return self.descriptor
}
func (self *ClassMember) Signature() string {
	return self.signature
}
func (self *ClassMember) AnnotationData() []int8 {
	return self.annotationData
}
func (self *ClassMember) Class() *Class {
	return self.class
}

func (self *ClassMember) ClassLoader() *ClassLoader {
	return self.class.classLoader
}
func (self *ClassMember) ConstantPool() *ConstantPool {
	return self.class.constantPool
}
