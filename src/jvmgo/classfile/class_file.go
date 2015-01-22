package classfile

//import "errors"

/*
ClassFile {
    u4             magic;
    u2             minor_version;
    u2             major_version;
    u2             constant_pool_count;
    cp_info        constant_pool[constant_pool_count-1];
    u2             access_flags;
    u2             this_class;
    u2             super_class;
    u2             interfaces_count;
    u2             interfaces[interfaces_count];
    u2             fields_count;
    field_info     fields[fields_count];
    u2             methods_count;
    method_info    methods[methods_count];
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type ClassFile struct {
  //magic           uint32
    minorVersion    uint16
    majorVersion    uint16
    constantPool    *ConstantPool
    accessFlags     uint16
    thisClass       uint16
    superClass      uint16
    interfaces      []uint16
    fields          []*FieldInfo
    methods         []*MethodInfo
    attributes      []AttributeInfo
}

func (self *ClassFile) readClass(reader *ClassReader) {
    self.readAndCheckMagic(reader)
    self.readVersions(reader)
    
    self.accessFlags = reader.readUint16()
    self.thisClass = reader.readUint16()
    self.superClass = reader.readUint16()
    self.interfaces = readInterfaces(reader)
    self.fields = readFields(reader, self.constantPool)
    self.methods = readMethods(reader, self.constantPool)
    self.attributes = readAttributes(reader, self.constantPool)
}

func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
    magic := reader.readUint32()
    if magic != 0xCAFEBABE {
        panic("Bad magic!")
    }
}

func (self *ClassFile) readVersions(reader *ClassReader) {
    self.minorVersion = reader.readUint16()
    self.majorVersion = reader.readUint16()
}

func (self *ClassFile) readConstantPool(reader *ClassReader) {
    self.constantPool = &ConstantPool{}
    self.constantPool.read(reader)
}

func readInterfaces(reader *ClassReader) ([]uint16) {
    interfacesCount := reader.readUint16()
    interfaces := make([]uint16, interfacesCount)
    for i := uint16(0); i < interfacesCount; i++ {
        interfaces[i] = reader.readUint16()
    }
    return interfaces
}

func readFields(reader *ClassReader, cp *ConstantPool) ([]*FieldInfo) {
    fieldsCount := reader.readUint16()
    fields := make([]*FieldInfo, fieldsCount)
    for i := uint16(0); i < fieldsCount; i++ {
        fields[i] = readFieldInfo(reader, cp)
    }
    return fields
}

func readMethods(reader *ClassReader, cp *ConstantPool) ([]*MethodInfo) {
    methodsCount := reader.readUint16()
    methods := make([]*MethodInfo, methodsCount)
    for i := uint16(0); i < methodsCount; i++ {
        methods[i] = readMethodInfo(reader, cp)
    }
    return methods
}

func readAttributes(reader *ClassReader, cp *ConstantPool) ([]AttributeInfo) {
    attributesCount := reader.readUint16()
    attributes := make([]AttributeInfo, attributesCount)
    for i := uint16(0); i < attributesCount; i++ {
        attributes[i] = readAttribute(reader, cp)
    }
    return attributes
}
