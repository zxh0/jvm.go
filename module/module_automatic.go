package module

// JAR without module-info.class
type AutomaticModule struct {
	info *Info
	// TODO
}

func NewAutomaticModule(path string) *AutomaticModule {
	return &AutomaticModule{
		// TODO
	}
}

func (module *AutomaticModule) GetInfo() *Info {
	return module.info
}
