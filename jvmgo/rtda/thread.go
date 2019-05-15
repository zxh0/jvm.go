package rtda

import (
	"fmt"
	"strings"
	"sync"

	"github.com/zxh0/jvm.go/jvmgo/options"
	"github.com/zxh0/jvm.go/jvmgo/rtda/heap"
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
	pc              int // the address of the instruction currently being executed
	stack           *Stack
	frameCache      *FrameCache
	jThread         *heap.Object // java.lang.Thread
	lock            *sync.Mutex  // state lock
	ch              chan int
	sleepingFlag    bool
	interruptedFlag bool
	parkingFlag     bool // used by Unsafe
	unparkedFlag    bool // used by Unsafe
	// todo
}

func NewThread(jThread *heap.Object) *Thread {
	stack := newStack(options.ThreadStackSize)
	thread := &Thread{
		stack:   stack,
		jThread: jThread,
		lock:    &sync.Mutex{},
		ch:      make(chan int),
	}
	thread.frameCache = newFrameCache(thread, 16) // todo
	return thread
}

// getters & setters
func (self *Thread) PC() int {
	return self.pc
}
func (self *Thread) SetPC(pc int) {
	self.pc = pc
}
func (self *Thread) JThread() *heap.Object {
	return self.jThread
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

	self.frameCache.returnFrame(top)
	return top
}

func (self *Thread) NewFrame(method *heap.Method) *Frame {
	if method.IsNative() {
		return newNativeFrame(self, method)
	} else {
		return self.frameCache.borrowFrame(method)
		//return newFrame(self, method)
	}
}

func (self *Thread) InvokeMethod(method *heap.Method) {
	//self._logInvoke(self.stack.size, method)
	currentFrame := self.CurrentFrame()
	newFrame := self.NewFrame(method)
	self.PushFrame(newFrame)
	argSlotsCount := method.ArgSlotCount()
	if argSlotsCount > 0 {
		_passArgs(currentFrame.operandStack, newFrame.localVars, argSlotsCount)
	}

	if method.IsSynchronized() {
		var monitor *heap.Monitor
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
func _passArgs(stack *OperandStack, vars *LocalVars, argSlotsCount uint) {
	args := stack.PopTops(argSlotsCount)
	for i := uint(0); i < argSlotsCount; i++ {
		arg := args[i]
		args[i] = nil
		vars.Set(i, arg)
	}
}
func (self *Thread) _logInvoke(stackSize uint, method *heap.Method) {
	space := strings.Repeat(" ", int(stackSize))
	className := method.Class().Name()
	methodName := method.Name()

	if method.IsStatic() {
		fmt.Printf("[method]%v thread:%p %v.%v()\n", space, self, className, methodName)
	} else {
		fmt.Printf("[method]%v thread:%p %v#%v()\n", space, self, className, methodName)
	}
}

func (self *Thread) InvokeMethodWithShim(method *heap.Method, args []interface{}) {
	shimFrame := newShimFrame(self, args)
	self.PushFrame(shimFrame)
	self.InvokeMethod(method)
}

func (self *Thread) InitClass(class *heap.Class) {
	initClass(self, class)
}

func (self *Thread) HandleUncaughtException(ex *heap.Object) {
	self.stack.clear()
	sysClass := heap.BootLoader().LoadClass("java/lang/System")
	sysErr := sysClass.GetStaticValue("out", "Ljava/io/PrintStream;").(*heap.Object)
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
	// vars.SetRef(1, JString("Exception in thread \"main\" ", newFrame))
	// self.PushFrame(newFrame)
}

// hack
func (self *Thread) HackSetJThread(jThread *heap.Object) {
	self.jThread = jThread
}
