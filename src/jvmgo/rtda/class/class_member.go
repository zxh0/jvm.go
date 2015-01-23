package class

type ClassMember struct {
    AccessFlags
    name        string
    descriptor  string
    class       *Class
}
