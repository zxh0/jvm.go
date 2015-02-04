package class

import cf "jvmgo/classfile"

type ConstantString struct {
    goStr   string
    jStr    *Obj
}

func newConstantString(stringInfo *cf.ConstantStringInfo) (*ConstantString) {
    goStr := stringInfo.String()
    return &ConstantString{goStr, nil}
}

// getters & setters
func (self *ConstantString) GoStr() (string) {
    return self.goStr
}
func (self *ConstantString) JStr() (*Obj) {
    return self.jStr
}
func (self *ConstantString) SetJStr(jStr *Obj) {
    self.jStr = jStr
}
