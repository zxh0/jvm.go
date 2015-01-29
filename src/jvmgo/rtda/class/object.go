package class

import . "jvmgo/any"

// object
type Obj struct {
    class   *Class
    fields  Any // []Any
}

func (self *Obj) Class() (*Class) {
    return self.class
}
func (self *Obj) Fields() (Any) {
    return self.fields
}

func (self *Obj) IsInstanceOf(class *Class) (bool) {
    // todo
    return self.class == class
}

// todo
func NewObj(fields Any) (*Obj) {
    return &Obj{nil, fields}
}
