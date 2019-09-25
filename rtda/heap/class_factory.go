package heap

import (
	"sync"

	"github.com/zxh0/jvm.go/classfile"
)

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.initCond = sync.NewCond(&sync.Mutex{})
	class.accessFlags = cf.AccessFlags
	class.copyConstantPool(cf)
	class.copyClassNames(cf)
	class.copyFields(cf)
	class.copyMethods(cf)
	class.copyAttributes(cf)
	return class
}

func (class *Class) copyConstantPool(cf *classfile.ClassFile) {
	class.constantPool = newConstantPool(class, &cf.ConstantPool)
}

func (class *Class) copyClassNames(cf *classfile.ClassFile) {
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
}

func (class *Class) copyFields(cf *classfile.ClassFile) {
	class.fields = make([]*Field, len(cf.Fields))
	for i, fieldInfo := range cf.Fields {
		class.fields[i] = newField(class, fieldInfo)
	}
}

func (class *Class) copyMethods(cf *classfile.ClassFile) {
	class.methods = make([]*Method, len(cf.Methods))
	for i, methodInfo := range cf.Methods {
		class.methods[i] = newMethod(class, methodInfo)
		class.methods[i].slot = uint(i)
	}
}

func (class *Class) copyAttributes(cf *classfile.ClassFile) {
	class.sourceFile = getSourceFile(cf)
	class.signature = getSignature(cf)
	class.annotationData = cf.RuntimeVisibleAnnotationsAttributeData()
	class.enclosingMethod = getEnclosingMethod(cf)
}

func getSourceFile(cf *classfile.ClassFile) string {
	if sfAttr := cf.SourceFileAttribute(); sfAttr != nil {
		return sfAttr.FileName()
	}
	return "Unknown" // todo
}

func getSignature(cf *classfile.ClassFile) string {
	if sigAttr := cf.SignatureAttribute(); sigAttr != nil {
		return sigAttr.Signature()
	}
	return ""
}

func getEnclosingMethod(cf *classfile.ClassFile) *EnclosingMethod {
	if emAttr := cf.EnclosingMethodAttribute(); emAttr != nil {
		methodName, methodDescriptor := emAttr.MethodNameAndDescriptor()
		return &EnclosingMethod{
			className:        emAttr.ClassName(),
			methodName:       methodName,
			methodDescriptor: methodDescriptor,
		}
	}
	return nil
}
