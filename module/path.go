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
func (path Path) ReadClass(className string) (Module, []byte) {
	for _, m := range path {
		if data, err := m.ReadClass(className); err == nil {
			return m, data
		}
	}
	//panic("class not found:" + className)
	return nil, nil
}

func (path Path) GetModuleByPackageName(pkgName string) Module {
	for _, module := range path {
		if module.HasPackage(pkgName) {
			return module
		}
	}
	return nil
}

func (path Path) getModuleByName(name string) Module {
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
