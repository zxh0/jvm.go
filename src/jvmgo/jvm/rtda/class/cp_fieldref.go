package class

import (
	"fmt"
	cf "jvmgo/classfile"
	"jvmgo/util"
)

type ConstantFieldref struct {
	className  string
	name       string
	descriptor string
	cp         *ConstantPool
	field      *Field
}

func newConstantFieldref(cp *ConstantPool, fieldrefInfo *cf.ConstantFieldrefInfo) *ConstantFieldref {
	fieldref := &ConstantFieldref{}
	fieldref.className = fieldrefInfo.ClassName()
	fieldref.name = fieldrefInfo.Name()
	fieldref.descriptor = fieldrefInfo.Descriptor()
	fieldref.cp = cp
	return fieldref
}

func (self *ConstantFieldref) String() string {
	return fmt.Sprintf("{ConstantFieldref className:%v name:%v descriptor:%v}",
		self.className, self.name, self.descriptor)
}

func (self *ConstantFieldref) InstanceField() *Field {
	if self.field == nil {
		self.resolveField(false)
	}
	return self.field
}

func (self *ConstantFieldref) StaticField() *Field {
	if self.field == nil {
		self.resolveField(true)
	}
	return self.field
}

func (self *ConstantFieldref) resolveField(isStatic bool) {
	classLoader := self.cp.class.classLoader
	fromClass := classLoader.LoadClass(self.className)

	for class := fromClass; class != nil; class = class.superClass {
		field := class.getField(self.name, self.descriptor, isStatic)
		if field != nil {
			self.field = field
			return
		}
	}

	// todo
	util.Panicf("field not found! %v", self)
}
