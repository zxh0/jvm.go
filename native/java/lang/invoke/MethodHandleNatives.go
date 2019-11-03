package invoke

import (
	"github.com/zxh0/jvm.go/instructions/references"
	"github.com/zxh0/jvm.go/native"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

const (
	MN_IS_METHOD            = 0x00010000 // method (not constructor)
	MN_IS_CONSTRUCTOR       = 0x00020000 // constructor
	MN_IS_FIELD             = 0x00040000 // field
	MN_IS_TYPE              = 0x00080000 // nested type
	MN_CALLER_SENSITIVE     = 0x00100000 // @CallerSensitive annotation detected
	MN_REFERENCE_KIND_SHIFT = 24         // refKind
	MN_REFERENCE_KIND_MASK  = 0x0F000000 >> MN_REFERENCE_KIND_SHIFT
)

func init() {
	_mhn(mhnInit, "init", "(Ljava/lang/invoke/MemberName;Ljava/lang/Object;)V")
	_mhn(resolve, "resolve", "(Ljava/lang/invoke/MemberName;Ljava/lang/Class;)Ljava/lang/invoke/MemberName;")
	_mhn(getConstant, "getConstant", "(I)I")
}

func _mhn(method native.Method, name, desc string) {
	native.Register("java/lang/invoke/MethodHandleNatives", name, desc, method)
}

// static native void init(MemberName self, Object ref);
// (Ljava/lang/invoke/MemberName;Ljava/lang/Object;)V
func mhnInit(frame *rtda.Frame) {
	mn := frame.GetRefVar(0)
	ref := frame.GetRefVar(1)
	//fmt.Printf("mn:%v  ref:%v \n", mn, ref)

	if ref.Class.Name == "java/lang/reflect/Method" {
		classObj := ref.GetFieldValue("clazz", "Ljava/lang/Class;").Ref
		class := classObj.GetGoClass()
		slot := ref.GetFieldValue("slot", "I").IntValue()
		method := class.Methods[slot]

		mn.SetFieldValue("clazz", "Ljava/lang/Class;", heap.NewRefSlot(classObj))
		mn.SetFieldValue("flags", "I", heap.NewIntSlot(getMNFlags(method)))
	} else {
		panic("TODO: mhnInit! " + ref.Class.Name)
	}
}

func getMNFlags(method *heap.Method) int32 {
	flags := int32(method.AccessFlags)
	if method.IsStatic() {
		flags |= MN_IS_METHOD | (references.RefInvokeStatic << MN_REFERENCE_KIND_SHIFT)
	} else if method.IsConstructor() {
		flags |= MN_IS_CONSTRUCTOR | (references.RefInvokeSpecial << MN_REFERENCE_KIND_SHIFT)
	} else {
		flags |= MN_IS_METHOD | (references.RefInvokeSpecial << MN_REFERENCE_KIND_SHIFT)
	}
	return flags
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
	nameStr := nameObj.JSToGoStr()

	frame.Thread.InvokeMethodWithShim(getSig, []heap.Slot{mnSlot})
	frame.Thread.CurrentFrame().AppendOnPopAction(func(shim *rtda.Frame) {
		sigObj := shim.TopRef(0)
		sigStr := sigObj.JSToGoStr()
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
