package module

// JAR without module-info.class
type AutomaticModule struct {
	BaseModule
	// TODO
}

func NewAutomaticModule(path string) *AutomaticModule {
	return &AutomaticModule{
		// TODO
	}
}

func (m *AutomaticModule) ReadClass(name string) ([]byte, error) {
	panic("TODO")
}
