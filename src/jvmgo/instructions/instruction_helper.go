package instructions

import (
    //"log"
    . "jvmgo/any"
    "jvmgo/rtda"
    rtc "jvmgo/rtda/class"
)

func branch(thread *rtda.Thread, offset int) {
    nextPC := thread.PC() + offset
    thread.CurrentFrame().SetNextPC(nextPC)
}

// todo: move to any.go?
func isLongOrDouble(x Any) (bool) {
    switch x.(type) {
    case int64: return true
    case float64: return true
    default: return false
    }
}

// todo
func checkArrIndex(index, len int) {
    if index < 0 || index >= len {
        panic("ArrayIndexOutOfBoundsException")
    }
}

func initClass(class *rtc.Class, thread *rtda.Thread) {
    uninitedClass := rtc.GetUpmostUninitializedClassOrInterface(class)
    if uninitedClass != nil {
        //log.Printf("init: %v", uninitedClass.Name())
        clinit := uninitedClass.GetClinitMethod()
        if clinit != nil {
            // hack!
            if uninitedClass.Name() == "java/lang/Character" {
                uninitedClass.MarkInitialized()
                return
            }

            // exec <clinit>
            newFrame := rtda.NewFrame(clinit)
            newFrame.SetOnPopAction(func() {
                uninitedClass.MarkInitialized()
            })
            thread.PushFrame(newFrame)
        } else {
            // no <clinit> method
            //log.Printf("%v has no <clinit>", uninitedClass.Name())
            uninitedClass.MarkInitialized()
        }
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
