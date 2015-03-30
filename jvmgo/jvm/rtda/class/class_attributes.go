package class

type Attributes struct {
	sourceFile      string
	annotationData  []int8 // RuntimeVisibleAnnotations_attribute
	enclosingMethod *EnclosingMethod
	signature       string
}

func (self *Attributes) SourceFile() string {
	return self.sourceFile
}

func (self *Attributes) AnnotationData() []int8 {
	return self.annotationData
}

func (self *Attributes) EnclosingMethod() *EnclosingMethod {
	return self.enclosingMethod
}

func (self *Attributes) Signature() string {
	return self.signature
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
