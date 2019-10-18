package jni

// #include "jni.h"
import "C"

type (
	JNIEnv    = *C.JNIEnv
	cstring   = *C.char
	cint      = C.int
	jsize     = C.jsize
	jboolean  = C.jboolean
	jchar     = C.jchar
	jint      = C.jint
	jobject   = C.jobject
	jclass    = C.jclass
	jstring   = C.jstring
	jmethodID = C.jmethodID
)
