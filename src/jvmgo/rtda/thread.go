package rtda

/*
JVM
  Thread
    pc
    Stack
      Frame
        LocalVars
        OperandStack
*/
type Thread struct {
    pc      int
    stack   *Stack
    // todo
}

func (self *Thread) PC() (int) {
    return self.pc
}

func (self *Thread) IsStackEmpty() (bool) {
    return self.stack.isEmpty()
}

func (self *Thread) CurrentFrame() (*Frame) {
    return self.stack.top()
}

func newThread(maxStackSize int) (*Thread) {
    stack := newStack(maxStackSize)
    return &Thread{0, stack}
}
