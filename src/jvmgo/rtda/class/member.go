package class

type ClassMember struct {
    AccessFlags
    name        string
    descriptor  string
    class       *Class
}

func (self *ClassMember) Name() (string) {
    return self.name
}
func (self *ClassMember) Class() (*Class) {
    return self.class
}
