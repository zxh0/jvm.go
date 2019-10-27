package module

import (
	"strings"

	"github.com/zxh0/jvm.go/vmutils"
)

type Graph struct {
	mList []Module
	mMap  map[string]Module // module by exported package
}

func newGraph(mList []Module) Graph {
	mMap := map[string]Module{}
	for _, m := range mList {
		for _, export := range m.GetInfo().Exports {
			mMap[vmutils.DotToSlash(export.Package)] = m
		}
	}

	return Graph{
		mList: mList,
		mMap:  mMap,
	}
}

func (g Graph) GetModules() []Module {
	return g.mList
}

func (g Graph) ReadClass(className string) (string, []byte) {
	pkgName := className[:strings.LastIndexByte(className, '/')]

	module, found := g.mMap[pkgName]
	if !found {
		panic("class not found:" + className)
	}

	//println("read class " + className + " from " + module.GetName())
	data, err := module.ReadClass(className)
	if err != nil {
		panic("class not found:" + className)
	}

	return "todo", data
}
