package invoke

import (
	"fmt"

	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

func init() {
	_mhn(getConstant, "getConstant", "(I)I")
	_mhn(mhnInit, "init", "(Ljava/lang/invoke/MemberName;Ljava/lang/Object;)V")
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
func mhnInit(frame *rtda.Frame) {
	mn := frame.GetRefVar(0)
	ref := frame.GetRefVar(1)

	if ref.Class.Name == "java/lang/reflect/Method" {
		classObj := ref.GetFieldValue("clazz", "Ljava/lang/Class;").Ref
		class := classObj.GetGoClass()
		slot := ref.GetFieldValue("slot", "I").IntValue()
		method := class.Methods[slot]

		mn.SetFieldValue("clazz", "Ljava/lang/Class;", heap.NewRefSlot(classObj))

		fmt.Printf("mhnInit! method:%v \n", method)
	}

	fmt.Printf("mn:%v  ref:%v \n", mn, ref)
	//panic("TODO: mhnInit!")
}

// static native MemberName resolve(MemberName self, Class<?> caller) throws LinkageError;
// (Ljava/lang/invoke/MemberName;Ljava/lang/Class;)Ljava/lang/invoke/MemberName;
func resolve(frame *rtda.Frame) {
	mnSlot := frame.GetLocalVar(0)
	mnObj := mnSlot.Ref
	// caller := frame.GetRefVar(1)
	// panic("TODO: resolve!")
	frame.PushRef(mnObj)

	clsObj := mnObj.GetFieldValue("clazz", "Ljava/lang/Class;").Ref
	nameObj := mnObj.GetFieldValue("name", "Ljava/lang/String;").Ref
	flags := mnObj.GetFieldValue("flags", "I").IntValue()
	getSig := mnObj.Class.GetInstanceMethod("getSignature", "()Ljava/lang/String;")

	cls := clsObj.GetGoClass()
	nameStr := heap.JSToGoStr(nameObj)

	frame.Thread.InvokeMethodWithShim(getSig, []heap.Slot{mnSlot})
	frame.Thread.CurrentFrame().AppendOnPopAction(func(shim *rtda.Frame) {
		sigObj := shim.TopRef(0)
		sigStr := heap.JSToGoStr(sigObj)
		if sigStr[0] == '(' {
			if m := getMethod(cls, nameStr, sigStr); m != nil {
				flags |= int32(m.AccessFlags)
				mnObj.SetFieldValue("flags", "I", heap.NewIntSlot(flags))
			}
		} else {
			panic("TODO")
		}
	})
}

// TODO
func getMethod(cls *heap.Class, name, descriptor string) *heap.Method {
	if m := cls.GetStaticMethod(name, descriptor); m != nil {
		return m
	}
	if m := cls.GetInstanceMethod(name, descriptor); m != nil {
		return m
	}
	return nil
}
