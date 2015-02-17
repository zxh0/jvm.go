package rtda

func (self *Thread) ThrowNPE() {
	self.ThrowException("java/lang/NullPointerException", "()V", nil)
}
