package heap

var _internedStrings = map[string]*Object{}

func getInternedString(goStr string) *Object {
	return _internedStrings[goStr]
}

// todo
func InternString(goStr string, jStr *Object) *Object {
	if internedStr, ok := _internedStrings[goStr]; ok {
		return internedStr
	}

	_internedStrings[goStr] = jStr
	return jStr
}
