package vmutils

import (
	"reflect"
	"runtime"
	"strings"
)

func GetFuncName(f interface{}) string {
	name := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	return removePackage(name)
}

func removePackage(funcName string) string {
	lastDotIdx := strings.LastIndexByte(funcName, '.')
	if lastDotIdx < 0 {
		return funcName
	}
	return funcName[lastDotIdx+1:]
}
