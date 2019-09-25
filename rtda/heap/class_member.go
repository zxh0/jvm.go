package heap

type ClassMember struct {
	AccessFlags
	name           string
	descriptor     string
	signature      string
	annotationData []byte // RuntimeVisibleAnnotations_attribute
	class          *Class
}

func (member *ClassMember) Name() string {
	return member.name
}
func (member *ClassMember) Descriptor() string {
	return member.descriptor
}
func (member *ClassMember) Signature() string {
	return member.signature
}
func (member *ClassMember) AnnotationData() []byte {
	return member.annotationData
}
func (member *ClassMember) Class() *Class {
	return member.class
}

// func (member *ClassMember) ClassLoader() *ClassLoader {
// 	return member.class.classLoader
// }
func (member *ClassMember) ConstantPool() *ConstantPool {
	return member.class.constantPool
}
