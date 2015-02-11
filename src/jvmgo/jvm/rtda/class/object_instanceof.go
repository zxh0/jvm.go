package class

func (self *Obj) IsInstanceOf(class *Class) (bool) {
    if class.IsInterface() {
        for k := self.class; k != nil; k = k.superClass {
            for _, i := range k.interfaces {
                if _interfaceXextendsY(i, class) {
                    return true
                }
            }
        }
    } else {
        for k := self.class; k != nil; k = k.superClass {
            if k == class {
                return true
            }
        }
    }
    return false
}
func _interfaceXextendsY(x, y *Class) (bool) {
    if x == y {
        return true
    }
    for _, superInterface := range x.interfaces {
        if _interfaceXextendsY(superInterface, y) {
            return true
        }
    }
    return false
}
