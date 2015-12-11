package heap

import (
	"strings"
)

type MethodDescriptorParser struct {
	descriptor string
	offset     int
	md         *MethodDescriptor
}

func (self *MethodDescriptorParser) parse() *MethodDescriptor {
	self.md = &MethodDescriptor{}
	self.md.d = self.descriptor
	self.startParams()
	self.parseParamTypes()
	self.endParams()
	self.parseReturnType()
	self.finish()
	return self.md
}

func (self *MethodDescriptorParser) startParams() {
	if self.readUint8() != '(' {
		self.causePanic()
	}
}
func (self *MethodDescriptorParser) endParams() {
	if self.readUint8() != ')' {
		self.causePanic()
	}
}
func (self *MethodDescriptorParser) finish() {
	if self.offset != len(self.descriptor) {
		self.causePanic()
	}
}

func (self *MethodDescriptorParser) causePanic() {
	panic("BAD descriptor: " + self.descriptor)
}

func (self *MethodDescriptorParser) readUint8() uint8 {
	b := self.descriptor[self.offset]
	self.offset++
	return b
}
func (self *MethodDescriptorParser) unreadUint8() {
	self.offset--
}

func (self *MethodDescriptorParser) parseParamTypes() {
	for {
		t := self.parseFieldType()
		if t != nil {
			self.md.addParameterType(t)
		} else {
			break
		}
	}
}
func (self *MethodDescriptorParser) parseReturnType() {
	t := self.parseFieldType()
	if t != nil {
		self.md.returnType = t
	} else {
		self.causePanic()
	}
}

func (self *MethodDescriptorParser) parseFieldType() *FieldType {
	switch self.readUint8() {
	case 'B':
		return baseTypeB
	case 'C':
		return baseTypeC
	case 'D':
		return baseTypeD
	case 'F':
		return baseTypeF
	case 'I':
		return baseTypeI
	case 'J':
		return baseTypeJ
	case 'S':
		return baseTypeS
	case 'Z':
		return baseTypeZ
	case 'V':
		return baseTypeV
	case 'L':
		return self.parseObjectType()
	case '[':
		return self.parseArrayType()
	default:
		self.unreadUint8()
		return nil
	}
}
func (self *MethodDescriptorParser) parseObjectType() *FieldType {
	unread := self.descriptor[self.offset:]
	semicolonIndex := strings.IndexRune(unread, ';')
	if semicolonIndex == -1 {
		self.causePanic()
		return nil
	} else {
		objStart := self.offset - 1
		objEnd := self.offset + semicolonIndex + 1
		self.offset = objEnd
		descriptor := self.descriptor[objStart:objEnd]
		return &FieldType{descriptor}
	}
}
func (self *MethodDescriptorParser) parseArrayType() *FieldType {
	arrStart := self.offset - 1
	self.parseFieldType()
	arrEnd := self.offset
	descriptor := self.descriptor[arrStart:arrEnd]
	return &FieldType{descriptor}
}
