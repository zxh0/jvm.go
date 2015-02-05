package rtda

import rtc "jvmgo/jvm/rtda/class"

// todo move to jvmgo/jvm/rtda/class_helper.go
func InitClass(class *rtc.Class, thread *Thread) {
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
