package module

import (
	"sort"
	"strings"
)

type Path []Module

func (path Path) Sort() {
	sort.Sort(path)
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
