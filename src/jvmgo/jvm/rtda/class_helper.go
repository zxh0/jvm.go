package rtda

import rtc "jvmgo/jvm/rtda/class"

// todo move to jvmgo/jvm/rtda/class_helper.go
func InitClass(class *rtc.Class, thread *Thread) {
    uninitedClass := getUpmostUninitializedClassOrInterface(class)
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

func getUpmostUninitializedClassOrInterface(from *rtc.Class) (*rtc.Class) {
    if !from.InitializationNotStarted() {
        return nil
    }
    for k := from.SuperClass(); k != nil; k = k.SuperClass() {
        if k.InitializationNotStarted() {
            return getUpmostUninitializedClassOrInterface(k)
        }
    }
    for _, i := range from.Interfaces() {
        if i.InitializationNotStarted() {
            return getUpmostUninitializedClassOrInterface(i)
        }
    }
    return from
}
