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
    classLoader := self.class.classLoader
    k := self.class 
    for {
        if k == class {
            return true
        }
        if k.superClassName != "" {
            k = classLoader.getClass(k.superClassName)
        } else {
            break;
        }
    }

    return false
}

// todo
func NewObj(fields Any) (*Obj) {
    return &Obj{nil, fields}
}
