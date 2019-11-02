package module

import (
	"github.com/zxh0/jvm.go/vmutils"
)

// JAR with module-info.class
type ModularJAR struct {
	BaseModule
	jar *vmutils.ZipFile
}

func NewModularJAR(path string) *ModularJAR {
	jar, err := vmutils.OpenZipFile(path)
	if err != nil {
		panic(err) // TODO
	}
	defer jar.Close()

	classData, err := jar.ReadFile("module-info.class")
	if err != nil {
		panic(err) // TODO
	}

	pkgs := map[string]bool{}
	for _, path := range jar.ListFiles() {
		if vmutils.IsClassFile(path) {
			pkgName := vmutils.StripFileName(path)
			pkgs[pkgName] = true
		}
	}

	return &ModularJAR{
		jar: jar,
		BaseModule: BaseModule{
			info: ParseModuleInfo(classData),
			pkgs: pkgs,
		},
	}
}

func (m *ModularJAR) GetPath() string {
	return "file:" + m.jar.AbsPath()
}

func (m *ModularJAR) ReadClass(name string) ([]byte, error) {
	if !m.jar.IsOpen() {
		if err := m.jar.Open(); err != nil {
			return nil, err
		}
	}
	return m.jar.ReadFile(name + ".class")
}
