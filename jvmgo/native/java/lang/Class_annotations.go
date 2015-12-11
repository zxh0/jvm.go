package lang

import (
	"github.com/zxh0/jvm.go/jvmgo/jutil"
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	"github.com/zxh0/jvm.go/jvmgo/rtda/heap"
)

func init() {
	_class(getRawAnnotations, "getRawAnnotations", "()[B")
}

// native byte[] getRawAnnotations();
// ()[B
func getRawAnnotations(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()

	class := this.Extra().(*heap.Class)
	goBytes := class.AnnotationData()
	if goBytes != nil {
		jBytes := jutil.CastUint8sToInt8s(goBytes)
		byteArr := heap.NewByteArray(jBytes)
		frame.OperandStack().PushRef(byteArr)
		return
	}

	frame.OperandStack().PushRef(nil)
}

// native byte[] getRawTypeAnnotations();
// ()[B
func getRawTypeAnnotations(frame *rtda.Frame) {
	// todo
}
