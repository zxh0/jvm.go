package heap

type ClassAttributes struct {
	sourceFile      string
	signature       string
	annotationData  []byte // RuntimeVisibleAnnotations_attribute
	enclosingMethod *EnclosingMethod
}

func (self *ClassAttributes) SourceFile() string {
	return self.sourceFile
}
func (self *ClassAttributes) Signature() string {
	return self.signature
}
func (self *ClassAttributes) AnnotationData() []byte {
	return self.annotationData
}
func (self *ClassAttributes) EnclosingMethod() *EnclosingMethod {
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
