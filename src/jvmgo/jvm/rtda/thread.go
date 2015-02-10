package rtda

import (
    "log"
    "strings"
    "jvmgo/any"
    rtc "jvmgo/jvm/rtda/class"
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
    pc      int // the address of the instruction currently being executed
    stack   *Stack
    jThread *rtc.Obj // java.lang.Thread
    // todo
}

func NewThread(maxStackSize uint, jThread *rtc.Obj) (*Thread) {
    stack := newStack(maxStackSize)
    return &Thread{0, stack, jThread}
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
func (self *Thread) StackDepth() uint {
    return self.stack.size
}

func (self *Thread) CurrentFrame() (*Frame) {
    return self.stack.top()
}
func (self *Thread) TopFrame() (*Frame) {
    return self.stack.top()
}
func (self *Thread) TopFrameN(n uint) (*Frame) {
    return self.stack.topN(n)
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

func (self *Thread) NewFrame(method *rtc.Method) (*Frame) {
    return newFrame(self, method)
}


func (self *Thread) InvokeMethod(method * rtc.Method) {
    //_logInvoke(self.stack.size, method)
    currentFrame := self.CurrentFrame()
    newFrame := self.NewFrame(method)
    self.PushFrame(newFrame)
    _passArgs(currentFrame.operandStack, newFrame.localVars, method.ActualArgCount())
}
func _passArgs(stack *OperandStack, vars *LocalVars, argCount uint) {
    if argCount > 0 {
        args := stack.popN(argCount)
        for i, j := uint(0), uint(0); i < argCount; i++ {
            arg := args[i]
            args[i] = nil
            vars.Set(i + j, arg)
            if any.IsLongOrDouble(arg) {
                j++
            }
        }
    }
}
func _logInvoke(stackSize uint, method * rtc.Method) {
    space := strings.Repeat(" ", int(stackSize))
    if method.IsStatic() {
        log.Printf("invoke method:%v %v.%v()", space, method.Class().Name(), method.Name())
    } else {
        log.Printf("invoke method:%v %v#%v()", space, method.Class().Name(), method.Name())
    }
}

// args not passed!
func (self *Thread) InvokeMethod2(method * rtc.Method) (*LocalVars) {
    //_logInvoke(self.stack.size, method)
    if !method.IsVoidReturnType() {
        // insert a garbage frame
        garbageMethod := rtc.NewGarbageMethod()
        garbageFrame := self.NewFrame(garbageMethod)
        self.PushFrame(garbageFrame)
    }

    newFrame := self.NewFrame(method)
    self.PushFrame(newFrame)
    return newFrame.localVars
}

// func (self *Thread) PushMethod(method *rtc.Method) (*Frame) {
//     frame := newFrame(self, method)
//     self.PushFrame(frame)
//     return frame
// }

func (self *Thread) HandleUncaughtException(ex *rtc.Obj) {
    self.stack.clear()
    sysClass := ex.Class().ClassLoader().LoadClass("java/lang/System")
    sysErr := sysClass.GetStaticValue("out", "Ljava/io/PrintStream;").(*rtc.Obj)
    printStackTrace := ex.Class().GetMethod("printStackTrace", "(Ljava/io/PrintStream;)V")

    // call ex.printStackTrace(System.err)
    newFrame := self.NewFrame(printStackTrace)
    vars := newFrame.localVars
    vars.SetRef(0, ex)
    vars.SetRef(1, sysErr)
    self.PushFrame(newFrame)
}
