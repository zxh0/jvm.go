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
		Name:     getModuleName(cf, modAttr.ModuleNameIndex),
		Flags:    modAttr.ModuleFlags,
		Version:  cf.GetUTF8(modAttr.ModuleVersionIndex),
		Requires: getRequires(cf, modAttr),
		Exports:  getExports(cf, modAttr),
		Opens:    getOpens(cf, modAttr),
		Uses:     getClassNames(cf, modAttr.UsesIndexTable),
		Provides: getProvides(cf, modAttr),
	}
}

func getRequires(cf *classfile.ClassFile,
	modAttr classfile.ModuleAttribute) []Require {

	requires := make([]Require, len(modAttr.RequiresTable))
	for i, cfRequire := range modAttr.RequiresTable {
		requires[i] = Require{
			Name:    getModuleName(cf, cfRequire.RequiresIndex),
			Flags:   cfRequire.RequiresFlags,
			Version: cf.GetUTF8(cfRequire.RequiresVersionIndex),
		}
	}

	return requires
}

func getExports(cf *classfile.ClassFile,
	modAttr classfile.ModuleAttribute) []Export {

	exports := make([]Export, len(modAttr.ExportsTable))
	for i, cfExport := range modAttr.ExportsTable {
		exports[i] = Export{
			Package: getPackageName(cf, cfExport.ExportsIndex),
			Flags:   cfExport.ExportsFlags,
			To:      getModuleNames(cf, cfExport.ExportsToIndexTable),
		}
	}

	return exports
}

func getOpens(cf *classfile.ClassFile,
	modAttr classfile.ModuleAttribute) []Open {

	opens := make([]Open, len(modAttr.OpensTable))
	for i, cfOpen := range modAttr.OpensTable {
		opens[i] = Open{
			Package: getPackageName(cf, cfOpen.OpensIndex),
			Flags:   cfOpen.OpensFlags,
			To:      getModuleNames(cf, cfOpen.OpensToIndexTable),
		}
	}

	return opens
}

func getProvides(cf *classfile.ClassFile,
	modAttr classfile.ModuleAttribute) []Provide {

	provides := make([]Provide, len(modAttr.ProvidesTable))
	for i, cfProvide := range modAttr.ProvidesTable {
		provides[i] = Provide{
			Service: getClassName(cf, cfProvide.ProvidesIndex),
			Impls:   getClassNames(cf, cfProvide.ProvidesWithIndexTable),
		}
	}

	return provides
}

func getModuleNames(cf *classfile.ClassFile, cpIdxes []uint16) []string {
	ss := make([]string, len(cpIdxes))
	for i, cpIdx := range cpIdxes {
		ss[i] = getModuleName(cf, cpIdx)
	}
	return ss
}

func getClassNames(cf *classfile.ClassFile, cpIdxes []uint16) []string {
	ss := make([]string, len(cpIdxes))
	for i, cpIdx := range cpIdxes {
		ss[i] = getClassName(cf, cpIdx)
	}
	return ss
}

func getModuleName(cf *classfile.ClassFile, cpIdx uint16) string {
	modInfo := cf.GetConstantInfo(cpIdx).(classfile.ConstantModuleInfo)
	return cf.GetUTF8(modInfo.NameIndex)
}

func getPackageName(cf *classfile.ClassFile, cpIdx uint16) string {
	pkgInfo := cf.GetConstantInfo(cpIdx).(classfile.ConstantPackageInfo)
	return cf.GetUTF8(pkgInfo.NameIndex)
}

func getClassName(cf *classfile.ClassFile, cpIdx uint16) string {
	clsInfo := cf.GetConstantInfo(cpIdx).(classfile.ConstantClassInfo)
	return cf.GetUTF8(clsInfo.NameIndex)
}
