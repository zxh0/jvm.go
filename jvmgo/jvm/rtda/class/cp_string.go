package class

import cf "github.com/zxh0/jvm.go/jvmgo/classfile"

type ConstantString struct {
	goStr string
	jStr  *Obj
}

func newConstantString(stringInfo *cf.ConstantStringInfo) *ConstantString {
	return &ConstantString{
		goStr: stringInfo.String(),
	}
}

// getters & setters
func (self *ConstantString) GoStr() string {
	return self.goStr
}
func (self *ConstantString) JStr() *Obj {
	return self.jStr
}
func (self *ConstantString) SetJStr(jStr *Obj) {
	self.jStr = jStr
}
