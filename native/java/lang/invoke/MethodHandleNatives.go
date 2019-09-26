package invoke

import (
	"fmt"

	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

func init() {
	_mhn(getConstant, "getConstant", "(I)I")
	_mhn(mhn_init, "init", "(Ljava/lang/invoke/MemberName;Ljava/lang/Object;)V")
	_mhn(resolve, "resolve", "(Ljava/lang/invoke/MemberName;Ljava/lang/Class;)Ljava/lang/invoke/MemberName;")
}

func _mhn(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("java/lang/invoke/MethodHandleNatives", name, desc, method)
}

// static native int getConstant(int which);
// (I)I
func getConstant(frame *rtda.Frame) {
	which := frame.GetIntVar(0)

	if which == 4 {
		frame.PushInt(1)
	} else {
		frame.PushInt(0)
	}
}

// static native void init(MemberName self, Object ref);
// (Ljava/lang/invoke/MemberName;Ljava/lang/Object;)V
func mhn_init(frame *rtda.Frame) {
	mn := frame.GetRefVar(0)
	ref := frame.GetRefVar(1)

	if ref.Class().Name() == "java/lang/reflect/Method" {
		classObj := ref.GetFieldValue("clazz", "Ljava/lang/Class;").Ref
		class := classObj.Extra().(*heap.Class)
		slot := ref.GetFieldValue("slot", "I").IntValue()
		method := class.Methods()[slot]

		mn.SetFieldValue("clazz", "Ljava/lang/Class;", heap.NewRefSlot(classObj))

		fmt.Printf("method:%v \n", method)
	}

	fmt.Printf("mn:%v  ref:%v \n", mn, ref)
	//panic("todo mhn_init...")
}

// static native MemberName resolve(MemberName self, Class<?> caller) throws LinkageError;
// (Ljava/lang/invoke/MemberName;Ljava/lang/Class;)Ljava/lang/invoke/MemberName;
func resolve(frame *rtda.Frame) {
	mn := frame.GetRefVar(0)
	// caller := frame.GetRefVar(1)

	// panic("todo resolve")
	frame.PushRef(mn)
}
