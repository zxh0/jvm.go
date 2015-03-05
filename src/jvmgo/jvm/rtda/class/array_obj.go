package class

func (self *Obj) IsArray() bool {
	return self.class.IsArray()
}
func (self *Obj) IsPrimitiveArray() bool {
	return self.class.IsPrimitiveArray()
}
