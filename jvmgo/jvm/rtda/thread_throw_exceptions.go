package rtda

import (
	"fmt"
	. "github.com/zxh0/jvm.go/jvmgo/any"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
)

func (self *Thread) throwException(className, initDesc string, initArgs ...Any) {
	class := rtc.BootLoader().LoadClass(className)
	exObj := class.NewObj()
	athrowFrame := newAthrowFrame(self, exObj, initArgs)
	self.PushFrame(athrowFrame)

	// init exObj
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

func (self *Thread) ThrowDivByZero() {
	self.throwExceptionS("java/lang/ArithmeticException", "/ by zero")
}

func (self *Thread) ThrowNegativeArraySizeException() {
	self.throwExceptionV("java/lang/NegativeArraySizeException")
}

func (self *Thread) ThrowIllegalArgumentException(msg string) {
	self.throwExceptionS("java/lang/IllegalArgumentException", msg)
}

func (self *Thread) ThrowInterruptedException(msg string) {
	self.throwExceptionS("java/lang/InterruptedException", msg)
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
