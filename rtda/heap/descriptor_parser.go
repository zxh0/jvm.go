package heap

type MethodDescriptorParser struct {
	d string
}

func calcArgSlotCount(descriptor string) uint {
	return parseMethodDescriptor(descriptor).argSlotCount()
}

func parseMethodDescriptor(descriptor string) ParsedDescriptor {
	parser := &MethodDescriptorParser{descriptor}
	return parser.parse()
}

func (parser *MethodDescriptorParser) parse() ParsedDescriptor {
	if paramTypes, ok := parser.parseParamTypes(); ok {
		if returnType, ok := parser.parseReturnType(); ok {
			return ParsedDescriptor{
				ParameterTypes: paramTypes,
				ReturnType:     returnType,
			}
		}
	}
	panic("invalid descriptor: " + parser.d) // TODO
}

func (parser *MethodDescriptorParser) parseReturnType() (FieldOrReturnType, bool) {
	if t, ok := parser.parseFieldType(); ok {
		return t, len(parser.d) == 0
	}
	return "V", parser.d == "V"
}

func (parser *MethodDescriptorParser) parseParamTypes() ([]FieldOrReturnType, bool) {
	if len(parser.d) == 0 && parser.d[0] != '(' {
		return nil, false
	}
	parser.d = parser.d[1:]

	var ts []FieldOrReturnType = nil
	for {
		if t, ok := parser.parseFieldType(); ok {
			ts = append(ts, t)
		} else {
			break
		}
	}

	if len(parser.d) == 0 && parser.d[0] != ')' {
		return nil, false
	}
	parser.d = parser.d[1:]
	return ts, true
}

func (parser *MethodDescriptorParser) parseFieldType() (FieldOrReturnType, bool) {
	if len(parser.d) > 0 {
		switch parser.d[0] {
		case 'B', 'C', 'D', 'F', 'I', 'J', 'S', 'Z':
			t := parser.d[0:1]
			parser.d = parser.d[1:]
			return FieldOrReturnType(t), true
		case 'L':
			return parser.parseObjectType()
		case '[':
			return parser.parseArrayType()
		}
	}
	return "", false
}

func (parser *MethodDescriptorParser) parseObjectType() (FieldOrReturnType, bool) {
	for i := 0; i < len(parser.d); i++ {
		if parser.d[i] == ';' {
			t := parser.d[:i+1]
			parser.d = parser.d[i+1:]
			return FieldOrReturnType(t), true
		}
	}
	return "", false
}

func (parser *MethodDescriptorParser) parseArrayType() (FieldOrReturnType, bool) {
	parser.d = parser.d[1:]
	t, ok := parser.parseFieldType()
	return "[" + t, ok
}
