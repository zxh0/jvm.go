package class

type Attributes struct {
	sourceFile      string
	annotationData  []int8 // RuntimeVisibleAnnotations_attribute
	enclosingMethod *EnclosingMethod
}

func (self *Attributes) SourceFile() string {
	return self.sourceFile
}
func (self *Attributes) AnnotationData() []int8 {
	return self.annotationData
}
func (self *Attributes) EnclosingMethodInfo() *EnclosingMethod {
	return self.enclosingMethod
}

type EnclosingMethod struct {
	className        string
	methodName       string
	methodDescriptor string
}
