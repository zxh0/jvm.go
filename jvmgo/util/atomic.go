package util

import (
	. "github.com/zxh0/jvm.go/jvmgo/any"
	"sync/atomic"
	"unsafe"
)

// copied from go/src/sync/atomic/value.go
type ifaceWords struct {
	typ  unsafe.Pointer
	data unsafe.Pointer
}

func CasInt32(any Any, old, _new int32) bool {
	ifws := (*ifaceWords)(unsafe.Pointer(&any))
	addr := (*int32)(ifws.data)
	return atomic.CompareAndSwapInt32(addr, old, _new)
}

func CasInt64(any Any, old, _new int64) bool {
	ifws := (*ifaceWords)(unsafe.Pointer(&any))
	addr := (*int64)(ifws.data)
	return atomic.CompareAndSwapInt64(addr, old, _new)
}
