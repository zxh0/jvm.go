package class

func (self *Class) isSubClassOf(c *Class) bool {
    for k := self.superClass; k != nil; k = k.superClass {
        if k == c {
            return true
        }
    }
    return false
}
func (self *Class) isSuperClassOf(c *Class) bool {
    return c.isSubClassOf(self)
}
