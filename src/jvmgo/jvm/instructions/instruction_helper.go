package instructions

import (
    "jvmgo/jvm/rtda"
    rtc "jvmgo/jvm/rtda/class"
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
