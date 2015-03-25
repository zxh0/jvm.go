package class

import (
	"fmt"
	cf "github.com/zxh0/jvm.go/jvmgo/classfile"
	"github.com/zxh0/jvm.go/jvmgo/util"
)

type ConstantFieldref struct {
	className  string
	name       string
	descriptor string
	cp         *ConstantPool
	field      *Field
}

func newConstantFieldref(cp *ConstantPool, fieldrefInfo *cf.ConstantFieldrefInfo) *ConstantFieldref {
	return &ConstantFieldref{
		className:  fieldrefInfo.ClassName(),
		name:       fieldrefInfo.Name(),
		descriptor: fieldrefInfo.Descriptor(),
		cp:         cp,
	}
}

func (self *ConstantFieldref) String() string {
	return fmt.Sprintf("{ConstantFieldref className:%v name:%v descriptor:%v}",
		self.className, self.name, self.descriptor)
}

func (self *ConstantFieldref) InstanceField() *Field {
	if self.field == nil {
		self.resolveInstanceField()
	}
	return self.field
}
func (self *ConstantFieldref) resolveInstanceField() {
	fromClass := bootLoader.LoadClass(self.className)

	for class := fromClass; class != nil; class = class.superClass {
		field := class.getField(self.name, self.descriptor, false)
		if field != nil {
			self.field = field
			return
		}
	}

	// todo
	util.Panicf("instance field not found! %v", self)
}

func (self *ConstantFieldref) StaticField() *Field {
	if self.field == nil {
		self.resolveStaticField()
	}
	return self.field
}
func (self *ConstantFieldref) resolveStaticField() {
	fromClass := bootLoader.LoadClass(self.className)

	for class := fromClass; class != nil; class = class.superClass {
		field := class.getField(self.name, self.descriptor, true)
		if field != nil {
			self.field = field
			return
		}
	}

	// todo
	util.Panicf("static field not found! %v", self)
}
