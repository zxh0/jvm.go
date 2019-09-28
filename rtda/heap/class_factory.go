package heap

import (
	"sync"

	"github.com/zxh0/jvm.go/classfile"
)

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.initCond = sync.NewCond(&sync.Mutex{})
	class.AccessFlags = AccessFlags(cf.AccessFlags)
	class.copyConstantPool(cf)
	class.copyClassNames(cf)
	class.copyFields(cf)
	class.copyMethods(cf)
	class.copyAttributes(cf)
	return class
}

func (class *Class) copyConstantPool(cf *classfile.ClassFile) {
	class.constantPool = newConstantPool(class, cf)
}

func (class *Class) copyClassNames(cf *classfile.ClassFile) {
	class.name = cf.GetClassName()
	class.superClassName = cf.GetSuperClassName()
	class.interfaceNames = cf.GetInterfaceNames()
}

func (class *Class) copyFields(cf *classfile.ClassFile) {
	class.fields = make([]*Field, len(cf.Fields))
	for i, fieldInfo := range cf.Fields {
		class.fields[i] = newField(class, cf, fieldInfo)
	}
}

func (class *Class) copyMethods(cf *classfile.ClassFile) {
	class.methods = make([]*Method, len(cf.Methods))
	for i, methodInfo := range cf.Methods {
		class.methods[i] = newMethod(class, cf, methodInfo)
		class.methods[i].slot = uint(i)
	}
}

func (class *Class) copyAttributes(cf *classfile.ClassFile) {
	class.sourceFile = cf.GetUTF8(cf.GetSourceFileIndex()) // TODO
	class.signature = cf.GetUTF8(cf.GetSignatureIndex())
	class.annotationData = cf.GetRuntimeVisibleAnnotationsAttributeData()
	class.enclosingMethod = getEnclosingMethod(cf)
}

func getEnclosingMethod(cf *classfile.ClassFile) *EnclosingMethod {
	if emAttr, found := cf.GetEnclosingMethodAttribute(); found {
		methodName, methodDescriptor := getNameAndType(cf, emAttr.MethodIndex)
		return &EnclosingMethod{
			className:        cf.GetClassNameOf(emAttr.ClassIndex),
			methodName:       methodName,
			methodDescriptor: methodDescriptor,
		}
	}
	return nil
}
