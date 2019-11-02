package module

// TODO
type Module interface {
	GetInfo() *Info
	GetName() string
	GetVersion() string
	GetPath() string
	HasPackage(pkgName string) bool
	ReadClass(className string) ([]byte, error)
}

type BaseModule struct {
	info *Info
	pkgs map[string]bool
}

func (bm *BaseModule) GetInfo() *Info {
	return bm.info
}

func (bm *BaseModule) GetName() string {
	return bm.info.Name
}
func (bm *BaseModule) GetVersion() string {
	return bm.info.Version
}

func (bm *BaseModule) HasPackage(pkgName string) bool {
	return bm.pkgs[pkgName]
}
