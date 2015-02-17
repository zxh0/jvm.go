package rtda

import (
	"fmt"
	rtc "jvmgo/jvm/rtda/class"
)

func (self *Thread) ThrowNPE() {
	self.ThrowException("java/lang/NullPointerException", "()V")
}

func (self *Thread) ThrowClassCastException(from, to *rtc.Class) {
	msg := fmt.Sprintf("%v cannot be cast to %v", from.Name(), to.Name())
	msgObj := NewJString(msg, from)
	self.ThrowException("java/lang/ClassCastException", "(Ljava/lang/String;)V", msgObj)
}
