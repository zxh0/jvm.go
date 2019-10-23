package heap

import (
	"sync"

	"github.com/zxh0/jvm.go/classfile"
)

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.InitCond = sync.NewCond(&sync.Mutex{})
	class.AccessFlags = classfile.AccessFlags(cf.AccessFlags)
	class.copyConstantPool(cf)
	class.copyClassNames(cf)
	class.copyFields(cf)
	class.copyMethods(cf)
	class.copyAttributes(cf)
	return class
}

func (class *Class) copyConstantPool(cf *classfile.ClassFile) {
	class.ConstantPool = newConstantPool(cf)
}

func (class *Class) copyClassNames(cf *classfile.ClassFile) {
	class.Name = cf.GetThisClassName()
	class.superClassName = cf.GetSuperClassName()
	class.interfaceNames = cf.GetInterfaceNames()
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
		class.Methods[i] = newMethod(class, cf, methodInfo)
		class.Methods[i].Slot = uint(i)
	}
}

func (class *Class) copyAttributes(cf *classfile.ClassFile) {
	class.SourceFile = cf.GetUTF8(cf.GetSourceFileIndex()) // TODO
	class.Signature = cf.GetUTF8(cf.GetSignatureIndex())
	class.AnnotationData = cf.GetRuntimeVisibleAnnotationsAttributeData()
	class.EnclosingMethod = getEnclosingMethod(cf)
}

func getEnclosingMethod(cf *classfile.ClassFile) *EnclosingMethod {
	if emAttr, found := cf.GetEnclosingMethodAttribute(); found {
		methodName, methodDescriptor := getNameAndType(cf, emAttr.MethodIndex)
		return &EnclosingMethod{
			ClassName:        cf.GetClassNameOf(emAttr.ClassIndex),
			MethodName:       methodName,
			MethodDescriptor: methodDescriptor,
		}
	}
	return nil
}
