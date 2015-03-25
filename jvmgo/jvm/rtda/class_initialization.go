package rtda

import (
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
)

// see: jls8 12.4.2. Detailed Initialization Procedure
// http://docs.oracle.com/javase/specs/jls/se8/html/jls-12.html#jls-12.4.2
func initClass(thread *Thread, class *rtc.Class) {
	uninitedClass := getUpmostUninitializedClassOrInterface(class)
	if uninitedClass != nil {
		clinit := uninitedClass.GetClinitMethod()
		if clinit != nil {
			// exec <clinit>
			uninitedClass.MarkBeingInitialized()
			newFrame := thread.NewFrame(clinit)
			newFrame.SetOnPopAction(func() {
				uninitedClass.MarkFullyInitialized()
			})
			thread.PushFrame(newFrame)
		} else {
			// no <clinit> method
			uninitedClass.MarkFullyInitialized()
		}
	}
}

func getUpmostUninitializedClassOrInterface(from *rtc.Class) *rtc.Class {
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
