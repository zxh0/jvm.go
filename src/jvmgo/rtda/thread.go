package rtda

import (
    "jvmgo/any"
    rtc "jvmgo/rtda/class"
)

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
    jThread *rtc.Obj // java.lang.Thread
    // todo
}

// getters & setters
func (self *Thread) PC() (int) {
    return self.pc
}
func (self *Thread) SetPC(pc int) {
    self.pc = pc
}
func (self *Thread) JThread() (*rtc.Obj) {
    return self.jThread
}
func (self *Thread) SetJThread(jThread *rtc.Obj) {
    self.jThread = jThread
}

func (self *Thread) IsStackEmpty() (bool) {
    return self.stack.isEmpty()
}

func (self *Thread) CurrentFrame() (*Frame) {
    return self.stack.top()
}
func (self *Thread) TopFrame() (*Frame) {
    return self.stack.top()
}

func (self *Thread) PushFrame(frame *Frame) {
    self.stack.push(frame)
}
func (self *Thread) PopFrame() (*Frame) {
    top := self.stack.pop()
    if top.onPopAction != nil {
        // todo
        top.onPopAction()
    }
    return top
}


func (self *Thread) InvokeMethod(method * rtc.Method) {
    currentFrame := self.CurrentFrame()
    newFrame := self.NewFrame(method)
    self.PushFrame(newFrame)
    _passArgs(currentFrame.OperandStack(), newFrame.LocalVars(), method.ActualArgCount())
}

func _passArgs(stack *OperandStack, vars *LocalVars, argCount uint) {
    if argCount > 0 {
        args := stack.PopN(argCount)
        for i, j := uint(0), uint(0); i < argCount; i++ {
            arg := args[i]
            vars.Set(i + j, arg)
            if any.IsLongOrDouble(arg) {
                j++
            }
        }
    }
}


func (self *Thread) NewFrame(method *rtc.Method) (*Frame) {
    return newFrame(self, method)
}

func NewThread(maxStackSize int, jThread *rtc.Obj) (*Thread) {
    stack := newStack(maxStackSize)
    return &Thread{0, stack, jThread}
}
