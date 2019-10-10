package heap

var (
	_shimClass  = &Class{Name: "~shim"}
	_returnCode = []byte{0xb1} // return
	_athrowCode = []byte{0xbf} // athrow

	_returnMethod = &Method{
		ClassMember: ClassMember{
			AccessFlags: AccessFlags(AccStatic),
			Name:        "<return>",
			Class:       _shimClass,
		},
		MethodData: MethodData{
			Code: _returnCode,
		},
	}

	_athrowMethod = &Method{
		ClassMember: ClassMember{
			AccessFlags: AccessFlags(AccStatic),
			Name:        "<athrow>",
			Class:       _shimClass,
		},
		MethodData: MethodData{
			Code: _athrowCode,
		},
	}
)

func ReturnMethod() *Method {
	return _returnMethod
}

func AthrowMethod() *Method {
	return _athrowMethod
}

func BootstrapMethod() *Method {
	method := &Method{}
	method.Class = &Class{Name: "~shim"}
	method.Name = "<bootstrap>"
	method.AccessFlags = AccStatic
	method.MaxStack = 8
	method.MaxLocals = 8
	method.ParamSlotCount = 2
	method.Code = []byte{0xff, 0xb1} // bootstrap, return
	return method
}
