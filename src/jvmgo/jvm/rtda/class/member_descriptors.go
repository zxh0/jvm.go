package class

type DescriptorReader struct {
    d       string
    offset  int
}
func (self *DescriptorReader) readUint8() uint8 {
    b := self.d[self.offset]
    self.offset++
    return b
}
func (self *DescriptorReader) unreadUint8() {
    self.offset--
}
func (self *DescriptorReader) startParams() {
    b := self.readUint8()
    if b != '(' {
        self.causePanic()
    }
}
func (self *DescriptorReader) endParams() {
    b := self.readUint8()
    if b != ')' {
        self.causePanic()
    }
}
func (self *DescriptorReader) readFieldType() (bool) {
    b := self.readUint8()
    switch b {
    case 'B', 'C', 'D', 'F', 'I', 'J', 'S', 'Z':
        return true
    case 'L':
        self.readObjectType()
        return true
    case '[':
        self.readArrayType()
        return true
    default:
        self.unreadUint8()
        return false
    }
}
func (self *DescriptorReader) readObjectType() {
    for ';' != self.readUint8() {}
}
func (self *DescriptorReader) readArrayType() {
    self.readFieldType()
}

func (self *DescriptorReader) causePanic() {
    panic("BAD descriptor: " + self.d)
}

// descriptor looks like: (IDLjava/lang/Thread;)Ljava/lang/Object;
func calcArgCount(descriptor string) (uint) {
    dr := &DescriptorReader{descriptor, 0}

    count := 0
    dr.startParams()
    for dr.readFieldType() {
        count++
    }
    dr.endParams()

    return uint(count)
}

// todo
// func IsBaseType(fieldDescriptor string) (bool) {
//     switch fieldDescriptor[0] {
//     case 'B', 'C', 'D', 'F', 'I', 'J', 'S', 'Z':
//         return true
//     default: return false
//     }
// }
