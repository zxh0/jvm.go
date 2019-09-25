package heap

var (
	_shimClass  = &Class{name: "~shim"}
	_returnCode = []byte{0xb1} // return
	_athrowCode = []byte{0xbf} // athrow

	_returnMethod = &Method{
		ClassMember: ClassMember{
			AccessFlags: AccessFlags{ACC_STATIC},
			name:        "<return>",
			class:       _shimClass,
		},
		code: _returnCode,
	}

	_athrowMethod = &Method{
		ClassMember: ClassMember{
			AccessFlags: AccessFlags{ACC_STATIC},
			name:        "<athrow>",
			class:       _shimClass,
		},
		code: _athrowCode,
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
	method.class = &Class{name: "~shim"}
	method.name = "<bootstrap>"
	method.accessFlags = ACC_STATIC
	method.maxStack = 8
	method.maxLocals = 8
	method.argSlotCount = 2
	method.code = []byte{0xff, 0xb1} // bootstrap, return
	return method
}
