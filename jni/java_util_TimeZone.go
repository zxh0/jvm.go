package jni

// #include "jni.h"
// jstring JNICALL Java_java_util_TimeZone_getSystemTimeZoneID(JNIEnv *env, jclass ign, jstring java_home);
// jstring JNICALL Java_java_util_TimeZone_getSystemGMTOffsetID(JNIEnv *env, jclass ign);
// JNIEnv NewJNIEnvWrapper();
import "C"
import (
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

func init() {
	_tz(getSystemTimeZoneID, "getSystemTimeZoneID", "(Ljava/lang/String;)Ljava/lang/String;")
	_tz(getSystemGMTOffsetID, "getSystemGMTOffsetID", "()Ljava/lang/String;")
}

func _tz(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("java/util/TimeZone", name, desc, method)
}

// private static native String getSystemTimeZoneID(String javaHome);
// (Ljava/lang/String;Ljava/lang/String;)Ljava/lang/String;
func getSystemTimeZoneID(frame *rtda.Frame) {
	javaHome := frame.GetRefVar(0)

	cEnv := getEnv(frame)
	cCls := jclassFromGo(nil)
	cJH := jstringFromGo(javaHome)
	cRet := C.Java_java_util_TimeZone_getSystemTimeZoneID(cEnv, cCls, cJH)

	zoneID := jstringToGo(cRet)
	frame.PushRef(zoneID)
}

// private static native String getSystemGMTOffsetID();
// ()Ljava/lang/String;
func getSystemGMTOffsetID(frame *rtda.Frame) {
	cEnv := getEnv(frame)
	cCls := jclassFromGo(nil)
	cRet := C.Java_java_util_TimeZone_getSystemGMTOffsetID(cEnv, cCls)

	zoneID := jstringToGo(cRet)
	frame.PushRef(zoneID)
}
