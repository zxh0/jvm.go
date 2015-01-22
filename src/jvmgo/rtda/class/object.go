package class

import . "jvmgo/any"

//type Ref *Obj

// object
type Obj struct {
    fields Any // []Any
}

func (self *Obj) Fields() (Any) {
    return self.fields
}

func NewObj(fields Any) (*Obj) {
    return &Obj{fields}
}
