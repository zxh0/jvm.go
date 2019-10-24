package module

import (
	"github.com/zxh0/jvm.go/vmutils"
)

// JAR with module-info.class
type ModularJAR struct {
	jar  *vmutils.ZipFile
	info *Info
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

	return &ModularJAR{
		jar:  jar,
		info: ParseModuleInfo(classData),
	}
}

func (module *ModularJAR) GetInfo() *Info {
	return module.info
}
