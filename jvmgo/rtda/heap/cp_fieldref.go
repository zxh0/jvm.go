package heap

import (
	"fmt"

	cf "github.com/zxh0/jvm.go/jvmgo/classfile"
	"github.com/zxh0/jvm.go/jvmgo/jutil"
)

type ConstantFieldref struct {
	ConstantMemberref
	field *Field
}

func newConstantFieldref(refInfo *cf.ConstantFieldrefInfo) *ConstantFieldref {
	ref := &ConstantFieldref{}
	ref.copy(&refInfo.ConstantMemberrefInfo)
	return ref
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
	jutil.Panicf("instance field not found! %v", self)
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
		if self._findInterfaceField(class) {
			return
		}
	}

	// todo
	jutil.Panicf("static field not found! %v", self)
}

func (self *ConstantFieldref) _findInterfaceField(class *Class) bool {
	for _, iface := range class.interfaces {
		for _, f := range iface.fields {
			if f.name == self.name && f.descriptor == self.descriptor {
				self.field = f
				return true
			}
		}
		if self._findInterfaceField(iface) {
			return true
		}
	}
	return false
}
