package heap

import (
	"sync"

	"github.com/zxh0/jvm.go/classfile"
)

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{
		AccessFlags:     classfile.AccessFlags(cf.AccessFlags),
		Name:            cf.GetThisClassName(),
		superClassName:  cf.GetSuperClassName(),
		interfaceNames:  cf.GetInterfaceNames(),
		SourceFile:      cf.GetUTF8(cf.GetSourceFileIndex()),
		Signature:       cf.GetUTF8(cf.GetSignatureIndex()),
		AnnotationData:  cf.GetRuntimeVisibleAnnotationsAttributeData(),
		EnclosingMethod: getEnclosingMethod(cf),
		InitCond:        sync.NewCond(&sync.Mutex{}),
	}
	class.ConstantPool = newConstantPool(class, cf)
	class.copyFields(cf)
	class.copyMethods(cf)
	return class
}

func (class *Class) copyFields(cf *classfile.ClassFile) {
	class.Fields = make([]*Field, len(cf.Fields))
	for i, fieldInfo := range cf.Fields {
		class.Fields[i] = newField(class, cf, fieldInfo)
	}
}

func (class *Class) copyMethods(cf *classfile.ClassFile) {
	class.Methods = make([]*Method, len(cf.Methods))
	for i, methodInfo := range cf.Methods {
		class.Methods[i] = newMethod(class, cf, methodInfo, uint(i))
	}
}

func getEnclosingMethod(cf *classfile.ClassFile) *EnclosingMethod {
	if emAttr, found := cf.GetEnclosingMethodAttribute(); found {
		methodName, methodDescriptor := cf.GetNameAndType(emAttr.MethodIndex)
		return &EnclosingMethod{
			ClassName:        cf.GetClassName(emAttr.ClassIndex),
			MethodName:       methodName,
			MethodDescriptor: methodDescriptor,
		}
	}
	return nil
}
