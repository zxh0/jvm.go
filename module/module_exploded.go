package module

import (
	"github.com/zxh0/jvm.go/vmutils"
)

type ExplodedModule struct {
	dir  *vmutils.Dir
	info *Info
}

func NewExplodedModule(path string) *ExplodedModule {
	dir, err := vmutils.NewDir(path)
	if err != nil {
		panic(err) // TODO
	}

	classData, err := dir.ReadFile("module-info.class")
	if err != nil {
		panic(err) // TODO
	}

	return &ExplodedModule{
		dir:  dir,
		info: ParseModuleInfo(classData),
	}
}

func (module *ExplodedModule) GetInfo() *Info {
	return module.info
}
