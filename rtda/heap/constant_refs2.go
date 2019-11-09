package heap

import (
	"github.com/zxh0/jvm.go/classfile"
)

type ConstantMethodType struct {
	// todo
}

type ConstantMethodHandle struct {
	referenceKind  uint8
	referenceIndex uint16
}

type ConstantInvokeDynamic struct {
	Name          string
	Type          string
	RefKind       int8
	FieldRef      *ConstantFieldRef
	MethodRef     *ConstantMethodRef
	bootstrapArgs []uint16
}

func newConstantMethodType(mtInfo classfile.ConstantMethodTypeInfo) *ConstantMethodType {
	return &ConstantMethodType{
		// todo
	}
}

func newConstantMethodHandle(mhInfo classfile.ConstantMethodHandleInfo) *ConstantMethodHandle {
	return &ConstantMethodHandle{
		referenceKind:  mhInfo.ReferenceKind,
		referenceIndex: mhInfo.ReferenceIndex,
	}
}

func newConstantInvokeDynamic(cf *classfile.ClassFile,
	indyInfo classfile.ConstantInvokeDynamicInfo) *ConstantInvokeDynamic {

	name, _type := cf.GetNameAndType(indyInfo.NameAndTypeIndex)
	bm := cf.GetBootstrapMethods()[indyInfo.BootstrapMethodAttrIndex]
	mh := cf.GetMethodHandleInfo(bm.BootstrapMethodRef)
	fieldRef, methodRef := getRef(cf, mh.ReferenceKind, mh.ReferenceIndex)

	return &ConstantInvokeDynamic{
		Name:          name,
		Type:          _type,
		FieldRef:      fieldRef,
		MethodRef:     methodRef,
		bootstrapArgs: bm.BootstrapArguments,
	}
}

// https://docs.oracle.com/javase/specs/jvms/se13/html/jvms-4.html#jvms-4.4.8
func getRef(cf *classfile.ClassFile, refKind uint8, refIdx uint16) (*ConstantFieldRef, *ConstantMethodRef) {
	switch refKind {
	case classfile.RefGetField, classfile.RefGetStatic,
		classfile.RefPutField, classfile.RefPutStatic:
		return newConstantFieldRef(cf, cf.GetFieldRefInfo(refIdx)), nil
	default: // TODO
		return nil, newConstantMethodRef(cf, cf.GetMethodRefInfo(refIdx))
	}
}

//func (indy *ConstantInvokeDynamic) MethodHandle() {
//	kMH := indy.cp.GetConstant(uint(indy.bootstrapMethodRef)).(*ConstantMethodHandle)
//	println(kMH)
//}
