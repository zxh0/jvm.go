package class

import (
	"jvmgo/classfile"
	"jvmgo/util"
)

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
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
	}
}

func (self *Class) copyAttributes(cf *classfile.ClassFile) {
	self.attributes = &Attributes{}

	if sfAttr := cf.SourceFileAttribute(); sfAttr != nil {
		self.attributes.sourceFile = sfAttr.FileName()
	} else {
		self.attributes.sourceFile = "Unknown" // todo
	}
	
	if rvaAttr := cf.RuntimeVisibleAnnotationsAttribute(); rvaAttr != nil {
		self.attributes.annotationData = util.CastUint8sToInt8s(rvaAttr.Info())
	}
}
