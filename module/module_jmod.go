package module

import (
	"github.com/zxh0/jvm.go/vmutils"
)

type JModModule struct {
	BaseModule
	jmod *vmutils.JModFile
}

func NewJModModule(path string) *JModModule {
	jmod, err := vmutils.OpenJModFile(path)
	if err != nil {
		panic(err) // TODO
	}

	classData, err := jmod.ReadFile("classes/module-info.class")
	if err != nil {
		panic(err) // TODO
	}

	jmod.Close()
	return &JModModule{
		jmod: jmod,
		BaseModule: BaseModule{
			info: ParseModuleInfo(classData),
		},
	}
}

func (m *JModModule) ReadClass(name string) ([]byte, error) {
	if !m.jmod.IsOpen() {
		if err := m.jmod.Open(); err != nil {
			return nil, err
		}
	}
	return m.jmod.ReadFile("classes/" + name + ".class")
}
