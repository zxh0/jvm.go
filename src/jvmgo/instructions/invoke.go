package instructions

import (
    "log"
    //. "jvmgo/any"
    "jvmgo/native"
    "jvmgo/rtda"
    "jvmgo/rtda/class"
)

// Invoke a class (static) method 
type invokestatic struct {Index16Instruction}
func (self *invokestatic) Execute(thread *rtda.Thread) {
    currentFrame := thread.CurrentFrame()
    currentMethod := currentFrame.Method()
    currentClass := currentMethod.Class()
    cp := currentClass.ConstantPool()
    cMethodRef := cp.GetConstant(self.index).(*class.ConstantMethodref)
    method := cMethodRef.Method()

    // init class
    classOfMethod := method.Class()
    if classOfMethod.NotInitialized() {
        if classOfMethod != currentClass || !currentMethod.IsClinit() {
            currentFrame.SetNextPC(thread.PC())
            initClass(classOfMethod, thread)
            return
        }
    }

    if method.IsNative() {
        if method.IsRegisterNatives() {
            // todo
            log.Print("skip registerNatives()!")
        } else {
            // exec native method
            nativeMethod := cMethodRef.NativeMethod().(native.NativeMethod)
            nativeMethod(currentFrame.OperandStack())
        }
        return
    }

    // create new frame
    newFrame := rtda.NewFrame(method)
    thread.PushFrame(newFrame)

    // pass args
    if argCount := method.ArgCount(); argCount > 0 {
        passArgs(currentFrame.OperandStack(), newFrame.LocalVars(), argCount)
    }
}

func passArgs(stack *rtda.OperandStack, vars *rtda.LocalVars, argCount uint) {
    args := stack.PopN(argCount)
    for i := uint(0); i < argCount; i++ {
        arg := args[i]
        vars.Set(i, arg)
        if isLongOrDouble(arg) {
            i++
        }
    }
}

// Invoke instance method;
// special handling for superclass, private, and instance initialization method invocations 
type invokespecial struct {Index16Instruction}
func (self *invokespecial) Execute(thread *rtda.Thread) {
    frame := thread.CurrentFrame()
    stack := frame.OperandStack()

    cp := frame.Method().Class().ConstantPool()
    cMethodRef := cp.GetConstant(self.index).(*class.ConstantMethodref)
    method := cMethodRef.Method()
    newFrame := rtda.NewFrame(method)

    // pass args
    argCount := 1 + method.ArgCount()
    passArgs(stack, newFrame.LocalVars(), argCount)

    thread.PushFrame(newFrame)
}

// Invoke instance method; dispatch based on class
type invokevirtual struct {Index16Instruction}
func (self *invokevirtual) Execute(thread *rtda.Thread) {
    // todo
    panic("todo invokevirtual")
}

// Invoke interface method
type invokeinterface struct {
    index   uint16
    count   uint8
    // 0
}
func (self *invokeinterface) fetchOperands(bcr *BytecodeReader) {
    self.index = bcr.readUint16()
    self.count = bcr.readUint8()
    bcr.readUint8() // must be 0
}
func (self *invokeinterface) Execute(thread *rtda.Thread) {
    // todo
    panic("todo invokeinterface")
}

// Invoke dynamic method
type invokedynamic struct {
    index uint16
    // 0
    // 0
}
func (self *invokedynamic) fetchOperands(bcr *BytecodeReader) {
    self.index = bcr.readUint16()
    bcr.readUint8() // must be 0
    bcr.readUint8() // must be 0
}
func (self *invokedynamic) Execute(thread *rtda.Thread) {
    // todo
    panic("todo invokedynamic")
}
