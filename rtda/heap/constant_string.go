package heap

type ConstantString struct {
	goStr string
	class *Class
	jStr  *Object
}

func newConstantString(class *Class, str string) *ConstantString {
	return &ConstantString{class: class, goStr: str}
}

func (s *ConstantString) GetJString() *Object {
	if s.jStr == nil {
		rt := s.class.bootLoader.rt
		s.jStr = rt.JSFromGoStr(s.goStr)
	}
	return s.jStr
}
