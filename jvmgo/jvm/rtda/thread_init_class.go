package rtda

import rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"

func (self *Thread) InitClass(class *rtc.Class) {
	uninitedClass := getUpmostUninitializedClassOrInterface(class)
	if uninitedClass != nil {
		//log.Printf("init: %v", uninitedClass.Name())
		clinit := uninitedClass.GetClinitMethod()
		if clinit != nil {
			// exec <clinit>
			uninitedClass.MarkInitializing()
			newFrame := self.NewFrame(clinit)
			newFrame.SetOnPopAction(func() {
				uninitedClass.MarkInitialized()
			})
			self.PushFrame(newFrame)
		} else {
			// no <clinit> method
			//log.Printf("%v has no <clinit>", uninitedClass.Name())
			uninitedClass.MarkInitialized()
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
