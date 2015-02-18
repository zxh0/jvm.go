package rtda

import (
	"fmt"
	. "jvmgo/any"
	rtc "jvmgo/jvm/rtda/class"
)

func (self *Thread) ThrowException(className, initDesc string, initArgs ...Any) {
	class := self.ClassLoader().LoadClass(className)
	ex := class.NewObj()
	athrowFrame := newAthrowFrame(self, ex, initArgs)
	self.PushFrame(athrowFrame)

	// init ex
	constructor := class.GetConstructor(initDesc)
	self.InvokeMethod(constructor)
}

func (self *Thread) ThrowNPE() {
	self.ThrowException("java/lang/NullPointerException", "()V")
}

func (self *Thread) ThrowClassCastException(from, to *rtc.Class) {
	msg := fmt.Sprintf("%v cannot be cast to %v", from.JlsName(), to.JlsName())
	msgObj := NewJString(msg, from)
	self.ThrowException("java/lang/ClassCastException", "(Ljava/lang/String;)V", msgObj)
}
