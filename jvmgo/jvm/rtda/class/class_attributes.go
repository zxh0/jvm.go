package class

type Attributes struct {
	sourceFile      string
	signature       string
	annotationData  []int8 // RuntimeVisibleAnnotations_attribute
	enclosingMethod *EnclosingMethod
}

func (self *Attributes) SourceFile() string {
	return self.sourceFile
}
func (self *Attributes) Signature() string {
	return self.signature
}
func (self *Attributes) AnnotationData() []int8 {
	return self.annotationData
}
func (self *Attributes) EnclosingMethod() *EnclosingMethod {
	return self.enclosingMethod
}

type EnclosingMethod struct {
	className        string
	methodName       string
	methodDescriptor string
}

func (self *EnclosingMethod) ClassName() string {
	return self.className
}
func (self *EnclosingMethod) MethodName() string {
	return self.methodName
}
func (self *EnclosingMethod) MethodDescriptor() string {
	return self.methodDescriptor
}
