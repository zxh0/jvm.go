package heap

type ConstantString struct {
	goStr string
	jStr  *Object
}

func newConstantString(str string) *ConstantString {
	return &ConstantString{goStr: str}
}

func (s *ConstantString) GetJString() *Object {
	if s.jStr == nil {
		s.jStr = JString(s.goStr)
	}
	return s.jStr
}
