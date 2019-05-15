package classfile

import (
	"github.com/zxh0/jvm.go/jvmgo/jutil"
)

// Constant pool tags
const (
	CONSTANT_Class              = 7
	CONSTANT_Fieldref           = 9
	CONSTANT_Methodref          = 10
	CONSTANT_InterfaceMethodref = 11
	CONSTANT_String             = 8
	CONSTANT_Integer            = 3
	CONSTANT_Float              = 4
	CONSTANT_Long               = 5
	CONSTANT_Double             = 6
	CONSTANT_NameAndType        = 12
	CONSTANT_Utf8               = 1
	CONSTANT_MethodHandle       = 15
	CONSTANT_MethodType         = 16
	CONSTANT_InvokeDynamic      = 18
)

/*
cp_info {
    u1 tag;
    u1 info[];
}
*/
type ConstantInfo interface{}

func readConstantInfo(reader *ClassReader, cp *ConstantPool) ConstantInfo {
	tag := reader.readUint8()
	switch tag {
	case CONSTANT_Integer:
		return readConstantIntegerInfo(reader)
	case CONSTANT_Float:
		return readConstantFloatInfo(reader)
	case CONSTANT_Long:
		return readConstantLongInfo(reader)
	case CONSTANT_Double:
		return readConstantDoubleInfo(reader)
	case CONSTANT_Utf8:
		return readConstantUtf8Info(reader)
	case CONSTANT_String:
		return readConstantStringInfo(reader, cp)
	case CONSTANT_Class:
		return readConstantClassInfo(reader, cp)
	case CONSTANT_Fieldref,
		CONSTANT_Methodref,
		CONSTANT_InterfaceMethodref:
		return readConstantMemberrefInfo(reader, cp, tag)
	case CONSTANT_NameAndType:
		return readConstantNameAndTypeInfo(reader)
	case CONSTANT_MethodType:
		return readConstantMethodTypeInfo(reader)
	case CONSTANT_MethodHandle:
		return readConstantMethodHandleInfo(reader)
	case CONSTANT_InvokeDynamic:
		return readConstantInvokeDynamicInfo(reader, cp)
	default: // todo
		jutil.Panicf("BAD constant pool tag: %v", tag)
		return nil
	}
}
