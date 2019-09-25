package heap

import (
	"strings"
)

type MethodDescriptorParser struct {
	descriptor string
	offset     int
	md         *MethodDescriptor
}

func (parser *MethodDescriptorParser) parse() *MethodDescriptor {
	parser.md = &MethodDescriptor{}
	parser.md.d = parser.descriptor
	parser.startParams()
	parser.parseParamTypes()
	parser.endParams()
	parser.parseReturnType()
	parser.finish()
	return parser.md
}

func (parser *MethodDescriptorParser) startParams() {
	if parser.readUint8() != '(' {
		parser.causePanic()
	}
}
func (parser *MethodDescriptorParser) endParams() {
	if parser.readUint8() != ')' {
		parser.causePanic()
	}
}
func (parser *MethodDescriptorParser) finish() {
	if parser.offset != len(parser.descriptor) {
		parser.causePanic()
	}
}

func (parser *MethodDescriptorParser) causePanic() {
	panic("BAD descriptor: " + parser.descriptor)
}

func (parser *MethodDescriptorParser) readUint8() uint8 {
	b := parser.descriptor[parser.offset]
	parser.offset++
	return b
}
func (parser *MethodDescriptorParser) unreadUint8() {
	parser.offset--
}

func (parser *MethodDescriptorParser) parseParamTypes() {
	for {
		t := parser.parseFieldType()
		if t != nil {
			parser.md.addParameterType(t)
		} else {
			break
		}
	}
}
func (parser *MethodDescriptorParser) parseReturnType() {
	t := parser.parseFieldType()
	if t != nil {
		parser.md.returnType = t
	} else {
		parser.causePanic()
	}
}

func (parser *MethodDescriptorParser) parseFieldType() *FieldType {
	switch parser.readUint8() {
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
		return parser.parseObjectType()
	case '[':
		return parser.parseArrayType()
	default:
		parser.unreadUint8()
		return nil
	}
}
func (parser *MethodDescriptorParser) parseObjectType() *FieldType {
	unread := parser.descriptor[parser.offset:]
	semicolonIndex := strings.IndexRune(unread, ';')
	if semicolonIndex == -1 {
		parser.causePanic()
		return nil
	} else {
		objStart := parser.offset - 1
		objEnd := parser.offset + semicolonIndex + 1
		parser.offset = objEnd
		descriptor := parser.descriptor[objStart:objEnd]
		return &FieldType{descriptor}
	}
}
func (parser *MethodDescriptorParser) parseArrayType() *FieldType {
	arrStart := parser.offset - 1
	parser.parseFieldType()
	arrEnd := parser.offset
	descriptor := parser.descriptor[arrStart:arrEnd]
	return &FieldType{descriptor}
}
