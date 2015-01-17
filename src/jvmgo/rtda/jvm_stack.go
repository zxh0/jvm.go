package rtda

// jvm stack
type Stack struct {
    size    uint16 
    frames  []*Frame
}

func (self *Stack) push(frame *Frame) {
    self.frames[self.size] = frame
    self.size++
}

func (self *Stack) pop() {
    self.frames[self.size] = nil
    self.size--
}
