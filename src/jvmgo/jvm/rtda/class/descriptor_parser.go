package class

import (
    "strings"
)

type MemberDescriptorParser struct {
    descriptor  string
    offset      int
    md          *MethodDescriptor
}

func newMemberDescriptorParser(descriptor string) (*MemberDescriptorParser) {
    return &MemberDescriptorParser{descriptor: descriptor}
}

func (self *MemberDescriptorParser) parse() (*MethodDescriptor) {
    self.md = &MethodDescriptor{}
    self.startParams()
    self.parseParameterTypes()
    self.endParams()
    self.parseReturnType()
    self.finish()
    return self.md
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

func (self *MemberDescriptorParser) parseParameterTypes() {
    for {
        t := self.readFieldType()
        if t != nil {
            self.md.addParameterType(t)
        } else {
            break
        }
    }
}
func (self *MemberDescriptorParser) parseReturnType() {
    t := self.readFieldType()
    if t != nil {
        self.md.returnType = t
    } else {
        self.causePanic()
    }
}

func (self *MemberDescriptorParser) readUint8() uint8 {
    b := self.descriptor[self.offset]
    self.offset++
    return b
}
func (self *MemberDescriptorParser) unreadUint8() {
    self.offset--
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
