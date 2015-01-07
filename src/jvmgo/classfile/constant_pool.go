package classfile

import (
    "fmt"
    "strconv"
)

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

type ConstantPool struct {
    //cpCount uint
    cpInfos []ConstantInfo
}

/*
cp_info {
    u1 tag;
    u1 info[];
}
*/
type ConstantInfo interface {
    // empty
}

func readConstantPool(reader *ClassReader) (*ConstantPool) {
    cpCount := reader.readUint16()
    cpInfos := make([]ConstantInfo, cpCount)

    // The constant_pool table is indexed from 1 to constant_pool_count - 1. 
    for i := uint16(1); i < cpCount; i++ {
        tag := reader.readUint8()
        cpInfos[i] = readConstantInfo(reader, tag)
        // http://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.4.5
        // All 8-byte constants take up two entries in the constant_pool table of the class file.
        // If a CONSTANT_Long_info or CONSTANT_Double_info structure is the item in the constant_pool
        // table at index n, then the next usable item in the pool is located at index n+2. 
        // The constant_pool index n+1 must be valid but is considered unusable. 
        if tag == CONSTANT_Long || tag == CONSTANT_Double {
            i++;
        }
    }

    return &ConstantPool{cpInfos}
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

// todo
func (self *ConstantPool) getUtf8(index uint16) (string) {
    cpInfo := self.cpInfos[index]
    if utf8Info, ok := cpInfo.(*ConstantUtf8Info); ok {
        return utf8Info.str
    } 
    
    // todo
    panic(fmt.Sprintf("Const#%v is not ConstantUtf8Info!", index))
}
