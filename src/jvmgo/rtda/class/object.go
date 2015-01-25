package class

import . "jvmgo/any"

//type Ref *Obj

// object
type Obj struct {
    class   *Class
    fields  Any // []Any
}

func (self *Obj) Fields() (Any) {
    return self.fields
}
