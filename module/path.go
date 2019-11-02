package module

import (
	"sort"
	"strings"
)

type Path []Module

func (path Path) Sort() {
	sort.Sort(path)
}

// TODO
func (path Path) ReadClass(className string) (string, []byte) {
	for _, m := range path {
		if data, err := m.ReadClass(className); err == nil {
			return "todo", data
		}
	}
	panic("class not found:" + className)
}

func (path Path) findModule(name string) Module {
	for _, module := range path {
		if module.GetName() == name {
			return module
		}
	}
	return nil
}

// sort.Interface
func (path Path) Len() int {
	return len(path)
}
func (path Path) Less(i, j int) bool {
	name1 := path[i].GetName()
	name2 := path[j].GetName()
	return strings.Compare(name1, name2) < 0
}
func (path Path) Swap(i, j int) {
	path[i], path[j] = path[j], path[i]
}
