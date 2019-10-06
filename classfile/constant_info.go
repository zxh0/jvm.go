package classfile

import (
	"fmt"
)

// Constant pool tags
const (
	ConstantUtf8               = 1  // Java 1.0.2
	ConstantInteger            = 3  // Java 1.0.2
	ConstantFloat              = 4  // Java 1.0.2
	ConstantLong               = 5  // Java 1.0.2
	ConstantDouble             = 6  // Java 1.0.2
	ConstantClass              = 7  // Java 1.0.2
	ConstantString             = 8  // Java 1.0.2
	ConstantFieldRef           = 9  // Java 1.0.2
	ConstantMethodRef          = 10 // Java 1.0.2
	ConstantInterfaceMethodRef = 11 // Java 1.0.2
	ConstantNameAndType        = 12 // Java 1.0.2
	ConstantMethodHandle       = 15 // Java 7
	ConstantMethodType         = 16 // Java 7
	ConstantInvokeDynamic      = 18 // Java 7
	ConstantModule             = 19 // Java 9
	ConstantPackage            = 20 // Java 9
	ConstantDynamic            = 17 // Java 11
)

/*
cp_info {
    u1 tag;
    u1 info[];
}
*/
type ConstantInfo interface{}

func readConstantInfo(reader *ClassReader) ConstantInfo {
	tag := reader.ReadUint8()
	switch tag {
	case ConstantInteger:
		return readConstantIntegerInfo(reader)
	case ConstantFloat:
		return readConstantFloatInfo(reader)
	case ConstantLong:
		return readConstantLongInfo(reader)
	case ConstantDouble:
		return readConstantDoubleInfo(reader)
	case ConstantUtf8:
		return readConstantUtf8Info(reader)
	case ConstantString:
		return readConstantStringInfo(reader)
	case ConstantClass:
		return readConstantClassInfo(reader)
	case ConstantModule:
		return readConstantModuleInfo(reader)
	case ConstantPackage:
		return readConstantPackageInfo(reader)
	case ConstantFieldRef:
		return readConstantFieldRefInfo(reader)
	case ConstantMethodRef:
		return readConstantMethodRefInfo(reader)
	case ConstantInterfaceMethodRef:
		return readConstantInterfaceMethodRefInfo(reader)
	case ConstantNameAndType:
		return readConstantNameAndTypeInfo(reader)
	case ConstantMethodType:
		return readConstantMethodTypeInfo(reader)
	case ConstantMethodHandle:
		return readConstantMethodHandleInfo(reader)
	case ConstantInvokeDynamic:
		return readConstantInvokeDynamicInfo(reader)
	default: // TODO
		panic(fmt.Errorf("invalid constant pool tag: %d", tag))
	}
}
