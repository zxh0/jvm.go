package heap

type ClassMember struct {
	AccessFlags
	name           string
	descriptor     string
	signature      string
	annotationData []byte // RuntimeVisibleAnnotations_attribute
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
func (self *ClassMember) AnnotationData() []byte {
	return self.annotationData
}
func (self *ClassMember) Class() *Class {
	return self.class
}

// func (self *ClassMember) ClassLoader() *ClassLoader {
// 	return self.class.classLoader
// }
func (self *ClassMember) ConstantPool() *ConstantPool {
	return self.class.constantPool
}
