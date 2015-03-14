package rtda

import (
	"fmt"
	. "jvmgo/any"
	"jvmgo/jvm/options"
	rtc "jvmgo/jvm/rtda/class"
	"strings"
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

func NewThread(jThread *rtc.Obj) *Thread {
	stack := newStack(options.ThreadStackSize)
	return &Thread{0, stack, jThread}
}

// getters & setters
func (self *Thread) PC() int {
	return self.pc
}
func (self *Thread) SetPC(pc int) {
	self.pc = pc
}
func (self *Thread) JThread() *rtc.Obj {
	return self.jThread
}

func (self *Thread) ClassLoader() *rtc.ClassLoader {
	return self.TopFrame().ClassLoader()
}

func (self *Thread) IsStackEmpty() bool {
	return self.stack.isEmpty()
}
func (self *Thread) StackDepth() uint {
	return self.stack.size
}

func (self *Thread) CurrentFrame() *Frame {
	return self.stack.top()
}
func (self *Thread) TopFrame() *Frame {
	return self.stack.top()
}
func (self *Thread) TopFrameN(n uint) *Frame {
	return self.stack.topN(n)
}

func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}
func (self *Thread) PopFrame() *Frame {
	top := self.stack.pop()
	if top.onPopAction != nil {
		// todo
		top.onPopAction()
	}
	return top
}

func (self *Thread) NewFrame(method *rtc.Method) *Frame {
	if method.IsNative() {
		return newNativeFrame(self, method)
	} else {
		return newFrame(self, method)
	}
}

func (self *Thread) InvokeMethod(method *rtc.Method) {
	//self._logInvoke(self.stack.size, method)
	currentFrame := self.CurrentFrame()
	newFrame := self.NewFrame(method)
	self.PushFrame(newFrame)
	actualArgCount := method.ActualArgCount()
	if actualArgCount > 0 {
		_passArgs(currentFrame.operandStack, newFrame.localVars, actualArgCount)
	}

	if method.IsSynchronized() {
		var monitor *rtc.Monitor
		if method.IsStatic() {
			classObj := method.Class().JClass()
			monitor = classObj.Monitor()
		} else {
			thisObj := newFrame.LocalVars().GetThis()
			monitor = thisObj.Monitor()
		}

		monitor.Enter(self)
		newFrame.SetOnPopAction(func() {
			monitor.Exit(self)
		})
	}
}
func _passArgs(stack *OperandStack, vars *LocalVars, argCount uint) {
	args := stack.PopTops(argCount)
	for i, j := uint(0), uint(0); i < argCount; i++ {
		arg := args[i]
		args[i] = nil
		vars.Set(i+j, arg)
		if IsLongOrDouble(arg) {
			j++
		}
	}
}
func (self *Thread) _logInvoke(stackSize uint, method *rtc.Method) {
	space := strings.Repeat(" ", int(stackSize))
	className := method.Class().Name()
	methodName := method.Name()

	if method.IsStatic() {
		fmt.Printf("[method]%v thread:%p %v.%v()\n", space, self, className, methodName)
	} else {
		fmt.Printf("[method]%v thread:%p %v#%v()\n", space, self, className, methodName)
	}
}

func (self *Thread) InvokeMethodWithShim(method *rtc.Method, args []Any) {
	shimFrame := newShimFrame(self, args)
	self.PushFrame(shimFrame)
	self.InvokeMethod(method)
}

func (self *Thread) HandleUncaughtException(ex *rtc.Obj) {
	self.stack.clear()
	sysClass := rtc.BootLoader().LoadClass("java/lang/System")
	sysErr := sysClass.GetStaticValue("out", "Ljava/io/PrintStream;").(*rtc.Obj)
	printStackTrace := ex.Class().GetInstanceMethod("printStackTrace", "(Ljava/io/PrintStream;)V")

	// call ex.printStackTrace(System.err)
	newFrame := self.NewFrame(printStackTrace)
	vars := newFrame.localVars
	vars.SetRef(0, ex)
	vars.SetRef(1, sysErr)
	self.PushFrame(newFrame)

	//
	// printString := sysErr.Class().GetInstanceMethod("print", "(Ljava/lang/String;)V")
	// newFrame = self.NewFrame(printString)
	// vars = newFrame.localVars
	// vars.SetRef(0, sysErr)
	// vars.SetRef(1, NewJString("Exception in thread \"main\" ", newFrame))
	// self.PushFrame(newFrame)
}

// hack
func (self *Thread) HackSetJThread(jThread *rtc.Obj) {
	self.jThread = jThread
}
