package rtda

import (
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
	"unsafe"
)

// see: jls8 12.4.2. Detailed Initialization Procedure
// http://docs.oracle.com/javase/specs/jls/se8/html/jls-12.html#jls-12.4.2
func initClass(thread *Thread, class *rtc.Class) {
	// step 1
	initCond := class.InitCond()
	initCond.L.Lock()
	defer initCond.L.Unlock()

	// step 2 & 3
	threadPtr := uintptr(unsafe.Pointer(thread))
	isInitializing, initThreadPtr := class.IsBeingInitialized()
	if isInitializing {
		if initThreadPtr != threadPtr {
			initCond.Wait()
		} else {
			return
		}
	}

	// step 4
	if class.IsFullyInitialized() {
		return
	}

	// step 5
	if class.IsInitializationFailed() {
		// todo
		panic("NoClassDefFoundError")
	}

	uninitedClass := getUpmostUninitializedClassOrInterface(class)
	if uninitedClass != nil {
		clinit := uninitedClass.GetClinitMethod()
		if clinit != nil {
			// exec <clinit>
			uninitedClass.MarkBeingInitialized(threadPtr)
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
