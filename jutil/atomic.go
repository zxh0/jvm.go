package jutil

import (
	"sync/atomic"
)

func CasInt32(any interface{}, old, _new int32) bool {
	addr, ok := any.(int32)
	if ok {
		return atomic.CompareAndSwapInt32(&addr, old, _new)
	}
	return false
}

func CasInt64(any interface{}, old, _new int64) bool {
	addr, ok := any.(int64)
	if ok {
		return atomic.CompareAndSwapInt64(&addr, old, _new)
	}
	return false
}
