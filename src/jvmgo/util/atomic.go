package util

import (
	. "jvmgo/any"
	"sync/atomic"
	"unsafe"
)

// copied from go/src/sync/atomic/value.go
type ifaceWords struct {
	typ  unsafe.Pointer
	data unsafe.Pointer
}

func CasInt32(any Any, old, _new int32) bool {
	ifws := ((*ifaceWords)(unsafe.Pointer(&any)))
	addr := (*int32)(ifws.data)
	return atomic.CompareAndSwapInt32(addr, old, _new)
}
