package class

// jvms8-6.5.instanceof
func (self *Obj) IsInstanceOf(class *Class) (bool) {
    s := self.class
    t := class

    if s == t {
        return true
    }

    if !s.IsArray() {
        if !s.IsInterface() {
            if !t.IsInterface() {
                return s.isSubClassOf(t)
            } else {
                return s.isImplements(t)
            }
        } else {
            if !t.IsInterface() {
                return t.isObject()
            } else {
                return t.isSuperInterfaceOf(t)
            }
        }
    } else {
        if !t.IsArray() {
            if !t.IsInterface() {
                return t.isObject()
            } else {
                // todo
            }
        } else {
            // todo
        }
    }


    // if class.IsInterface() {
    //     for k := self.class; k != nil; k = k.superClass {
    //         for _, i := range k.interfaces {
    //             if i.isSubInterfaceOf(class) {
    //                 return true
    //             }
    //         }
    //     }
    // } else {
    //     for k := self.class; k != nil; k = k.superClass {
    //         if k == class {
    //             return true
    //         }
    //     }
    // }
    return false
}

// self implements iface
func (self *Class) isImplements(iface *Class) bool {
    for k := self; k != nil; k = k.superClass {
        for _, i := range k.interfaces {
            if i == iface || i.isSubInterfaceOf(iface) {
                return true
            }
        }
    }
    return false
}

// iface extends self
func (self *Class) isSuperInterfaceOf(iface *Class) bool {
    return iface.isSubInterfaceOf(self)
}
// self extends iface
func (self *Class) isSubInterfaceOf(iface *Class) bool {
    for _, superInterface := range self.interfaces {
        if superInterface == iface || superInterface.isSubInterfaceOf(iface) {
            return true
        }
    }
    return false
}
