package jni

// #cgo CFLAGS: -I/Users/zxh/.sdkman/candidates/java/8.0.222-zulu/include
// #cgo CFLAGS: -I/Users/zxh/.sdkman/candidates/java/8.0.222-zulu/include/darwin
// #include <stdio.h>
// #include <stdlib.h>
// #include "jni.h"
//
// extern jint (JNICALL GetVersion)(JNIEnv *env);
// extern jclass (JNICALL DefineClass) (JNIEnv *env, const char *name, jobject loader, const jbyte *buf, jsize len);
// extern jclass (JNICALL FindClass) (JNIEnv *env, const char *name);
// jmethodID (JNICALL FromReflectedMethod) (JNIEnv *env, jobject method);
// jfieldID (JNICALL FromReflectedField) (JNIEnv *env, jobject field);
// jobject (JNICALL ToReflectedMethod) (JNIEnv *env, jclass cls, jmethodID methodID, jboolean isStatic);
// jclass (JNICALL GetSuperclass) (JNIEnv *env, jclass sub);
// jboolean (JNICALL IsAssignableFrom) (JNIEnv *env, jclass sub, jclass sup);
// jobject (JNICALL ToReflectedField) (JNIEnv *env, jclass cls, jfieldID fieldID, jboolean isStatic);
// jint (JNICALL Throw) (JNIEnv *env, jthrowable obj);
// jint (JNICALL ThrowNew) (JNIEnv *env, jclass clazz, const char *msg);
// jthrowable (JNICALL ExceptionOccurred) (JNIEnv *env);
// void (JNICALL ExceptionDescribe) (JNIEnv *env);
// void (JNICALL ExceptionClear) (JNIEnv *env);
// void (JNICALL FatalError) (JNIEnv *env, const char *msg);
// jint (JNICALL PushLocalFrame) (JNIEnv *env, jint capacity);
// jobject (JNICALL PopLocalFrame) (JNIEnv *env, jobject result);
// jobject (JNICALL NewGlobalRef) (JNIEnv *env, jobject lobj);
// void (JNICALL DeleteGlobalRef) (JNIEnv *env, jobject gref);
// void (JNICALL DeleteLocalRef) (JNIEnv *env, jobject obj);
// jboolean (JNICALL IsSameObject) (JNIEnv *env, jobject obj1, jobject obj2);
// jobject (JNICALL NewLocalRef) (JNIEnv *env, jobject ref);
// jint (JNICALL EnsureLocalCapacity) (JNIEnv *env, jint capacity);
// jobject (JNICALL AllocObject) (JNIEnv *env, jclass clazz);
// jobject (JNICALL NewObject) (JNIEnv *env, jclass clazz, jmethodID methodID, ...);
// jobject (JNICALL NewObjectV) (JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
// jobject (JNICALL NewObjectA) (JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);
// jclass (JNICALL GetObjectClass) (JNIEnv *env, jobject obj);
// jboolean (JNICALL IsInstanceOf) (JNIEnv *env, jobject obj, jclass clazz);
// jmethodID (JNICALL GetMethodID) (JNIEnv *env, jclass clazz, const char *name, const char *sig);
// jobject (JNICALL CallObjectMethod) (JNIEnv *env, jobject obj, jmethodID methodID, ...);
// jobject (JNICALL CallObjectMethodV) (JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
// jobject (JNICALL CallObjectMethodA) (JNIEnv *env, jobject obj, jmethodID methodID, const jvalue * args);
// jboolean (JNICALL CallBooleanMethod) (JNIEnv *env, jobject obj, jmethodID methodID, ...);
// jboolean (JNICALL CallBooleanMethodV) (JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
// jboolean (JNICALL CallBooleanMethodA) (JNIEnv *env, jobject obj, jmethodID methodID, const jvalue * args);
// jbyte (JNICALL CallByteMethod) (JNIEnv *env, jobject obj, jmethodID methodID, ...);
// jbyte (JNICALL CallByteMethodV) (JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
// jbyte (JNICALL CallByteMethodA) (JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);
// jchar (JNICALL CallCharMethod) (JNIEnv *env, jobject obj, jmethodID methodID, ...);
// jchar (JNICALL CallCharMethodV) (JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
// jchar (JNICALL CallCharMethodA) (JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);
// jshort (JNICALL CallShortMethod) (JNIEnv *env, jobject obj, jmethodID methodID, ...);
// jshort (JNICALL CallShortMethodV) (JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
// jshort (JNICALL CallShortMethodA) (JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);
// jint (JNICALL CallIntMethod) (JNIEnv *env, jobject obj, jmethodID methodID, ...);
// jint (JNICALL CallIntMethodV) (JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
// jint (JNICALL CallIntMethodA) (JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);
// jlong (JNICALL CallLongMethod) (JNIEnv *env, jobject obj, jmethodID methodID, ...);
// jlong (JNICALL CallLongMethodV) (JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
// jlong (JNICALL CallLongMethodA) (JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);
// jfloat (JNICALL CallFloatMethod) (JNIEnv *env, jobject obj, jmethodID methodID, ...);
// jfloat (JNICALL CallFloatMethodV) (JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
// jfloat (JNICALL CallFloatMethodA) (JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);
// jdouble (JNICALL CallDoubleMethod) (JNIEnv *env, jobject obj, jmethodID methodID, ...);
// jdouble (JNICALL CallDoubleMethodV) (JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
// jdouble (JNICALL CallDoubleMethodA) (JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);
// void (JNICALL CallVoidMethod) (JNIEnv *env, jobject obj, jmethodID methodID, ...);
// void (JNICALL CallVoidMethodV) (JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
// void (JNICALL CallVoidMethodA) (JNIEnv *env, jobject obj, jmethodID methodID, const jvalue * args);
// jobject (JNICALL CallNonvirtualObjectMethod) (JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
// jobject (JNICALL CallNonvirtualObjectMethodV) (JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args);
// jobject (JNICALL CallNonvirtualObjectMethodA) (JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue * args);
// jboolean (JNICALL CallNonvirtualBooleanMethod) (JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
// jboolean (JNICALL CallNonvirtualBooleanMethodV) (JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args);
// jboolean (JNICALL CallNonvirtualBooleanMethodA) (JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue * args);
// jbyte (JNICALL CallNonvirtualByteMethod) (JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
// jbyte (JNICALL CallNonvirtualByteMethodV) (JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args);
// jbyte (JNICALL CallNonvirtualByteMethodA) (JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue *args);
// jchar (JNICALL CallNonvirtualCharMethod) (JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
// jchar (JNICALL CallNonvirtualCharMethodV) (JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args);
// jchar (JNICALL CallNonvirtualCharMethodA) (JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue *args);
// jshort (JNICALL CallNonvirtualShortMethod) (JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
// jshort (JNICALL CallNonvirtualShortMethodV) (JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args);
// jshort (JNICALL CallNonvirtualShortMethodA) (JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue *args);
// jint (JNICALL CallNonvirtualIntMethod) (JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
// jint (JNICALL CallNonvirtualIntMethodV) (JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args);
// jint (JNICALL CallNonvirtualIntMethodA) (JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue *args);
// jlong (JNICALL CallNonvirtualLongMethod) (JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
// jlong (JNICALL CallNonvirtualLongMethodV) (JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args);
// jlong (JNICALL CallNonvirtualLongMethodA) (JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue *args);
// jfloat (JNICALL CallNonvirtualFloatMethod) (JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
// jfloat (JNICALL CallNonvirtualFloatMethodV) (JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args);
// jfloat (JNICALL CallNonvirtualFloatMethodA) (JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue *args);
// jdouble (JNICALL CallNonvirtualDoubleMethod) (JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
// jdouble (JNICALL CallNonvirtualDoubleMethodV) (JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args);
// jdouble (JNICALL CallNonvirtualDoubleMethodA) (JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue *args);
// void (JNICALL CallNonvirtualVoidMethod) (JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
// void (JNICALL CallNonvirtualVoidMethodV) (JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args);
// void (JNICALL CallNonvirtualVoidMethodA) (JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue * args);
// jfieldID (JNICALL GetFieldID) (JNIEnv *env, jclass clazz, const char *name, const char *sig);
// jobject (JNICALL GetObjectField) (JNIEnv *env, jobject obj, jfieldID fieldID);
// jboolean (JNICALL GetBooleanField) (JNIEnv *env, jobject obj, jfieldID fieldID);
// jbyte (JNICALL GetByteField) (JNIEnv *env, jobject obj, jfieldID fieldID);
// jchar (JNICALL GetCharField) (JNIEnv *env, jobject obj, jfieldID fieldID);
// jshort (JNICALL GetShortField) (JNIEnv *env, jobject obj, jfieldID fieldID);
// jint (JNICALL GetIntField) (JNIEnv *env, jobject obj, jfieldID fieldID);
// jlong (JNICALL GetLongField) (JNIEnv *env, jobject obj, jfieldID fieldID);
// jfloat (JNICALL GetFloatField) (JNIEnv *env, jobject obj, jfieldID fieldID);
// jdouble (JNICALL GetDoubleField) (JNIEnv *env, jobject obj, jfieldID fieldID);
// void (JNICALL SetObjectField) (JNIEnv *env, jobject obj, jfieldID fieldID, jobject val);
// void (JNICALL SetBooleanField) (JNIEnv *env, jobject obj, jfieldID fieldID, jboolean val);
// void (JNICALL SetByteField) (JNIEnv *env, jobject obj, jfieldID fieldID, jbyte val);
// void (JNICALL SetCharField) (JNIEnv *env, jobject obj, jfieldID fieldID, jchar val);
// void (JNICALL SetShortField) (JNIEnv *env, jobject obj, jfieldID fieldID, jshort val);
// void (JNICALL SetIntField) (JNIEnv *env, jobject obj, jfieldID fieldID, jint val);
// void (JNICALL SetLongField) (JNIEnv *env, jobject obj, jfieldID fieldID, jlong val);
// void (JNICALL SetFloatField) (JNIEnv *env, jobject obj, jfieldID fieldID, jfloat val);
// void (JNICALL SetDoubleField) (JNIEnv *env, jobject obj, jfieldID fieldID, jdouble val);
// jmethodID (JNICALL GetStaticMethodID) (JNIEnv *env, jclass clazz, const char *name, const char *sig);
// jobject (JNICALL CallStaticObjectMethod) (JNIEnv *env, jclass clazz, jmethodID methodID, ...);
// jobject (JNICALL CallStaticObjectMethodV) (JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
// jobject (JNICALL CallStaticObjectMethodA) (JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);
// jboolean (JNICALL CallStaticBooleanMethod) (JNIEnv *env, jclass clazz, jmethodID methodID, ...);
// jboolean (JNICALL CallStaticBooleanMethodV) (JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
// jboolean (JNICALL CallStaticBooleanMethodA) (JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);
// jbyte (JNICALL CallStaticByteMethod) (JNIEnv *env, jclass clazz, jmethodID methodID, ...);
// jbyte (JNICALL CallStaticByteMethodV) (JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
// jbyte (JNICALL CallStaticByteMethodA) (JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);
// jchar (JNICALL CallStaticCharMethod) (JNIEnv *env, jclass clazz, jmethodID methodID, ...);
// jchar (JNICALL CallStaticCharMethodV) (JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
// jchar (JNICALL CallStaticCharMethodA) (JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);
// jshort (JNICALL CallStaticShortMethod) (JNIEnv *env, jclass clazz, jmethodID methodID, ...);
// jshort (JNICALL CallStaticShortMethodV) (JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
// jshort (JNICALL CallStaticShortMethodA) (JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);
// jint (JNICALL CallStaticIntMethod) (JNIEnv *env, jclass clazz, jmethodID methodID, ...);
// jint (JNICALL CallStaticIntMethodV) (JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
// jint (JNICALL CallStaticIntMethodA) (JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);
// jlong (JNICALL CallStaticLongMethod) (JNIEnv *env, jclass clazz, jmethodID methodID, ...);
// jlong (JNICALL CallStaticLongMethodV) (JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
// jlong (JNICALL CallStaticLongMethodA) (JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);
// jfloat (JNICALL CallStaticFloatMethod) (JNIEnv *env, jclass clazz, jmethodID methodID, ...);
// jfloat (JNICALL CallStaticFloatMethodV) (JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
// jfloat (JNICALL CallStaticFloatMethodA) (JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);
// jdouble (JNICALL CallStaticDoubleMethod) (JNIEnv *env, jclass clazz, jmethodID methodID, ...);
// jdouble (JNICALL CallStaticDoubleMethodV) (JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
// jdouble (JNICALL CallStaticDoubleMethodA) (JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);
// void (JNICALL CallStaticVoidMethod) (JNIEnv *env, jclass cls, jmethodID methodID, ...);
// void (JNICALL CallStaticVoidMethodV) (JNIEnv *env, jclass cls, jmethodID methodID, va_list args);
// void (JNICALL CallStaticVoidMethodA) (JNIEnv *env, jclass cls, jmethodID methodID, const jvalue * args);
// jfieldID (JNICALL GetStaticFieldID) (JNIEnv *env, jclass clazz, const char *name, const char *sig);
// jobject (JNICALL GetStaticObjectField) (JNIEnv *env, jclass clazz, jfieldID fieldID);
// jboolean (JNICALL GetStaticBooleanField) (JNIEnv *env, jclass clazz, jfieldID fieldID);
// jbyte (JNICALL GetStaticByteField) (JNIEnv *env, jclass clazz, jfieldID fieldID);
// jchar (JNICALL GetStaticCharField) (JNIEnv *env, jclass clazz, jfieldID fieldID);
// jshort (JNICALL GetStaticShortField) (JNIEnv *env, jclass clazz, jfieldID fieldID);
// jint (JNICALL GetStaticIntField) (JNIEnv *env, jclass clazz, jfieldID fieldID);
// jlong (JNICALL GetStaticLongField) (JNIEnv *env, jclass clazz, jfieldID fieldID);
// jfloat (JNICALL GetStaticFloatField) (JNIEnv *env, jclass clazz, jfieldID fieldID);
// jdouble (JNICALL GetStaticDoubleField) (JNIEnv *env, jclass clazz, jfieldID fieldID);
// void (JNICALL SetStaticObjectField) (JNIEnv *env, jclass clazz, jfieldID fieldID, jobject value);
// void (JNICALL SetStaticBooleanField) (JNIEnv *env, jclass clazz, jfieldID fieldID, jboolean value);
// void (JNICALL SetStaticByteField) (JNIEnv *env, jclass clazz, jfieldID fieldID, jbyte value);
// void (JNICALL SetStaticCharField) (JNIEnv *env, jclass clazz, jfieldID fieldID, jchar value);
// void (JNICALL SetStaticShortField) (JNIEnv *env, jclass clazz, jfieldID fieldID, jshort value);
// void (JNICALL SetStaticIntField) (JNIEnv *env, jclass clazz, jfieldID fieldID, jint value);
// void (JNICALL SetStaticLongField) (JNIEnv *env, jclass clazz, jfieldID fieldID, jlong value);
// void (JNICALL SetStaticFloatField) (JNIEnv *env, jclass clazz, jfieldID fieldID, jfloat value);
// void (JNICALL SetStaticDoubleField) (JNIEnv *env, jclass clazz, jfieldID fieldID, jdouble value);
// jstring (JNICALL NewString) (JNIEnv *env, const jchar *unicode, jsize len);
// jsize (JNICALL GetStringLength) (JNIEnv *env, jstring str);
// const jchar *(JNICALL GetStringChars) (JNIEnv *env, jstring str, jboolean *isCopy);
// void (JNICALL ReleaseStringChars) (JNIEnv *env, jstring str, const jchar *chars);
// jstring (JNICALL NewStringUTF) (JNIEnv *env, const char *utf);
// jsize (JNICALL GetStringUTFLength) (JNIEnv *env, jstring str);
// const char* (JNICALL GetStringUTFChars) (JNIEnv *env, jstring str, jboolean *isCopy);
// void (JNICALL ReleaseStringUTFChars) (JNIEnv *env, jstring str, const char* chars);
// jsize (JNICALL GetArrayLength) (JNIEnv *env, jarray array);
// jobjectArray (JNICALL NewObjectArray) (JNIEnv *env, jsize len, jclass clazz, jobject init);
// jobject (JNICALL GetObjectArrayElement) (JNIEnv *env, jobjectArray array, jsize index);
// void (JNICALL SetObjectArrayElement) (JNIEnv *env, jobjectArray array, jsize index, jobject val);
// jbooleanArray (JNICALL NewBooleanArray) (JNIEnv *env, jsize len);
// jbyteArray (JNICALL NewByteArray) (JNIEnv *env, jsize len);
// jcharArray (JNICALL NewCharArray) (JNIEnv *env, jsize len);
// jshortArray (JNICALL NewShortArray) (JNIEnv *env, jsize len);
// jintArray (JNICALL NewIntArray) (JNIEnv *env, jsize len);
// jlongArray (JNICALL NewLongArray) (JNIEnv *env, jsize len);
// jfloatArray (JNICALL NewFloatArray) (JNIEnv *env, jsize len);
// jdoubleArray (JNICALL NewDoubleArray) (JNIEnv *env, jsize len);
// jboolean * (JNICALL GetBooleanArrayElements) (JNIEnv *env, jbooleanArray array, jboolean *isCopy);
// jbyte * (JNICALL GetByteArrayElements) (JNIEnv *env, jbyteArray array, jboolean *isCopy);
// jchar * (JNICALL GetCharArrayElements) (JNIEnv *env, jcharArray array, jboolean *isCopy);
// jshort * (JNICALL GetShortArrayElements) (JNIEnv *env, jshortArray array, jboolean *isCopy);
// jint * (JNICALL GetIntArrayElements) (JNIEnv *env, jintArray array, jboolean *isCopy);
// jlong * (JNICALL GetLongArrayElements) (JNIEnv *env, jlongArray array, jboolean *isCopy);
// jfloat * (JNICALL GetFloatArrayElements) (JNIEnv *env, jfloatArray array, jboolean *isCopy);
// jdouble * (JNICALL GetDoubleArrayElements) (JNIEnv *env, jdoubleArray array, jboolean *isCopy);
// void (JNICALL ReleaseBooleanArrayElements) (JNIEnv *env, jbooleanArray array, jboolean *elems, jint mode);
// void (JNICALL ReleaseByteArrayElements) (JNIEnv *env, jbyteArray array, jbyte *elems, jint mode);
// void (JNICALL ReleaseCharArrayElements) (JNIEnv *env, jcharArray array, jchar *elems, jint mode);
// void (JNICALL ReleaseShortArrayElements) (JNIEnv *env, jshortArray array, jshort *elems, jint mode);
// void (JNICALL ReleaseIntArrayElements) (JNIEnv *env, jintArray array, jint *elems, jint mode);
// void (JNICALL ReleaseLongArrayElements) (JNIEnv *env, jlongArray array, jlong *elems, jint mode);
// void (JNICALL ReleaseFloatArrayElements) (JNIEnv *env, jfloatArray array, jfloat *elems, jint mode);
// void (JNICALL ReleaseDoubleArrayElements) (JNIEnv *env, jdoubleArray array, jdouble *elems, jint mode);
// void (JNICALL GetBooleanArrayRegion) (JNIEnv *env, jbooleanArray array, jsize start, jsize l, jboolean *buf);
// void (JNICALL GetByteArrayRegion) (JNIEnv *env, jbyteArray array, jsize start, jsize len, jbyte *buf);
// void (JNICALL GetCharArrayRegion) (JNIEnv *env, jcharArray array, jsize start, jsize len, jchar *buf);
// void (JNICALL GetShortArrayRegion) (JNIEnv *env, jshortArray array, jsize start, jsize len, jshort *buf);
// void (JNICALL GetIntArrayRegion) (JNIEnv *env, jintArray array, jsize start, jsize len, jint *buf);
// void (JNICALL GetLongArrayRegion) (JNIEnv *env, jlongArray array, jsize start, jsize len, jlong *buf);
// void (JNICALL GetFloatArrayRegion) (JNIEnv *env, jfloatArray array, jsize start, jsize len, jfloat *buf);
// void (JNICALL GetDoubleArrayRegion) (JNIEnv *env, jdoubleArray array, jsize start, jsize len, jdouble *buf);
// void (JNICALL SetBooleanArrayRegion) (JNIEnv *env, jbooleanArray array, jsize start, jsize l, const jboolean *buf);
// void (JNICALL SetByteArrayRegion) (JNIEnv *env, jbyteArray array, jsize start, jsize len, const jbyte *buf);
// void (JNICALL SetCharArrayRegion) (JNIEnv *env, jcharArray array, jsize start, jsize len, const jchar *buf);
// void (JNICALL SetShortArrayRegion) (JNIEnv *env, jshortArray array, jsize start, jsize len, const jshort *buf);
// void (JNICALL SetIntArrayRegion) (JNIEnv *env, jintArray array, jsize start, jsize len, const jint *buf);
// void (JNICALL SetLongArrayRegion) (JNIEnv *env, jlongArray array, jsize start, jsize len, const jlong *buf);
// void (JNICALL SetFloatArrayRegion) (JNIEnv *env, jfloatArray array, jsize start, jsize len, const jfloat *buf);
// void (JNICALL SetDoubleArrayRegion) (JNIEnv *env, jdoubleArray array, jsize start, jsize len, const jdouble *buf);
// jint (JNICALL RegisterNatives) (JNIEnv *env, jclass clazz, const JNINativeMethod *methods, jint nMethods);
// jint (JNICALL UnregisterNatives) (JNIEnv *env, jclass clazz);
// jint (JNICALL MonitorEnter) (JNIEnv *env, jobject obj);
// jint (JNICALL MonitorExit) (JNIEnv *env, jobject obj);
// jint (JNICALL GetJavaVM) (JNIEnv *env, JavaVM **vm);
// void (JNICALL GetStringRegion) (JNIEnv *env, jstring str, jsize start, jsize len, jchar *buf);
// void (JNICALL GetStringUTFRegion) (JNIEnv *env, jstring str, jsize start, jsize len, char *buf);
// void * (JNICALL GetPrimitiveArrayCritical) (JNIEnv *env, jarray array, jboolean *isCopy);
// void (JNICALL ReleasePrimitiveArrayCritical) (JNIEnv *env, jarray array, void *carray, jint mode);
// const jchar * (JNICALL GetStringCritical) (JNIEnv *env, jstring string, jboolean *isCopy);
// void (JNICALL ReleaseStringCritical) (JNIEnv *env, jstring string, const jchar *cstring);
// jweak (JNICALL NewWeakGlobalRef) (JNIEnv *env, jobject obj);
// void (JNICALL DeleteWeakGlobalRef) (JNIEnv *env, jweak ref);
// jboolean (JNICALL ExceptionCheck) (JNIEnv *env);
// jobject (JNICALL NewDirectByteBuffer) (JNIEnv* env, void* address, jlong capacity);
// void* (JNICALL GetDirectBufferAddress) (JNIEnv* env, jobject buf);
// jlong (JNICALL GetDirectBufferCapacity) (JNIEnv* env, jobject buf);
// jobjectRefType (JNICALL GetObjectRefType) (JNIEnv* env, jobject obj);
//
// // https://github.com/golang/go/wiki/cgo#function-variables
// static inline JNIEnv createJNIEnv() {
//   struct JNINativeInterface_ * myEnv = malloc(sizeof(struct JNINativeInterface_));
//   myEnv->GetVersion = GetVersion;
//   return myEnv;
// }
// static inline void test(JNIEnv *env) {
//   printf("%d", (*env)->GetVersion(env));
// }
import "C"
import (
	"fmt"
	"os"
	"unsafe"
)

var version int32 = 100

//export GetVersion
func GetVersion(env *C.JNIEnv) C.jint {
	println("hello JNI")
	return C.int(*(*int32)((*env).reserved0))
}

func Test() {
	myEnv := C.createJNIEnv()
	fmt.Printf("%T\n", myEnv)
	fmt.Printf("%v\n", myEnv.reserved0)
	fmt.Printf("%T\n", myEnv.GetVersion)
	fmt.Printf("%v\n", myEnv.GetVersion)
	myEnv.reserved0 = unsafe.Pointer(&version)
	C.test(&myEnv)
	os.Exit(0)
}
