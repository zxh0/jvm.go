package jni

// #include <stdlib.h>
// #include "jni.h"
// JNIEnv NewJNIEnvWrapper();
import "C"
import (
	"unsafe"

	"github.com/zxh0/jvm.go/interpreter"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
	"github.com/zxh0/jvm.go/vmutils"
)

func getEnv(frame *rtda.Frame) JNIEnv {
	thread := frame.Thread
	if env := thread.JNIEnv; env != nil {
		return env.(JNIEnv)
	}
	env := C.NewJNIEnvWrapper()
	env.reserved0 = unsafe.Pointer(thread)
	thread.JNIEnv = &env
	return &env
}
func getThread(env JNIEnv) *rtda.Thread {
	return (*rtda.Thread)((*env).reserved0)
}

//export GoPanic
func GoPanic(err cstring) {
	panic(C.GoString(err))
}

//export GoGetParamCount
func GoGetParamCount(clazz jclass, methodID jmethodID) cint {
	method := getMethod(clazz, methodID)
	return cint(method.ParamCount)
}

//export GoGetVersion
func GoGetVersion(env JNIEnv) jint {
	return 0 // TODO
}

//export GoFindClass
func GoFindClass(env JNIEnv, _name cstring) jclass {
	name := C.GoString(_name)
	//println("GoFindClass:", name)
	cls := heap.BootLoader().FindLoadedClass(name)
	if cls == nil {
		panic("GoFindClass: " + name)
	}
	return jclassFromGo(cls)
}

//export GoNewGlobalRef
func GoNewGlobalRef(env JNIEnv, lobj jobject) jobject {
	// TODO
	return lobj
}

//export GoDeleteLocalRef
func GoDeleteLocalRef(env JNIEnv, obj jobject) {
	// TODO
}

//export GoNewString
func GoNewString(env JNIEnv, unicode *jchar, _len jsize) jstring {
	goBytes := C.GoBytes(unsafe.Pointer(unicode), _len*2)
	jChars := vmutils.CastBytesToUint16s(goBytes)
	jStr := heap.JSFromChars(jChars)
	//println("GoNewString:", heap.JSToGoStr(jStr))
	return jstringFromGo(jStr)
}

//export GoNewStringUTF
func GoNewStringUTF(env JNIEnv, utf cstring) jstring {
	goStr := C.GoString(utf)
	//println("GoNewStringUTF:", goStr)
	jStr := heap.JSFromGoStr(goStr)
	return jstringFromGo(jStr)
}

//export GoGetStringLength
func GoGetStringLength(env JNIEnv, str jstring) jsize {
	jStr := jstringToGo(str)
	//println("GoGetStringLength:", heap.JSToGoStr(jStr))
	return jsize(len(heap.JSToChars(jStr)))
}

//export GoGetStringCritical
func GoGetStringCritical(env JNIEnv, str jstring, isCopy *jboolean) *jchar {
	jStr := jstringToGo(str)
	//println("GoGetStringCritical:", heap.JSToGoStr(jStr))
	jChars := heap.JSToChars(jStr)
	nChars := len(jChars)
	cChars := (*jchar)(C.malloc(2 * C.size_t(nChars)))
	cCharsAsSlice := cCharArrToSlice(cChars, nChars)
	copy(cCharsAsSlice, jChars)
	if isCopy != nil {
		*isCopy = C.JNI_TRUE
	}
	return cChars
}

//export GoReleaseStringCritical
func GoReleaseStringCritical(env JNIEnv, str jstring, cstr *jchar) {
	C.free(unsafe.Pointer(cstr))
}

//export GoGetMethodID
func GoGetMethodID(env JNIEnv, clazz jclass, _name, _sig cstring) jmethodID {
	name := C.GoString(_name)
	sig := C.GoString(_sig)
	cls := jclassToGo(clazz)
	//println("GoGetMethodID:", cls.Name, name, sig)
	method := cls.GetInstanceMethod(name, sig)
	return jmethodIDFromGo(method.Slot)
}

//export GoGetStaticMethodID
func GoGetStaticMethodID(env JNIEnv, clazz jclass, _name, _sig cstring) jmethodID {
	name := C.GoString(_name)
	sig := C.GoString(_sig)
	cls := jclassToGo(clazz)
	//println("GoGetStaticMethodID:", cls.Name, name, sig)
	method := cls.GetStaticMethod(name, sig)
	return jmethodIDFromGo(method.Slot)
}

//export GoCallStaticObjectMethodA
func GoCallStaticObjectMethodA(env JNIEnv, clazz jclass, methodID jmethodID, _args unsafe.Pointer) jobject {
	method := getMethod(clazz, methodID)
	//println("GoCallStaticObjectMethodA:", method.Class.Name, method.Name, method.Descriptor)
	thread := getThread(env)
	args := vaListToSlots(_args, method.ParsedDescriptor)
	retVal := interpreter.ExecMethod(thread, method, args) // TODO
	return jobjectFromGo(retVal.Ref)
}

//export GoExceptionCheck
func GoExceptionCheck(env JNIEnv) jboolean {
	// TODO
	return C.JNI_FALSE
}
