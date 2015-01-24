package instructions

import (
    "log"
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
            nativeMethod := cMethodRef.NativeMethod().(func(*rtda.OperandStack))
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
