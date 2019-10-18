package jni

// #include "jni.h"
// JNIEnv NewJNIEnvWrapper();
import "C"
import (
	"unsafe"

	"github.com/zxh0/jvm.go/rtda/heap"
	"github.com/zxh0/jvm.go/vmutils"
)

// TODO: jxxxFromGo

func jclassFromGo(cls *heap.Class) jclass {
	return jclass(unsafe.Pointer(cls))
}
func jclassToGo(ptr jclass) *heap.Class {
	return (*heap.Class)(unsafe.Pointer(ptr))
}

func jobjectFromGo(obj *heap.Object) jobject {
	return jobject(unsafe.Pointer(obj))
}
func jobjectToGo(ptr jobject) *heap.Object {
	return (*heap.Object)(unsafe.Pointer(ptr))
}

func jstringFromGo(obj *heap.Object) jstring {
	return jstring(unsafe.Pointer(obj))
}
func jstringToGo(ptr jstring) *heap.Object {
	return (*heap.Object)(unsafe.Pointer(ptr))
}

func jmethodIDFromGo(id uint) jmethodID {
	return jmethodID(unsafe.Pointer(uintptr(id)))
}
func jmethodIDToGo(mid jmethodID) uint {
	return uint(uintptr(unsafe.Pointer(mid)))
}

func getMethod(clazz jclass, methodID jmethodID) *heap.Method {
	cls := jclassToGo(clazz)
	mid := jmethodIDToGo(methodID)
	return cls.Methods[mid]
}

func vaListToSlots(args unsafe.Pointer, pd heap.ParsedDescriptor) []heap.Slot {
	nParams := len(pd.ParameterTypes)
	rawArgBytes := C.GoBytes(args, cint(nParams*8))
	jVals := vmutils.CastBytesToInt64s(rawArgBytes)

	slots := make([]heap.Slot, 0, nParams)
	for i, paramType := range pd.ParameterTypes {
		jVal := jVals[i]
		switch paramType {
		case "B", "C", "F", "I", "S", "Z":
			slots = append(slots, heap.Slot{Val: jVal})
		case "D", "J":
			slots = append(slots, heap.Slot{Val: jVal}, heap.EmptySlot)
		default:
			ref := (*heap.Object)(unsafe.Pointer(uintptr(jVal)))
			slots = append(slots, heap.NewRefSlot(ref))
		}
	}

	return slots
}

// https://github.com/golang/go/wiki/cgo#turning-c-arrays-into-go-slices
func cCharArrToSlice(ptr *jchar, n int) []uint16 {
	return (*[1 << 28]uint16)(unsafe.Pointer(ptr))[:n:n]
}
