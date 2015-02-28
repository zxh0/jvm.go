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

	class := this.Extra().(*rtc.Class)
	if attrs := class.Attributes(); attrs != nil {
		if data := attrs.AnnotationData(); data != nil {
			byteArr := rtc.NewByteArray(data, frame.ClassLoader())
			frame.OperandStack().PushRef(byteArr)
			return
		}
	}

	frame.OperandStack().PushRef(nil)
}

// native byte[] getRawTypeAnnotations();
// ()[B
func getRawTypeAnnotations(frame *rtda.Frame) {
	// todo
}
