package rtda

// jvm stack
type Stack struct {
    maxSize int
    size    int 
    frames  []*Frame
}

func (self *Stack) push(frame *Frame) {
    if self.size >= self.maxSize {
        // todo
        panic("StackOverflowError")
    }
    if self.size == len(self.frames) {
        // todo
    }

    self.frames[self.size] = frame
    self.size++
}

func (self *Stack) pop() {
    self.size--
    self.frames[self.size] = nil
}

func newStack(maxSize int) (*Stack) {
    frames := make([]*Frame, 8)
    return &Stack{maxSize, 0, frames}
}
