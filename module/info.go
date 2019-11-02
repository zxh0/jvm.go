package module

import (
	"github.com/zxh0/jvm.go/classfile"
)

type Info struct {
	Name     string
	Flags    uint16
	Version  string
	Requires []Require
	Exports  []Export
	Opens    []Open
	Uses     []string
	Provides []Provide
}

type Require struct {
	Name    string
	Flags   uint16
	Version string
}

type Export struct {
	Package string
	Flags   uint16
	To      []string
}

type Open struct {
	Package string
	Flags   uint16
	To      []string
}

type Provide struct {
	Service string
	Impls   []string
}

func ParseModuleInfo(classData []byte) *Info {
	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err) // TODO
	}

	modAttr, ok := cf.GetModuleAttribute()
	if !ok {
		panic("no module attribute")
	}

	return &Info{
		Name:     cf.GetModuleName(modAttr.ModuleNameIndex),
		Flags:    modAttr.ModuleFlags,
		Version:  cf.GetRawUTF8(modAttr.ModuleVersionIndex),
		Requires: getRequires(cf, modAttr),
		Exports:  getExports(cf, modAttr),
		Opens:    getOpens(cf, modAttr),
		Uses:     cf.GetClassNames(modAttr.UsesIndexTable),
		Provides: getProvides(cf, modAttr),
	}
}

func getRequires(cf *classfile.ClassFile,
	modAttr classfile.ModuleAttribute) []Require {

	requires := make([]Require, len(modAttr.RequiresTable))
	for i, cfRequire := range modAttr.RequiresTable {
		requires[i] = Require{
			Name:    cf.GetModuleName(cfRequire.RequiresIndex),
			Flags:   cfRequire.RequiresFlags,
			Version: cf.GetRawUTF8(cfRequire.RequiresVersionIndex),
		}
	}

	return requires
}

func getExports(cf *classfile.ClassFile,
	modAttr classfile.ModuleAttribute) []Export {

	exports := make([]Export, len(modAttr.ExportsTable))
	for i, cfExport := range modAttr.ExportsTable {
		exports[i] = Export{
			Package: cf.GetPackageName(cfExport.ExportsIndex),
			Flags:   cfExport.ExportsFlags,
			To:      cf.GetModuleNames(cfExport.ExportsToIndexTable),
		}
	}

	return exports
}

func getOpens(cf *classfile.ClassFile,
	modAttr classfile.ModuleAttribute) []Open {

	opens := make([]Open, len(modAttr.OpensTable))
	for i, cfOpen := range modAttr.OpensTable {
		opens[i] = Open{
			Package: cf.GetPackageName(cfOpen.OpensIndex),
			Flags:   cfOpen.OpensFlags,
			To:      cf.GetModuleNames(cfOpen.OpensToIndexTable),
		}
	}

	return opens
}

func getProvides(cf *classfile.ClassFile,
	modAttr classfile.ModuleAttribute) []Provide {

	provides := make([]Provide, len(modAttr.ProvidesTable))
	for i, cfProvide := range modAttr.ProvidesTable {
		provides[i] = Provide{
			Service: cf.GetClassName(cfProvide.ProvidesIndex),
			Impls:   cf.GetClassNames(cfProvide.ProvidesWithIndexTable),
		}
	}

	return provides
}
