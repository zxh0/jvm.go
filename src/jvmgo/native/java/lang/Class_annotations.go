package lang

import (
	"jvmgo/jvm/rtda"
	rtc "jvmgo/jvm/rtda/class"
)

func init() {
	_class(getRawAnnotations, "getRawAnnotations", "()[B")
}

// native byte[] getRawAnnotations();
// ()[B
func getRawAnnotations(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()

	class := this.Class()
	bytes := class.AnnotationData()

	stack := frame.OperandStack()
	if bytes != nil {
		byteArr := rtc.NewByteArray(bytes, frame.ClassLoader())
		stack.PushRef(byteArr)
	} else {
		stack.PushRef(nil)
	}
}

// native byte[] getRawTypeAnnotations();
// ()[B
func getRawTypeAnnotations(frame *rtda.Frame) {
	// todo
}
