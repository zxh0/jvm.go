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

func (self *Thread) ThrowNegativeArraySizeException() {
	self.ThrowException("java/lang/NegativeArraySizeException", "()V")
}

func (self *Thread) ThrowDivByZero() {
	msgObj := NewJString("/ by zero", self)
	self.ThrowException("java/lang/ArithmeticException", "(Ljava/lang/String;)V", msgObj)
}

func (self *Thread) ThrowArrayIndexOutOfBoundsException(index int32) {
	msg := fmt.Sprintf("%v", index)
	msgObj := NewJString(msg, self)
	self.ThrowException("java/lang/ArrayIndexOutOfBoundsException", "(Ljava/lang/String;)V", msgObj)
}

func (self *Thread) ThrowClassCastException(from, to *rtc.Class) {
	msg := fmt.Sprintf("%v cannot be cast to %v", from.NameJlsFormat(), to.NameJlsFormat())
	msgObj := NewJString(msg, from)
	self.ThrowException("java/lang/ClassCastException", "(Ljava/lang/String;)V", msgObj)
}

func (self *Thread) ThrowClassNotFoundException(name string) {
	msgObj := NewJString(name, self)
	self.ThrowException("java/lang/ClassNotFoundException", "(Ljava/lang/String;)V", msgObj)
}
