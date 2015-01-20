package class

import . "jvmgo/any"

type Field struct {
    //name    string
    class   *Class
    slot    uint
}

// getters
func (self *Field) Class() (*Class) {
    return self.class
}

func (self *Field) GetValue(ref *Obj) (Any) {
    fields := ref.fields.([]Any)
    return fields[self.slot]
}

func (self *Field) GetStaticValue() (Any) {
    fields := self.class.fields.([]Any)
    return fields[self.slot]
}
