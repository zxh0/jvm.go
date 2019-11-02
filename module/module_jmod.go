package module

import (
	"strings"

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
	defer jmod.Close()

	classData, err := jmod.ReadFile("classes/module-info.class")
	if err != nil {
		panic(err) // TODO
	}

	pkgs := map[string]bool{}
	for _, path := range jmod.ListFiles() {
		if vmutils.IsClassFile(path) {
			pkgName := vmutils.StripFileName(path)
			pkgName = strings.TrimPrefix(pkgName, "classes/")
			pkgs[pkgName] = true
		}
	}

	return &JModModule{
		jmod: jmod,
		BaseModule: BaseModule{
			info: ParseModuleInfo(classData),
			pkgs: pkgs,
		},
	}
}

func (m *JModModule) GetPath() string {
	return "file:" + m.jmod.AbsPath()
}

func (m *JModModule) HasPackage(pkgName string) bool {
	return m.pkgs[pkgName]
}

func (m *JModModule) ReadClass(name string) ([]byte, error) {
	if !m.jmod.IsOpen() {
		if err := m.jmod.Open(); err != nil {
			return nil, err
		}
	}
	return m.jmod.ReadFile("classes/" + name + ".class")
}
