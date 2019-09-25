package heap

type ClassAttributes struct {
	sourceFile      string
	signature       string
	annotationData  []byte // RuntimeVisibleAnnotations_attribute
	enclosingMethod *EnclosingMethod
}

func (attrs *ClassAttributes) SourceFile() string {
	return attrs.sourceFile
}
func (attrs *ClassAttributes) Signature() string {
	return attrs.signature
}
func (attrs *ClassAttributes) AnnotationData() []byte {
	return attrs.annotationData
}
func (attrs *ClassAttributes) EnclosingMethod() *EnclosingMethod {
	return attrs.enclosingMethod
}

type EnclosingMethod struct {
	className        string
	methodName       string
	methodDescriptor string
}

func (em *EnclosingMethod) ClassName() string {
	return em.className
}
func (em *EnclosingMethod) MethodName() string {
	return em.methodName
}
func (em *EnclosingMethod) MethodDescriptor() string {
	return em.methodDescriptor
}
