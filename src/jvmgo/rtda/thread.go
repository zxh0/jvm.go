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
    pc      uint
    stack   *Stack
    // todo
}

// todo
func (self *Thread) loop() {
    for !self.stack.isEmpty() {
        currentFrame := self.stack.top()
        currentFrame.executeOneInstruction()
    }
}

func (self *Thread) CurrentFrame() (*Frame) {
    return self.stack.top()
}

func newThread(maxStackSize int) (*Thread) {
    stack := newStack(maxStackSize)
    return &Thread{0, stack}
}
