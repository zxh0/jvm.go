package rtda

import (
	"fmt"

	"github.com/zxh0/jvm.go/rtda/heap"
)

func (thread *Thread) throwException(className, initDesc string, initArgs ...heap.Slot) {
	class := thread.Runtime.BootLoader().LoadClass(className)
	exObj := class.NewObj()
	athrowFrame := newAthrowFrame(thread, exObj, initArgs)
	thread.PushFrame(athrowFrame)

	// init exObj
	constructor := class.GetConstructor(initDesc)
	thread.InvokeMethod(constructor)
}

func (thread *Thread) throwExceptionV(className string) {
	thread.throwException(className, "()V")
}
func (thread *Thread) throwExceptionS(className, msg string) {
	msgObj := thread.Runtime.JSFromGoStr(msg)
	thread.throwException(className, "(Ljava/lang/String;)V", heap.NewRefSlot(msgObj))
}

func (thread *Thread) ThrowNPE() {
	thread.throwExceptionV("java/lang/NullPointerException")
}

func (thread *Thread) ThrowNegativeArraySizeException() {
	thread.throwExceptionV("java/lang/NegativeArraySizeException")
}

func (thread *Thread) ThrowArrayIndexOutOfBoundsExceptionNoMsg() {
	thread.throwExceptionV("java/lang/ArrayIndexOutOfBoundsException")
}

func (thread *Thread) ThrowDivByZero() {
	thread.throwExceptionS("java/lang/ArithmeticException", "/ by zero")
}

func (thread *Thread) ThrowIllegalArgumentException(msg string) {
	thread.throwExceptionS("java/lang/IllegalArgumentException", msg)
}

func (thread *Thread) ThrowInterruptedException(msg string) {
	thread.throwExceptionS("java/lang/InterruptedException", msg)
}

func (thread *Thread) ThrowClassNotFoundException(name string) {
	thread.throwExceptionS("java/lang/ClassNotFoundException", name)
}

func (thread *Thread) ThrowFileNotFoundException(name string) {
	msg := name + " (No such file or directory)"
	thread.throwExceptionS("java/io/FileNotFoundException", msg)
}

func (thread *Thread) ThrowArrayIndexOutOfBoundsException(index int32) {
	msg := fmt.Sprintf("%v", index)
	thread.throwExceptionS("java/lang/ArrayIndexOutOfBoundsException", msg)
}

func (thread *Thread) ThrowClassCastException(from, to *heap.Class) {
	msg := fmt.Sprintf("%v cannot be cast to %v", from.NameJlsFormat(), to.NameJlsFormat())
	thread.throwExceptionS("java/lang/ClassCastException", msg)
}

func (thread *Thread) ThrowIOException(name string) {
	thread.throwExceptionS("java/lang/IOException", name)
}
