package classfile

import "jvmgo/util"

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
    readInfo(reader *ClassReader)
}

// todo ugly code
func newConstantInfo(tag uint8, cp *ConstantPool) (ConstantInfo) {
    switch tag {
    case CONSTANT_Integer:              return &ConstantIntegerInfo{}
    case CONSTANT_Float:                return &ConstantFloatInfo{}
    case CONSTANT_Long:                 return &ConstantLongInfo{}
    case CONSTANT_Double:               return &ConstantDoubleInfo{}
    case CONSTANT_Utf8:                 return &ConstantUtf8Info{}
    case CONSTANT_String:               return &ConstantStringInfo{cp:cp}
    case CONSTANT_Class:                return &ConstantClassInfo{cp:cp}
    case CONSTANT_Fieldref:             c := &ConstantFieldrefInfo{};           c.cp = cp; return c
    case CONSTANT_Methodref:            c := &ConstantMethodrefInfo{};          c.cp = cp; return c
    case CONSTANT_InterfaceMethodref:   c := &ConstantInterfaceMethodrefInfo{}; c.cp = cp; return c
    case CONSTANT_NameAndType:          return &ConstantNameAndTypeInfo{}
    case CONSTANT_MethodType:           return &ConstantMethodTypeInfo{}
    case CONSTANT_MethodHandle:         return &ConstantMethodHandleInfo{}
    case CONSTANT_InvokeDynamic:        return &ConstantInvokeDynamicInfo{}
    default: // todo
        util.Panicf("BAD constant pool tag: %v", tag)
        return nil
    }
}
