package heap

import (
	"sync"

	"github.com/zxh0/jvm.go/jvmgo/classfile"
)

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.initCond = sync.NewCond(&sync.Mutex{})
	class.accessFlags = cf.AccessFlags()
	class.copyConstantPool(cf)
	class.copyClassNames(cf)
	class.copyFields(cf)
	class.copyMethods(cf)
	class.copyAttributes(cf)
	return class
}

func (self *Class) copyConstantPool(cf *classfile.ClassFile) {
	self.constantPool = newConstantPool(self, cf.ConstantPool())
}

func (self *Class) copyClassNames(cf *classfile.ClassFile) {
	self.name = cf.ClassName()
	self.superClassName = cf.SuperClassName()
	self.interfaceNames = cf.InterfaceNames()
}

func (self *Class) copyFields(cf *classfile.ClassFile) {
	self.fields = make([]*Field, len(cf.Fields()))
	for i, fieldInfo := range cf.Fields() {
		self.fields[i] = newField(self, fieldInfo)
	}
}

func (self *Class) copyMethods(cf *classfile.ClassFile) {
	self.methods = make([]*Method, len(cf.Methods()))
	for i, methodInfo := range cf.Methods() {
		self.methods[i] = newMethod(self, methodInfo)
		self.methods[i].slot = uint(i)
	}
}

func (self *Class) copyAttributes(cf *classfile.ClassFile) {
	self.sourceFile = getSourceFile(cf)
	self.signature = getSignature(cf)
	self.annotationData = cf.RuntimeVisibleAnnotationsAttributeData()
	self.enclosingMethod = getEnclosingMethod(cf)
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
