package instructions

import (
    //"log"
    "unicode/utf8"
    //. "jvmgo/any"
    "jvmgo/rtda"
    rtc "jvmgo/rtda/class"
)

func branch(frame *rtda.Frame, offset int) {
    nextPC := frame.Thread().PC() + offset
    frame.SetNextPC(nextPC)
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
            uninitedClass.MarkInitializing()
            newFrame := thread.NewFrame(clinit)
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

func newJString(goStr string, thread *rtda.Thread) (*rtc.Obj) {
    currentFrame := thread.CurrentFrame()

    // new string
    stringClass := currentFrame.Method().Class().ClassLoader().StringClass()
    jStr := stringClass.NewObj()
    //currentFrame.OperandStack().PushRef(jStr)
    // init string
    codePoints := string2CodePoints(goStr)
    initMethod := stringClass.GetMethod("<init>", "([III)V") //public String(int[] codePoints, int offset, int count)
    newFrame := thread.NewFrame(initMethod)
    localVars := newFrame.LocalVars()
    localVars.SetRef(0, jStr) // this
    localVars.SetRef(1, rtc.NewIntArray(codePoints))
    localVars.SetInt(2, 0)
    localVars.SetInt(3, int32(len(codePoints)))
    thread.PushFrame(newFrame)

    return jStr
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

