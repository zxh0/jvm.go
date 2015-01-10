package classfile

import "strconv"

// Constant pool tags
const (
    CONSTANT_Class              =  7
    CONSTANT_Fieldref           =  9
    CONSTANT_Methodref          = 10
    CONSTANT_InterfaceMethodref = 11
    CONSTANT_String             =  8
    CONSTANT_Integer            =  3
    CONSTANT_Float              =  4
    CONSTANT_Long               =  5
    CONSTANT_Double             =  6
    CONSTANT_NameAndType        = 12
    CONSTANT_Utf8               =  1
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
type ConstantInfo interface {
    // empty
}

func readConstantInfo(reader *ClassReader, tag uint8) (ConstantInfo) {
    switch tag {
    case CONSTANT_Integer: return readConstantIntegerInfo(reader)
    case CONSTANT_Float: return readConstantFloatInfo(reader)
    case CONSTANT_Long: return readConstantLongInfo(reader)
    case CONSTANT_Double: return readConstantDoubleInfo(reader)
    case CONSTANT_Utf8: return readConstantUtf8Info(reader)
    case CONSTANT_String: return readConstantStringInfo(reader)
    case CONSTANT_Class: return readConstantClassInfo(reader)
    case CONSTANT_MethodType: return readConstantMethodTypeInfo(reader)
    case CONSTANT_NameAndType: return readConstantNameAndTypeInfo(reader)
    case CONSTANT_Fieldref: return readConstantFieldrefInfo(reader)
    case CONSTANT_Methodref: return readConstantMethodrefInfo(reader)
    case CONSTANT_InterfaceMethodref: return readConstantInterfaceMethodrefInfo(reader)
    case CONSTANT_MethodHandle: return readConstantMethodHandleInfo(reader)
    case CONSTANT_InvokeDynamic: return readConstantInvokeDynamicInfo(reader)
    // todo
    default: panic("Invalid Constant pool tag: " + strconv.Itoa(int(tag)))
    }
}
