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
    pc      uint32
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

func newThread(maxStackSize int) (*Thread) {
    stack := newStack(maxStackSize)
    return &Thread{0, stack}
}
