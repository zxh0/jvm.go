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
	native.ForClass("java/lang/invoke/MethodHandleNatives").
		RemovePrefix("mhn_").
		Register(mhn_init, "(Ljava/lang/invoke/MemberName;Ljava/lang/Object;)V").
		Register(resolve, "(Ljava/lang/invoke/MemberName;Ljava/lang/Class;Z)Ljava/lang/invoke/MemberName;").
		Register(getConstant, "(I)I")
}

// static native void init(MemberName self, Object ref);
// (Ljava/lang/invoke/MemberName;Ljava/lang/Object;)V
func mhn_init(frame *rtda.Frame) {
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

// static native MemberName resolve(MemberName self, Class<?> caller, boolean speculativeResolve) throws LinkageError, ClassNotFoundException;
// (Ljava/lang/invoke/MemberName;Ljava/lang/Class;Z)Ljava/lang/invoke/MemberName;
func resolve(frame *rtda.Frame) {
	mnSlot := frame.GetLocalVar(0)
	mnObj := mnSlot.Ref
	// caller := frame.GetRefVar(1)
	// speculativeResolve := frame.GetBooleanVar(2)
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
			if m := cls.GetMethod(nameStr, sigStr); m != nil {
				flags |= int32(m.AccessFlags)
				mnObj.SetFieldValue("flags", "I", heap.NewIntSlot(flags))
			}
		} else {
			panic("TODO: MHN: resolve!" + nameStr + "||" + sigStr)
		}
	})
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
