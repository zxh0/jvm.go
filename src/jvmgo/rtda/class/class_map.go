package class

type ClassMap struct {
    _map map[string]*Class
}

func (self ClassMap) getClass(name string) (*Class) {
    return self._map[name]
}

func (self ClassMap) putClass(name string, class *Class) {
    self._map[name] = class
}
