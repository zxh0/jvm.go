package module

import (
	"github.com/zxh0/jvm.go/vmutils"
)

type ExplodedModule struct {
	BaseModule
	dir *vmutils.Dir
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
		dir: dir,
		BaseModule: BaseModule{
			info: ParseModuleInfo(classData),
		},
	}
}

func (m *ExplodedModule) ReadClass(name string) ([]byte, error) {
	return m.dir.ReadFile(name + ".class")
}
