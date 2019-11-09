package invoke

import (
	"github.com/zxh0/jvm.go/classfile"
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

/*
	/// MemberName support

    static native void init(MemberName self, Object ref);
    static native void expand(MemberName self);
    static native MemberName resolve(MemberName self, Class<?> caller,
            boolean speculativeResolve) throws LinkageError, ClassNotFoundException;
    static native int getMembers(Class<?> defc, String matchName, String matchSig,
            int matchFlags, Class<?> caller, int skip, MemberName[] results);

    /// Field layout queries parallel to jdk.internal.misc.Unsafe:
    static native long objectFieldOffset(MemberName self);  // e.g., returns vmindex
    static native long staticFieldOffset(MemberName self);  // e.g., returns vmindex
    static native Object staticFieldBase(MemberName self);  // e.g., returns clazz
    static native Object getMemberVMInfo(MemberName self);  // returns {vmindex,vmtarget}

    /// CallSite support

	// Tell the JVM that we need to change the target of a CallSite.
	static native void setCallSiteTargetNormal(CallSite site, MethodHandle target);
	static native void setCallSiteTargetVolatile(CallSite site, MethodHandle target);

	static native void copyOutBootstrapArguments(Class<?> caller, int[] indexInfo,
												int start, int end,
												Object[] buf, int pos,
												boolean resolve,
												Object ifNotAvailable);
*/
func init() {
	native.ForClass("java/lang/invoke/MethodHandleNatives").
		RemovePrefix("mhn_").
		Register(mhn_init, "(Ljava/lang/invoke/MemberName;Ljava/lang/Object;)V").
		Register(resolve, "(Ljava/lang/invoke/MemberName;Ljava/lang/Class;Z)Ljava/lang/invoke/MemberName;").
		Register(objectFieldOffset, "(Ljava/lang/invoke/MemberName;)J").
		Register(staticFieldOffset, "(Ljava/lang/invoke/MemberName;)J").
		Register(staticFieldBase, "(Ljava/lang/invoke/MemberName;)Ljava/lang/Object;")
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
		flags |= MN_IS_METHOD | (classfile.RefInvokeStatic << MN_REFERENCE_KIND_SHIFT)
	} else if method.IsConstructor() {
		flags |= MN_IS_CONSTRUCTOR | (classfile.RefInvokeSpecial << MN_REFERENCE_KIND_SHIFT)
	} else {
		flags |= MN_IS_METHOD | (classfile.RefInvokeSpecial << MN_REFERENCE_KIND_SHIFT)
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
			if f := cls.GetField(nameStr, sigStr); f != nil {
				flags |= int32(f.AccessFlags)
				mnObj.SetFieldValue("flags", "I", heap.NewIntSlot(flags))
			}
		}
	})
}

// static native Object staticFieldBase(MemberName self);
func staticFieldBase(frame *rtda.Frame) {
	mName := frame.GetRefVar(0)
	class := mName.GetFieldValue("clazz", "*").Ref.GetGoClass()
	name := mName.GetFieldValue("name", "*").Ref.JSToGoStr()
	field := class.GetField(name, "*") // TODO: check static
	frame.PushRef(field.Class.JClass)
}

// static native long staticFieldOffset(MemberName self);
func staticFieldOffset(frame *rtda.Frame) {
	mName := frame.GetRefVar(0)
	class := mName.GetFieldValue("clazz", "*").Ref.GetGoClass()
	name := mName.GetFieldValue("name", "*").Ref.JSToGoStr()
	field := class.GetField(name, "*") // TODO: check static
	frame.PushLong(int64(field.SlotId))
}

// static native long objectFieldOffset(MemberName self);
func objectFieldOffset(frame *rtda.Frame) {
	mName := frame.GetRefVar(0)
	class := mName.GetFieldValue("clazz", "*").Ref.GetGoClass()
	name := mName.GetFieldValue("name", "*").Ref.JSToGoStr()
	field := class.GetField(name, "*") // TODO: check non-static
	frame.PushLong(int64(field.SlotId))
}
