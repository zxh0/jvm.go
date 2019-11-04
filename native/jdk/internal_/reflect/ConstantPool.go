package reflect

import (
	"github.com/zxh0/jvm.go/native"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

func init() {
	native.ForClass("jdk/internal/reflect/ConstantPool").
		Register(getLongAt0, "(Ljava/lang/Object;I)J").
		Register(getUTF8At0, "(Ljava/lang/Object;I)Ljava/lang/String;")
}

// private native long getLongAt0(Object o, int i);
// (Ljava/lang/Object;I)J
func getLongAt0(frame *rtda.Frame) {
	class, index := _getArgs(frame)
	val := class.GetConstant(index).(int64)
	frame.PushLong(val)
}

// private native String getUTF8At0(Object o, int i);
// (Ljava/lang/Object;I)Ljava/lang/String;
func getUTF8At0(frame *rtda.Frame) {
	class, index := _getArgs(frame)
	kUtf8 := class.GetConstant(index).(string)
	jStr := frame.GetRuntime().JSFromGoStr(kUtf8)
	frame.PushRef(jStr)
}

func _getArgs(frame *rtda.Frame) (class *heap.Class, index uint) {
	this := frame.GetThis()
	index = uint(frame.GetIntVar(2))
	class = this.Extra.(*heap.Class)
	return
}
