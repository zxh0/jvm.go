package native

import (
	"strings"

	"github.com/zxh0/jvm.go/vmutils"
)

type RegForClass struct {
	className string
	fnPrefix  string
}

func ForClass(className string) *RegForClass {
	return &RegForClass{className: className}
}

func (reg *RegForClass) RemovePrefix(fnPrefix string) *RegForClass {
	reg.fnPrefix = fnPrefix
	return reg
}

func (reg *RegForClass) Register(method Method, descriptor string) *RegForClass {
	methodName := vmutils.GetFuncName(method)
	if reg.fnPrefix != "" {
		methodName = strings.TrimPrefix(methodName, reg.fnPrefix)
	}
	Register(reg.className, methodName, descriptor, method)
	return reg
}
