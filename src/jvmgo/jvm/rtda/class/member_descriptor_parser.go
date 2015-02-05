package class

import (
    "strings"
)

type MemberDescriptorParser struct {
    descriptor  string
    offset      int
}

func (self *MemberDescriptorParser) readUint8() uint8 {
    b := self.descriptor[self.offset]
    self.offset++
    return b
}
func (self *MemberDescriptorParser) unreadUint8() {
    self.offset--
}

// todo
func calcArgCount(descriptor string) (uint) {
    return 0
}
func isVoidReturnType(descriptor string) bool {
    return false
}

func parseMethodDescriptor(descriptor string) (*MethodDescriptor) {
    parser := MemberDescriptorParser{descriptor, 0}
    md := &MethodDescriptor{}

    // parse parameter types
    parser.startParams()
    for {
        t := parser.readFieldType()
        if t != nil {
            md.addParameterType(t)
        } else {
            break
        }
    }
    parser.endParams()

    // parse return type
    t := parser.readFieldType()
    if t != nil {
        md.returnType = t
    } else {
        parser.causePanic()
    }

    parser.finish()
    return md
}

func (self *MemberDescriptorParser) startParams() {
    b := self.readUint8()
    if b != '(' {
        self.causePanic()
    }
}
func (self *MemberDescriptorParser) endParams() {
    b := self.readUint8()
    if b != ')' {
        self.causePanic()
    }
}
func (self *MemberDescriptorParser) finish() {
    if self.offset != len(self.descriptor) {
        self.causePanic()
    }
}

func (self *MemberDescriptorParser) readFieldType() (*FieldType) {
    switch self.readUint8() {
    case 'B': return baseTypeB
    case 'C': return baseTypeC
    case 'D': return baseTypeD
    case 'F': return baseTypeF
    case 'I': return baseTypeI
    case 'J': return baseTypeJ
    case 'S': return baseTypeS
    case 'Z': return baseTypeZ
    case 'V': return baseTypeV
    case 'L': return self.readObjectType()
    case '[': return self.readArrayType()
    default:
        self.unreadUint8()
        return nil
    }
}
func (self *MemberDescriptorParser) readObjectType() (*FieldType) {
    unread := self.descriptor[self.offset:]
    semicolonIndex := strings.IndexRune(unread, ';')
    if semicolonIndex == -1 {
        self.causePanic()
        return nil
    } else {
        objStart := self.offset - 1
        objEnd := self.offset + semicolonIndex + 1
        self.offset = objEnd
        descriptor := self.descriptor[objStart: objEnd]
        return &FieldType{descriptor}
    }
}
func (self *MemberDescriptorParser) readArrayType() (*FieldType) {
    arrStart := self.offset - 1
    self.readFieldType()
    arrEnd := self.offset
    descriptor := self.descriptor[arrStart: arrEnd]
    return &FieldType{descriptor}
}

func (self *MemberDescriptorParser) causePanic() {
    panic("BAD descriptor: " + self.descriptor)
}
