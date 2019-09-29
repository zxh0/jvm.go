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
	panic("invalid descriptor: " + parser.descriptor)
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
		if t := parser.parseFieldType(); t != "" {
			parser.md.addParameterType(t)
		} else {
			break
		}
	}
}
func (parser *MethodDescriptorParser) parseReturnType() {
	if t := parser.parseFieldType(); t != "" {
		parser.md.ReturnType = t
	} else {
		parser.causePanic()
	}
}

func (parser *MethodDescriptorParser) parseFieldType() FieldType {
	switch parser.readUint8() {
	case 'B':
		return "B"
	case 'C':
		return "C"
	case 'D':
		return "D"
	case 'F':
		return "F"
	case 'I':
		return "I"
	case 'J':
		return "J"
	case 'S':
		return "S"
	case 'Z':
		return "Z"
	case 'V':
		return "V"
	case 'L':
		return parser.parseObjectType()
	case '[':
		return parser.parseArrayType()
	default:
		parser.unreadUint8()
		return ""
	}
}
func (parser *MethodDescriptorParser) parseObjectType() FieldType {
	unread := parser.descriptor[parser.offset:]
	semicolonIndex := strings.IndexRune(unread, ';')
	if semicolonIndex == -1 {
		parser.causePanic()
		return ""
	} else {
		objStart := parser.offset - 1
		objEnd := parser.offset + semicolonIndex + 1
		parser.offset = objEnd
		descriptor := parser.descriptor[objStart:objEnd]
		return FieldType(descriptor)
	}
}
func (parser *MethodDescriptorParser) parseArrayType() FieldType {
	arrStart := parser.offset - 1
	parser.parseFieldType()
	arrEnd := parser.offset
	descriptor := parser.descriptor[arrStart:arrEnd]
	return FieldType(descriptor)
}
