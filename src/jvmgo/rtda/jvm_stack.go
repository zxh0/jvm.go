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
        self.expand()
    }

    self.frames[self.size] = frame
    self.size++
}

func (self *Stack) expand() {
    newLen := self.size + 8
    if newLen > self.maxSize {
        newLen = self.maxSize
    }
    newFrames := make([]*Frame, newLen)
    copy(newFrames, self.frames) // func copy(dst, src []T) int
    self.frames = newFrames
}

func (self *Stack) pop() {
    self.size--
    self.frames[self.size] = nil
}

func (self *Stack) top() (*Frame) {
    return self.frames[self.size - 1]
}

func (self *Stack) isEmpty() (bool) {
    return self.size > 0
}

func newStack(maxSize int) (*Stack) {
    frames := make([]*Frame, 8)
    return &Stack{maxSize, 0, frames}
}
