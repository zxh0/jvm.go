package class

type ClassLoader interface {
    LoadClass(name string) (*Class)
}
