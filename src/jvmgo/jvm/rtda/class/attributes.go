package class

type Attributes struct {
	annotationData []int8 // RuntimeVisibleAnnotations_attribute
	enclosingMethod *EnclosingMethod
}

func (self *Attributes) AnnotationData() []int8 {
	return self.annotationData
}
func (self *Attributes) EnclosingMethodInfo() (string, string, string) {
	return self.enclosingMethod.className,
		self.enclosingMethod.methodName,
		self.enclosingMethod.methodDescriptor
}

type EnclosingMethod struct {
	className string
	methodName string
	methodDescriptor string
}
