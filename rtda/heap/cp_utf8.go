package heap

type ConstantUtf8 struct {
	str string
}

func newConstantUtf8(utf8Str string) *ConstantUtf8 {
	return &ConstantUtf8{
		str: utf8Str,
	}
}

func (self *ConstantUtf8) Str() string {
	return self.str
}
