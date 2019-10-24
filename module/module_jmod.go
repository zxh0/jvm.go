package module

import (
	"github.com/zxh0/jvm.go/vmutils"
)

type JModModule struct {
	jmod *vmutils.JModFile
	info *Info
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
		info: ParseModuleInfo(classData),
	}
}

func (module *JModModule) GetInfo() *Info {
	return module.info
}
