package instructions

import (
    //"log"
    "jvmgo/rtda"
    "jvmgo/rtda/class"
)

// Invoke a class (static) method 
type invokestatic struct {Index16Instruction}
func (self *invokestatic) Execute(frame *rtda.Frame) {
    thread := frame.Thread()
    currentFrame := frame
    currentMethod := currentFrame.Method()
    currentClass := currentMethod.Class()
    cp := currentClass.ConstantPool()
    cMethodRef := cp.GetConstant(self.index).(*class.ConstantMethodref)
    method := cMethodRef.StaticMethod()

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
        // exec native method
        nativeMethod := method.NativeMethod().(func(*rtda.Frame))
        nativeMethod(currentFrame)
        return
    }

    // create new frame
    newFrame := thread.NewFrame(method)
    thread.PushFrame(newFrame)

    // pass args
    if argCount := method.ArgCount(); argCount > 0 {
        passArgs(currentFrame.OperandStack(), newFrame.LocalVars(), argCount)
    }
}
