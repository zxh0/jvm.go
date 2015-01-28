package instructions

import (
    //"log"
    "unicode/utf8"
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

func passArgs(stack *rtda.OperandStack, vars *rtda.LocalVars, argCount uint) {
    args := stack.PopN(argCount)
    for i, j := uint(0), uint(0); i < argCount; i++ {
        arg := args[i]
        vars.Set(i + j, arg)
        if isLongOrDouble(arg) {
            j++
        }
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

func newJString(goStr string, thread *rtda.Thread) {
    currentFrame := thread.CurrentFrame()

    // new string
    stringClass := currentFrame.Method().Class().ClassLoader().StringClass()
    jStr := stringClass.NewObj()
    currentFrame.OperandStack().PushRef(jStr)
    // init string
    codePoints := string2CodePoints(goStr)
    initMethod := stringClass.GetMethod("<init>", "([III)V") //public String(int[] codePoints, int offset, int count)
    newFrame := rtda.NewFrame(initMethod)
    localVars := newFrame.LocalVars()
    localVars.SetRef(0, jStr) // this
    localVars.SetRef(1, rtc.NewIntArray(codePoints))
    localVars.SetInt(2, 0)
    localVars.SetInt(3, int32(len(codePoints)))
    thread.PushFrame(newFrame)
}

func string2CodePoints(str string) ([]rune) {
    runeCount := utf8.RuneCountInString(str)
    codePoints := make([]rune, runeCount)
    i := 0
    for len(str) > 0 {
        r, size := utf8.DecodeRuneInString(str)
        codePoints[i] = r
        i++
        str = str[size:]
    }
    return codePoints
}

