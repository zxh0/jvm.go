package invoke

import (
	"fmt"

	. "github.com/zxh0/jvm.go/jvmgo/any"
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
)

func init() {
	_mhn(getConstant, "getConstant", "(I)I")
	_mhn(mhn_init, "init", "(Ljava/lang/invoke/MemberName;Ljava/lang/Object;)V")
	_mhn(resolve, "resolve", "(Ljava/lang/invoke/MemberName;Ljava/lang/Class;)Ljava/lang/invoke/MemberName;")
}

func _mhn(method Any, name, desc string) {
	rtc.RegisterNativeMethod("java/lang/invoke/MethodHandleNatives", name, desc, method)
}

// static native int getConstant(int which);
// (I)I
func getConstant(frame *rtda.Frame) {
	vars := frame.LocalVars()
	which := vars.GetInt(0)

	if which == 4 {
		frame.OperandStack().PushInt(1)
	} else {
		frame.OperandStack().PushInt(0)
	}
}

// static native void init(MemberName self, Object ref);
// (Ljava/lang/invoke/MemberName;Ljava/lang/Object;)V
func mhn_init(frame *rtda.Frame) {
	vars := frame.LocalVars()
	mn := vars.GetRef(0)
	ref := vars.GetRef(1)

	if ref.Class().Name() == "java/lang/reflect/Method" {
		classObj := ref.GetFieldValue("clazz", "Ljava/lang/Class;").(*rtc.Obj)
		class := classObj.Extra().(*rtc.Class)
		slot := ref.GetFieldValue("slot", "I").(int32)
		method := class.Methods()[slot]

		mn.SetFieldValue("clazz", "Ljava/lang/Class;", classObj)

		fmt.Printf("method:%v \n", method)
	}

	fmt.Printf("mn:%v  ref:%v \n", mn, ref)
	//panic("todo mhn_init...")
}

// static native MemberName resolve(MemberName self, Class<?> caller) throws LinkageError;
// (Ljava/lang/invoke/MemberName;Ljava/lang/Class;)Ljava/lang/invoke/MemberName;
func resolve(frame *rtda.Frame) {
	vars := frame.LocalVars()
	mn := vars.GetRef(0)
	// caller := vars.GetRef(1)

	// panic("todo resolve")
	stack := frame.OperandStack()
	stack.PushRef(mn)
}
