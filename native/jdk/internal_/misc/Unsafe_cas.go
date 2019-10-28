package misc

import (
	"sync/atomic"
	"unsafe"

	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

func init() {
	_unsafe(compareAndSetInt, "compareAndSetInt", "(Ljava/lang/Object;JII)Z")
	_unsafe(compareAndSetLong, "compareAndSetLong", "(Ljava/lang/Object;JJJ)Z")
	_unsafe(compareAndSetReference, "compareAndSetReference", "(Ljava/lang/Object;JLjava/lang/Object;Ljava/lang/Object;)Z")
}

// public final native boolean compareAndSetInt(Object o, long offset, int expected, int x);
// (Ljava/lang/Object;JII)Z
func compareAndSetInt(frame *rtda.Frame) {
	obj := frame.GetRefVar(1)
	offset := frame.GetLongVar(2)
	expected := frame.GetIntVar(4)
	newVal := frame.GetIntVar(5)

	fields := obj.Fields
	if slots, ok := fields.([]heap.Slot); ok {
		// object
		swapped := atomic.CompareAndSwapInt64(&(slots[offset].Val), int64(expected), int64(newVal))
		frame.PushBoolean(swapped)
	} else if ints, ok := fields.([]int32); ok {
		// int[]
		swapped := atomic.CompareAndSwapInt32(&ints[offset], expected, newVal)
		frame.PushBoolean(swapped)
	} else {
		// todo
		panic("todo: compareAndSwapInt!")
	}
}

// public final native boolean compareAndSetLong(Object o, long offset, long expected, long x);
// (Ljava/lang/Object;JJJ)Z
func compareAndSetLong(frame *rtda.Frame) {
	obj := frame.GetRefVar(1)
	offset := frame.GetLongVar(2)
	expected := frame.GetLongVar(4)
	newVal := frame.GetLongVar(6)

	fields := obj.Fields
	if slots, ok := fields.([]heap.Slot); ok {
		// object
		swapped := atomic.CompareAndSwapInt64(&(slots[offset].Val), expected, newVal)
		frame.PushBoolean(swapped)
	} else if ints, ok := fields.([]int64); ok {
		// long[]
		swapped := atomic.CompareAndSwapInt64(&ints[offset], expected, newVal)
		frame.PushBoolean(swapped)
	} else {
		// todo
		panic("todo: compareAndSetLong!")
	}
}

// public final native boolean compareAndSetReference(Object o, long offset, Object expected, Object x)
// (Ljava/lang/Object;JLjava/lang/Object;Ljava/lang/Object;)Z
func compareAndSetReference(frame *rtda.Frame) {
	obj := frame.GetRefVar(1)
	offset := frame.GetLongVar(2)
	expected := frame.GetRefVar(4)
	newVal := frame.GetRefVar(5)

	// todo
	fields := obj.Fields
	if slots, ok := fields.([]heap.Slot); ok {
		// object
		swapped := _casObj(obj, slots, offset, expected, newVal)
		frame.PushBoolean(swapped)
	} else if objs, ok := fields.([]*heap.Object); ok {
		// ref[]
		swapped := _casArr(objs, offset, expected, newVal)
		frame.PushBoolean(swapped)
	} else {
		// todo
		panic("todo: compareAndSetReference!")
	}
}
func _casObj(obj *heap.Object, fields []heap.Slot, offset int64, expected, newVal *heap.Object) bool {
	// todo
	obj.LockState()
	defer obj.UnlockState()

	current := fields[offset].Ref
	if current == expected {
		fields[offset].Ref = newVal
		return true
	} else {
		return false
	}
}
func _casArr(objs []*heap.Object, offset int64, expected, newVal *heap.Object) bool {
	// cast to []unsafe.Pointer
	ps := *((*[]unsafe.Pointer)(unsafe.Pointer(&objs)))

	addr := &ps[offset]
	old := unsafe.Pointer(expected)
	_new := unsafe.Pointer(newVal)
	return atomic.CompareAndSwapPointer(addr, old, _new)
}
