package module

// TODO
type Module interface {
	GetInfo() *Info
	GetName() string
	GetVersion() string
	ReadClass(className string) ([]byte, error)
}

type BaseModule struct {
	info *Info
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
