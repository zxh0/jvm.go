package rtda

import (
	"fmt"
	. "jvmgo/any"
	rtc "jvmgo/jvm/rtda/class"
)

func (self *Thread) throwException(className, initDesc string, initArgs ...Any) {
	class := rtc.BootLoader().LoadClass(className)
	ex := class.NewObj()
	athrowFrame := newAthrowFrame(self, ex, initArgs)
	self.PushFrame(athrowFrame)

	// init ex
	constructor := class.GetConstructor(initDesc)
	self.InvokeMethod(constructor)
}

func (self *Thread) throwExceptionV(className string) {
	self.throwException(className, "()V")
}
func (self *Thread) throwExceptionS(className, msg string) {
	msgObj := NewJString(msg)
	self.throwException(className, "(Ljava/lang/String;)V", msgObj)
}

func (self *Thread) ThrowNPE() {
	self.throwExceptionV("java/lang/NullPointerException")
}

func (self *Thread) ThrowNegativeArraySizeException() {
	self.throwExceptionV("java/lang/NegativeArraySizeException")
}

func (self *Thread) ThrowDivByZero() {
	self.throwExceptionS("java/lang/ArithmeticException", "/ by zero")
}

func (self *Thread) ThrowClassNotFoundException(name string) {
	self.throwExceptionS("java/lang/ClassNotFoundException", name)
}

func (self *Thread) ThrowFileNotFoundException(name string) {
	msg := name + " (No such file or directory)"
	self.throwExceptionS("java/io/FileNotFoundException", msg)
}

func (self *Thread) ThrowArrayIndexOutOfBoundsException(index int32) {
	msg := fmt.Sprintf("%v", index)
	self.throwExceptionS("java/lang/ArrayIndexOutOfBoundsException", msg)
}

func (self *Thread) ThrowClassCastException(from, to *rtc.Class) {
	msg := fmt.Sprintf("%v cannot be cast to %v", from.NameJlsFormat(), to.NameJlsFormat())
	self.throwExceptionS("java/lang/ClassCastException", msg)
}
