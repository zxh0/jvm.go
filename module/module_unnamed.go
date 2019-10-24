package module

type UnnamedModule struct {
	// TODO
	info *Info
}

func (module UnnamedModule) GetInfo() *Info {
	return module.info
}
