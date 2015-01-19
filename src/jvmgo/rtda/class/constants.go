package class

type ConstantInt struct {
    val int32
}
func (self *ConstantInt) Val() (int32) {
    return self.val
}

type ConstantLong struct {
    val int64
}
func (self *ConstantLong) Val() (int64) {
    return self.val
}

type ConstantFloat struct {
    val float32
}
func (self *ConstantFloat) Val() (float32) {
    return self.val
}

type ConstantDouble struct {
    val float64
}
func (self *ConstantDouble) Val() (float64) {
    return self.val
}
