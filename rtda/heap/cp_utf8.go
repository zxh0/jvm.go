package heap

type ConstantUtf8 struct {
	str string
}

func newConstantUtf8(utf8Str string) *ConstantUtf8 {
	return &ConstantUtf8{
		str: utf8Str,
	}
}

func (utf8 *ConstantUtf8) Str() string {
	return utf8.str
}
