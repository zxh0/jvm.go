package class

type ConstantInt struct {
    val int32
}
func (self *ConstantInt) Val() (int32) {
    return self.val
}

type ConstantFloat struct {
    val float32
}
func (self *ConstantFloat) Val() (float32) {
    return self.val
}
