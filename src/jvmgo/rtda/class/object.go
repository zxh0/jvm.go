package class

import . "jvmgo/any"

//type Ref *Obj

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

// todo
func NewObj(fields Any) (*Obj) {
    return &Obj{nil, fields}
}
