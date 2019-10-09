package jni

// #cgo CFLAGS: -I/Users/zxh/.sdkman/candidates/java/8.0.222-zulu/include
// #cgo CFLAGS: -I/Users/zxh/.sdkman/candidates/java/8.0.222-zulu/include/darwin
// #include <stdio.h>
// #include <stdlib.h>
// #include "jni.h"
//
// jint GetVersion(JNIEnv *env);
// jclass DefineClass(JNIEnv *env, char *name, jobject loader, jbyte *buf, jsize len);
// jclass FindClass(JNIEnv *env, char *name);
// jmethodID FromReflectedMethod(JNIEnv *env, jobject method);
// jfieldID FromReflectedField(JNIEnv *env, jobject field);
// jobject ToReflectedMethod(JNIEnv *env, jclass cls, jmethodID methodID, jboolean isStatic);
// jclass GetSuperclass(JNIEnv *env, jclass sub);
// jboolean IsAssignableFrom(JNIEnv *env, jclass sub, jclass sup);
// jobject ToReflectedField(JNIEnv *env, jclass cls, jfieldID fieldID, jboolean isStatic);
// jint Throw(JNIEnv *env, jthrowable obj);
// jint ThrowNew(JNIEnv *env, jclass clazz, char *msg);
// jthrowable ExceptionOccurred(JNIEnv *env);
// void ExceptionDescribe(JNIEnv *env);
// void ExceptionClear(JNIEnv *env);
// void FatalError(JNIEnv *env, char *msg);
// jint PushLocalFrame(JNIEnv *env, jint capacity);
// jobject PopLocalFrame(JNIEnv *env, jobject result);
// jobject NewGlobalRef(JNIEnv *env, jobject lobj);
// void DeleteGlobalRef(JNIEnv *env, jobject gref);
// void DeleteLocalRef(JNIEnv *env, jobject obj);
// jboolean IsSameObject(JNIEnv *env, jobject obj1, jobject obj2);
// jobject NewLocalRef(JNIEnv *env, jobject ref);
// jint EnsureLocalCapacity(JNIEnv *env, jint capacity);
// jobject AllocObject(JNIEnv *env, jclass clazz);
// jobject NewObject(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
// jobject NewObjectV(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
// jobject NewObjectA(JNIEnv *env, jclass clazz, jmethodID methodID, jvalue *args);
// jclass GetObjectClass(JNIEnv *env, jobject obj);
// jboolean IsInstanceOf(JNIEnv *env, jobject obj, jclass clazz);
// jmethodID GetMethodID(JNIEnv *env, jclass clazz, char *name, char *sig);
// jobject CallObjectMethod(JNIEnv *env, jobject obj, jmethodID methodID, ...);
// jobject CallObjectMethodV(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
// jobject CallObjectMethodA(JNIEnv *env, jobject obj, jmethodID methodID, jvalue * args);
// jboolean CallBooleanMethod(JNIEnv *env, jobject obj, jmethodID methodID, ...);
// jboolean CallBooleanMethodV(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
// jboolean CallBooleanMethodA(JNIEnv *env, jobject obj, jmethodID methodID, jvalue * args);
// jbyte CallByteMethod(JNIEnv *env, jobject obj, jmethodID methodID, ...);
// jbyte CallByteMethodV(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
// jbyte CallByteMethodA(JNIEnv *env, jobject obj, jmethodID methodID, jvalue *args);
// jchar CallCharMethod(JNIEnv *env, jobject obj, jmethodID methodID, ...);
// jchar CallCharMethodV(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
// jchar CallCharMethodA(JNIEnv *env, jobject obj, jmethodID methodID, jvalue *args);
// jshort CallShortMethod(JNIEnv *env, jobject obj, jmethodID methodID, ...);
// jshort CallShortMethodV(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
// jshort CallShortMethodA(JNIEnv *env, jobject obj, jmethodID methodID, jvalue *args);
// jint CallIntMethod(JNIEnv *env, jobject obj, jmethodID methodID, ...);
// jint CallIntMethodV(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
// jint CallIntMethodA(JNIEnv *env, jobject obj, jmethodID methodID, jvalue *args);
// jlong CallLongMethod(JNIEnv *env, jobject obj, jmethodID methodID, ...);
// jlong CallLongMethodV(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
// jlong CallLongMethodA(JNIEnv *env, jobject obj, jmethodID methodID, jvalue *args);
// jfloat CallFloatMethod(JNIEnv *env, jobject obj, jmethodID methodID, ...);
// jfloat CallFloatMethodV(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
// jfloat CallFloatMethodA(JNIEnv *env, jobject obj, jmethodID methodID, jvalue *args);
// jdouble CallDoubleMethod(JNIEnv *env, jobject obj, jmethodID methodID, ...);
// jdouble CallDoubleMethodV(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
// jdouble CallDoubleMethodA(JNIEnv *env, jobject obj, jmethodID methodID, jvalue *args);
// void CallVoidMethod(JNIEnv *env, jobject obj, jmethodID methodID, ...);
// void CallVoidMethodV(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
// void CallVoidMethodA(JNIEnv *env, jobject obj, jmethodID methodID, jvalue * args);
// jobject CallNonvirtualObjectMethod(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
// jobject CallNonvirtualObjectMethodV(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args);
// jobject CallNonvirtualObjectMethodA(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, jvalue * args);
// jboolean CallNonvirtualBooleanMethod(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
// jboolean CallNonvirtualBooleanMethodV(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args);
// jboolean CallNonvirtualBooleanMethodA(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, jvalue * args);
// jbyte CallNonvirtualByteMethod(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
// jbyte CallNonvirtualByteMethodV(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args);
// jbyte CallNonvirtualByteMethodA(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, jvalue *args);
// jchar CallNonvirtualCharMethod(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
// jchar CallNonvirtualCharMethodV(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args);
// jchar CallNonvirtualCharMethodA(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, jvalue *args);
// jshort CallNonvirtualShortMethod(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
// jshort CallNonvirtualShortMethodV(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args);
// jshort CallNonvirtualShortMethodA(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, jvalue *args);
// jint CallNonvirtualIntMethod(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
// jint CallNonvirtualIntMethodV(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args);
// jint CallNonvirtualIntMethodA(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, jvalue *args);
// jlong CallNonvirtualLongMethod(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
// jlong CallNonvirtualLongMethodV(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args);
// jlong CallNonvirtualLongMethodA(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, jvalue *args);
// jfloat CallNonvirtualFloatMethod(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
// jfloat CallNonvirtualFloatMethodV(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args);
// jfloat CallNonvirtualFloatMethodA(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, jvalue *args);
// jdouble CallNonvirtualDoubleMethod(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
// jdouble CallNonvirtualDoubleMethodV(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args);
// jdouble CallNonvirtualDoubleMethodA(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, jvalue *args);
// void CallNonvirtualVoidMethod(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
// void CallNonvirtualVoidMethodV(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args);
// void CallNonvirtualVoidMethodA(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, jvalue * args);
// jfieldID GetFieldID(JNIEnv *env, jclass clazz, char *name, char *sig);
// jobject GetObjectField(JNIEnv *env, jobject obj, jfieldID fieldID);
// jboolean GetBooleanField(JNIEnv *env, jobject obj, jfieldID fieldID);
// jbyte GetByteField(JNIEnv *env, jobject obj, jfieldID fieldID);
// jchar GetCharField(JNIEnv *env, jobject obj, jfieldID fieldID);
// jshort GetShortField(JNIEnv *env, jobject obj, jfieldID fieldID);
// jint GetIntField(JNIEnv *env, jobject obj, jfieldID fieldID);
// jlong GetLongField(JNIEnv *env, jobject obj, jfieldID fieldID);
// jfloat GetFloatField(JNIEnv *env, jobject obj, jfieldID fieldID);
// jdouble GetDoubleField(JNIEnv *env, jobject obj, jfieldID fieldID);
// void SetObjectField(JNIEnv *env, jobject obj, jfieldID fieldID, jobject val);
// void SetBooleanField(JNIEnv *env, jobject obj, jfieldID fieldID, jboolean val);
// void SetByteField(JNIEnv *env, jobject obj, jfieldID fieldID, jbyte val);
// void SetCharField(JNIEnv *env, jobject obj, jfieldID fieldID, jchar val);
// void SetShortField(JNIEnv *env, jobject obj, jfieldID fieldID, jshort val);
// void SetIntField(JNIEnv *env, jobject obj, jfieldID fieldID, jint val);
// void SetLongField(JNIEnv *env, jobject obj, jfieldID fieldID, jlong val);
// void SetFloatField(JNIEnv *env, jobject obj, jfieldID fieldID, jfloat val);
// void SetDoubleField(JNIEnv *env, jobject obj, jfieldID fieldID, jdouble val);
// jmethodID GetStaticMethodID(JNIEnv *env, jclass clazz, char *name, char *sig);
// jobject CallStaticObjectMethod(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
// jobject CallStaticObjectMethodV(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
// jobject CallStaticObjectMethodA(JNIEnv *env, jclass clazz, jmethodID methodID, jvalue *args);
// jboolean CallStaticBooleanMethod(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
// jboolean CallStaticBooleanMethodV(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
// jboolean CallStaticBooleanMethodA(JNIEnv *env, jclass clazz, jmethodID methodID, jvalue *args);
// jbyte CallStaticByteMethod(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
// jbyte CallStaticByteMethodV(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
// jbyte CallStaticByteMethodA(JNIEnv *env, jclass clazz, jmethodID methodID, jvalue *args);
// jchar CallStaticCharMethod(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
// jchar CallStaticCharMethodV(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
// jchar CallStaticCharMethodA(JNIEnv *env, jclass clazz, jmethodID methodID, jvalue *args);
// jshort CallStaticShortMethod(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
// jshort CallStaticShortMethodV(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
// jshort CallStaticShortMethodA(JNIEnv *env, jclass clazz, jmethodID methodID, jvalue *args);
// jint CallStaticIntMethod(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
// jint CallStaticIntMethodV(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
// jint CallStaticIntMethodA(JNIEnv *env, jclass clazz, jmethodID methodID, jvalue *args);
// jlong CallStaticLongMethod(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
// jlong CallStaticLongMethodV(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
// jlong CallStaticLongMethodA(JNIEnv *env, jclass clazz, jmethodID methodID, jvalue *args);
// jfloat CallStaticFloatMethod(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
// jfloat CallStaticFloatMethodV(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
// jfloat CallStaticFloatMethodA(JNIEnv *env, jclass clazz, jmethodID methodID, jvalue *args);
// jdouble CallStaticDoubleMethod(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
// jdouble CallStaticDoubleMethodV(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
// jdouble CallStaticDoubleMethodA(JNIEnv *env, jclass clazz, jmethodID methodID, jvalue *args);
// void CallStaticVoidMethod(JNIEnv *env, jclass cls, jmethodID methodID, ...);
// void CallStaticVoidMethodV(JNIEnv *env, jclass cls, jmethodID methodID, va_list args);
// void CallStaticVoidMethodA(JNIEnv *env, jclass cls, jmethodID methodID, jvalue * args);
// jfieldID GetStaticFieldID(JNIEnv *env, jclass clazz, char *name, char *sig);
// jobject GetStaticObjectField(JNIEnv *env, jclass clazz, jfieldID fieldID);
// jboolean GetStaticBooleanField(JNIEnv *env, jclass clazz, jfieldID fieldID);
// jbyte GetStaticByteField(JNIEnv *env, jclass clazz, jfieldID fieldID);
// jchar GetStaticCharField(JNIEnv *env, jclass clazz, jfieldID fieldID);
// jshort GetStaticShortField(JNIEnv *env, jclass clazz, jfieldID fieldID);
// jint GetStaticIntField(JNIEnv *env, jclass clazz, jfieldID fieldID);
// jlong GetStaticLongField(JNIEnv *env, jclass clazz, jfieldID fieldID);
// jfloat GetStaticFloatField(JNIEnv *env, jclass clazz, jfieldID fieldID);
// jdouble GetStaticDoubleField(JNIEnv *env, jclass clazz, jfieldID fieldID);
// void SetStaticObjectField(JNIEnv *env, jclass clazz, jfieldID fieldID, jobject value);
// void SetStaticBooleanField(JNIEnv *env, jclass clazz, jfieldID fieldID, jboolean value);
// void SetStaticByteField(JNIEnv *env, jclass clazz, jfieldID fieldID, jbyte value);
// void SetStaticCharField(JNIEnv *env, jclass clazz, jfieldID fieldID, jchar value);
// void SetStaticShortField(JNIEnv *env, jclass clazz, jfieldID fieldID, jshort value);
// void SetStaticIntField(JNIEnv *env, jclass clazz, jfieldID fieldID, jint value);
// void SetStaticLongField(JNIEnv *env, jclass clazz, jfieldID fieldID, jlong value);
// void SetStaticFloatField(JNIEnv *env, jclass clazz, jfieldID fieldID, jfloat value);
// void SetStaticDoubleField(JNIEnv *env, jclass clazz, jfieldID fieldID, jdouble value);
// jstring NewString(JNIEnv *env, jchar *unicode, jsize len);
// jsize GetStringLength(JNIEnv *env, jstring str);
// jchar* GetStringChars(JNIEnv *env, jstring str, jboolean *isCopy);
// void ReleaseStringChars(JNIEnv *env, jstring str, jchar *chars);
// jstring NewStringUTF(JNIEnv *env, char *utf);
// jsize GetStringUTFLength(JNIEnv *env, jstring str);
// char* GetStringUTFChars(JNIEnv *env, jstring str, jboolean *isCopy);
// void ReleaseStringUTFChars(JNIEnv *env, jstring str, char* chars);
// jsize GetArrayLength(JNIEnv *env, jarray array);
// jobjectArray NewObjectArray(JNIEnv *env, jsize len, jclass clazz, jobject init);
// jobject GetObjectArrayElement(JNIEnv *env, jobjectArray array, jsize index);
// void SetObjectArrayElement(JNIEnv *env, jobjectArray array, jsize index, jobject val);
// jbooleanArray NewBooleanArray(JNIEnv *env, jsize len);
// jbyteArray NewByteArray(JNIEnv *env, jsize len);
// jcharArray NewCharArray(JNIEnv *env, jsize len);
// jshortArray NewShortArray(JNIEnv *env, jsize len);
// jintArray NewIntArray(JNIEnv *env, jsize len);
// jlongArray NewLongArray(JNIEnv *env, jsize len);
// jfloatArray NewFloatArray(JNIEnv *env, jsize len);
// jdoubleArray NewDoubleArray(JNIEnv *env, jsize len);
// jboolean * GetBooleanArrayElements(JNIEnv *env, jbooleanArray array, jboolean *isCopy);
// jbyte * GetByteArrayElements(JNIEnv *env, jbyteArray array, jboolean *isCopy);
// jchar * GetCharArrayElements(JNIEnv *env, jcharArray array, jboolean *isCopy);
// jshort * GetShortArrayElements(JNIEnv *env, jshortArray array, jboolean *isCopy);
// jint * GetIntArrayElements(JNIEnv *env, jintArray array, jboolean *isCopy);
// jlong * GetLongArrayElements(JNIEnv *env, jlongArray array, jboolean *isCopy);
// jfloat * GetFloatArrayElements(JNIEnv *env, jfloatArray array, jboolean *isCopy);
// jdouble * GetDoubleArrayElements(JNIEnv *env, jdoubleArray array, jboolean *isCopy);
// void ReleaseBooleanArrayElements(JNIEnv *env, jbooleanArray array, jboolean *elems, jint mode);
// void ReleaseByteArrayElements(JNIEnv *env, jbyteArray array, jbyte *elems, jint mode);
// void ReleaseCharArrayElements(JNIEnv *env, jcharArray array, jchar *elems, jint mode);
// void ReleaseShortArrayElements(JNIEnv *env, jshortArray array, jshort *elems, jint mode);
// void ReleaseIntArrayElements(JNIEnv *env, jintArray array, jint *elems, jint mode);
// void ReleaseLongArrayElements(JNIEnv *env, jlongArray array, jlong *elems, jint mode);
// void ReleaseFloatArrayElements(JNIEnv *env, jfloatArray array, jfloat *elems, jint mode);
// void ReleaseDoubleArrayElements(JNIEnv *env, jdoubleArray array, jdouble *elems, jint mode);
// void GetBooleanArrayRegion(JNIEnv *env, jbooleanArray array, jsize start, jsize l, jboolean *buf);
// void GetByteArrayRegion(JNIEnv *env, jbyteArray array, jsize start, jsize len, jbyte *buf);
// void GetCharArrayRegion(JNIEnv *env, jcharArray array, jsize start, jsize len, jchar *buf);
// void GetShortArrayRegion(JNIEnv *env, jshortArray array, jsize start, jsize len, jshort *buf);
// void GetIntArrayRegion(JNIEnv *env, jintArray array, jsize start, jsize len, jint *buf);
// void GetLongArrayRegion(JNIEnv *env, jlongArray array, jsize start, jsize len, jlong *buf);
// void GetFloatArrayRegion(JNIEnv *env, jfloatArray array, jsize start, jsize len, jfloat *buf);
// void GetDoubleArrayRegion(JNIEnv *env, jdoubleArray array, jsize start, jsize len, jdouble *buf);
// void SetBooleanArrayRegion(JNIEnv *env, jbooleanArray array, jsize start, jsize l, jboolean *buf);
// void SetByteArrayRegion(JNIEnv *env, jbyteArray array, jsize start, jsize len, jbyte *buf);
// void SetCharArrayRegion(JNIEnv *env, jcharArray array, jsize start, jsize len, jchar *buf);
// void SetShortArrayRegion(JNIEnv *env, jshortArray array, jsize start, jsize len, jshort *buf);
// void SetIntArrayRegion(JNIEnv *env, jintArray array, jsize start, jsize len, jint *buf);
// void SetLongArrayRegion(JNIEnv *env, jlongArray array, jsize start, jsize len, jlong *buf);
// void SetFloatArrayRegion(JNIEnv *env, jfloatArray array, jsize start, jsize len, jfloat *buf);
// void SetDoubleArrayRegion(JNIEnv *env, jdoubleArray array, jsize start, jsize len, jdouble *buf);
// jint RegisterNatives(JNIEnv *env, jclass clazz, JNINativeMethod *methods, jint nMethods);
// jint UnregisterNatives(JNIEnv *env, jclass clazz);
// jint MonitorEnter(JNIEnv *env, jobject obj);
// jint MonitorExit(JNIEnv *env, jobject obj);
// jint GetJavaVM(JNIEnv *env, JavaVM **vm);
// void GetStringRegion(JNIEnv *env, jstring str, jsize start, jsize len, jchar *buf);
// void GetStringUTFRegion(JNIEnv *env, jstring str, jsize start, jsize len, char *buf);
// void * GetPrimitiveArrayCritical(JNIEnv *env, jarray array, jboolean *isCopy);
// void ReleasePrimitiveArrayCritical(JNIEnv *env, jarray array, void *carray, jint mode);
// jchar * GetStringCritical(JNIEnv *env, jstring string, jboolean *isCopy);
// void ReleaseStringCritical(JNIEnv *env, jstring string, jchar *cstring);
// jweak NewWeakGlobalRef(JNIEnv *env, jobject obj);
// void DeleteWeakGlobalRef(JNIEnv *env, jweak ref);
// jboolean ExceptionCheck(JNIEnv *env);
// jobject NewDirectByteBuffer(JNIEnv* env, void* address, jlong capacity);
// void* GetDirectBufferAddress(JNIEnv* env, jobject buf);
// jlong GetDirectBufferCapacity(JNIEnv* env, jobject buf);
// jobjectRefType GetObjectRefType(JNIEnv* env, jobject obj);
//
// // https://github.com/golang/go/wiki/cgo#function-variables
// static inline JNIEnv createJNIEnv() {
//   struct JNINativeInterface_ * myEnv = malloc(sizeof(struct JNINativeInterface_));
//   myEnv->GetVersion = GetVersion;
//   /*myEnv->DefineClass = DefineClass;
//   myEnv->FindClass = FindClass;
//   myEnv->FromReflectedMethod = FromReflectedMethod;
//   myEnv->FromReflectedField = FromReflectedField;
//   myEnv->ToReflectedMethod = ToReflectedMethod;
//   myEnv->GetSuperclass = GetSuperclass;
//   myEnv->IsAssignableFrom = IsAssignableFrom;
//   myEnv->ToReflectedField = ToReflectedField;
//   myEnv->Throw = Throw;
//   myEnv->ThrowNew = ThrowNew;
//   myEnv->ExceptionOccurred = ExceptionOccurred;
//   myEnv->ExceptionDescribe = ExceptionDescribe;
//   myEnv->ExceptionClear = ExceptionClear;
//   myEnv->FatalError = FatalError;
//   myEnv->PushLocalFrame = PushLocalFrame;
//   myEnv->PopLocalFrame = PopLocalFrame;
//   myEnv->NewGlobalRef = NewGlobalRef;
//   myEnv->DeleteGlobalRef = DeleteGlobalRef;
//   myEnv->DeleteLocalRef = DeleteLocalRef;
//   myEnv->IsSameObject = IsSameObject;
//   myEnv->NewLocalRef = NewLocalRef;
//   myEnv->EnsureLocalCapacity = EnsureLocalCapacity;
//   myEnv->AllocObject = AllocObject;
//   //myEnv->NewObject = NewObject;
//   //myEnv->NewObjectV = NewObjectV;
//   myEnv->NewObjectA = NewObjectA;
//   myEnv->GetObjectClass = GetObjectClass;
//   myEnv->IsInstanceOf = IsInstanceOf;
//   myEnv->GetMethodID = GetMethodID;
//   //myEnv->CallObjectMethod = CallObjectMethod;
//   //myEnv->CallObjectMethodV = CallObjectMethodV;
//   myEnv->CallObjectMethodA = CallObjectMethodA;
//   //myEnv->CallBooleanMethod = CallBooleanMethod;
//   //myEnv->CallBooleanMethodV = CallBooleanMethodV;
//   myEnv->CallBooleanMethodA = CallBooleanMethodA;
//   //myEnv->CallByteMethod = CallByteMethod;
//   //myEnv->CallByteMethodV = CallByteMethodV;
//   myEnv->CallByteMethodA = CallByteMethodA;
//   //myEnv->CallCharMethod = CallCharMethod;
//   //myEnv->CallCharMethodV = CallCharMethodV;
//   myEnv->CallCharMethodA = CallCharMethodA;
//   //myEnv->CallShortMethod = CallShortMethod;
//   //myEnv->CallShortMethodV = CallShortMethodV;
//   myEnv->CallShortMethodA = CallShortMethodA;
//   //myEnv->CallIntMethod = CallIntMethod;
//   //myEnv->CallIntMethodV = CallIntMethodV;
//   myEnv->CallIntMethodA = CallIntMethodA;
//   //myEnv->CallLongMethod = CallLongMethod;
//   //myEnv->CallLongMethodV = CallLongMethodV;
//   myEnv->CallLongMethodA = CallLongMethodA;
//   //myEnv->CallFloatMethod = CallFloatMethod;
//   //myEnv->CallFloatMethodV = CallFloatMethodV;
//   myEnv->CallFloatMethodA = CallFloatMethodA;
//   //myEnv->CallDoubleMethod = CallDoubleMethod;
//   //myEnv->CallDoubleMethodV = CallDoubleMethodV;
//   myEnv->CallDoubleMethodA = CallDoubleMethodA;
//   //myEnv->CallVoidMethod = CallVoidMethod;
//   //myEnv->CallVoidMethodV = CallVoidMethodV;
//   myEnv->CallVoidMethodA = CallVoidMethodA;
//   //myEnv->CallNonvirtualObjectMethod = CallNonvirtualObjectMethod;
//   //myEnv->CallNonvirtualObjectMethodV = CallNonvirtualObjectMethodV;
//   myEnv->CallNonvirtualObjectMethodA = CallNonvirtualObjectMethodA;
//   //myEnv->CallNonvirtualBooleanMethod = CallNonvirtualBooleanMethod;
//   //myEnv->CallNonvirtualBooleanMethodV = CallNonvirtualBooleanMethodV;
//   myEnv->CallNonvirtualBooleanMethodA = CallNonvirtualBooleanMethodA;
//   //myEnv->CallNonvirtualByteMethod = CallNonvirtualByteMethod;
//   //myEnv->CallNonvirtualByteMethodV = CallNonvirtualByteMethodV;
//   myEnv->CallNonvirtualByteMethodA = CallNonvirtualByteMethodA;
//   //myEnv->CallNonvirtualCharMethod = CallNonvirtualCharMethod;
//   //myEnv->CallNonvirtualCharMethodV = CallNonvirtualCharMethodV;
//   myEnv->CallNonvirtualCharMethodA = CallNonvirtualCharMethodA;
//   //myEnv->CallNonvirtualShortMethod = CallNonvirtualShortMethod;
//   //myEnv->CallNonvirtualShortMethodV = CallNonvirtualShortMethodV;
//   myEnv->CallNonvirtualShortMethodA = CallNonvirtualShortMethodA;
//   //myEnv->CallNonvirtualIntMethod = CallNonvirtualIntMethod;
//   //myEnv->CallNonvirtualIntMethodV = CallNonvirtualIntMethodV;
//   myEnv->CallNonvirtualIntMethodA = CallNonvirtualIntMethodA;
//   //myEnv->CallNonvirtualLongMethod = CallNonvirtualLongMethod;
//   //myEnv->CallNonvirtualLongMethodV = CallNonvirtualLongMethodV;
//   myEnv->CallNonvirtualLongMethodA = CallNonvirtualLongMethodA;
//   //myEnv->CallNonvirtualFloatMethod = CallNonvirtualFloatMethod;
//   //myEnv->CallNonvirtualFloatMethodV = CallNonvirtualFloatMethodV;
//   myEnv->CallNonvirtualFloatMethodA = CallNonvirtualFloatMethodA;
//   //myEnv->CallNonvirtualDoubleMethod = CallNonvirtualDoubleMethod;
//   //myEnv->CallNonvirtualDoubleMethodV = CallNonvirtualDoubleMethodV;
//   myEnv->CallNonvirtualDoubleMethodA = CallNonvirtualDoubleMethodA;
//   //myEnv->CallNonvirtualVoidMethod = CallNonvirtualVoidMethod;
//   //myEnv->CallNonvirtualVoidMethodV = CallNonvirtualVoidMethodV;
//   myEnv->CallNonvirtualVoidMethodA = CallNonvirtualVoidMethodA;
//   myEnv->GetFieldID = GetFieldID;
//   myEnv->GetObjectField = GetObjectField;
//   myEnv->GetBooleanField = GetBooleanField;
//   myEnv->GetByteField = GetByteField;
//   myEnv->GetCharField = GetCharField;
//   myEnv->GetShortField = GetShortField;
//   myEnv->GetIntField = GetIntField;
//   myEnv->GetLongField = GetLongField;
//   myEnv->GetFloatField = GetFloatField;
//   myEnv->GetDoubleField = GetDoubleField;
//   myEnv->SetObjectField = SetObjectField;
//   myEnv->SetBooleanField = SetBooleanField;
//   myEnv->SetByteField = SetByteField;
//   myEnv->SetCharField = SetCharField;
//   myEnv->SetShortField = SetShortField;
//   myEnv->SetIntField = SetIntField;
//   myEnv->SetLongField = SetLongField;
//   myEnv->SetFloatField = SetFloatField;
//   myEnv->SetDoubleField = SetDoubleField;
//   myEnv->GetStaticMethodID = GetStaticMethodID;
//   //myEnv->CallStaticObjectMethod = CallStaticObjectMethod;
//   //myEnv->CallStaticObjectMethodV = CallStaticObjectMethodV;
//   myEnv->CallStaticObjectMethodA = CallStaticObjectMethodA;
//   //myEnv->CallStaticBooleanMethod = CallStaticBooleanMethod;
//   //myEnv->CallStaticBooleanMethodV = CallStaticBooleanMethodV;
//   myEnv->CallStaticBooleanMethodA = CallStaticBooleanMethodA;
//   //myEnv->CallStaticByteMethod = CallStaticByteMethod;
//   //myEnv->CallStaticByteMethodV = CallStaticByteMethodV;
//   myEnv->CallStaticByteMethodA = CallStaticByteMethodA;
//   //myEnv->CallStaticCharMethod = CallStaticCharMethod;
//   //myEnv->CallStaticCharMethodV = CallStaticCharMethodV;
//   myEnv->CallStaticCharMethodA = CallStaticCharMethodA;
//   //myEnv->CallStaticShortMethod = CallStaticShortMethod;
//   //myEnv->CallStaticShortMethodV = CallStaticShortMethodV;
//   myEnv->CallStaticShortMethodA = CallStaticShortMethodA;
//   //myEnv->CallStaticIntMethod = CallStaticIntMethod;
//   //myEnv->CallStaticIntMethodV = CallStaticIntMethodV;
//   myEnv->CallStaticIntMethodA = CallStaticIntMethodA;
//   //myEnv->CallStaticLongMethod = CallStaticLongMethod;
//   //myEnv->CallStaticLongMethodV = CallStaticLongMethodV;
//   myEnv->CallStaticLongMethodA = CallStaticLongMethodA;
//   //myEnv->CallStaticFloatMethod = CallStaticFloatMethod;
//   //myEnv->CallStaticFloatMethodV = CallStaticFloatMethodV;
//   myEnv->CallStaticFloatMethodA = CallStaticFloatMethodA;
//   //myEnv->CallStaticDoubleMethod = CallStaticDoubleMethod;
//   //myEnv->CallStaticDoubleMethodV = CallStaticDoubleMethodV;
//   myEnv->CallStaticDoubleMethodA = CallStaticDoubleMethodA;
//   //myEnv->CallStaticVoidMethod = CallStaticVoidMethod;
//   //myEnv->CallStaticVoidMethodV = CallStaticVoidMethodV;
//   myEnv->CallStaticVoidMethodA = CallStaticVoidMethodA;
//   myEnv->GetStaticFieldID = GetStaticFieldID;
//   myEnv->GetStaticObjectField = GetStaticObjectField;
//   myEnv->GetStaticBooleanField = GetStaticBooleanField;
//   myEnv->GetStaticByteField = GetStaticByteField;
//   myEnv->GetStaticCharField = GetStaticCharField;
//   myEnv->GetStaticShortField = GetStaticShortField;
//   myEnv->GetStaticIntField = GetStaticIntField;
//   myEnv->GetStaticLongField = GetStaticLongField;
//   myEnv->GetStaticFloatField = GetStaticFloatField;
//   myEnv->GetStaticDoubleField = GetStaticDoubleField;
//   myEnv->SetStaticObjectField = SetStaticObjectField;
//   myEnv->SetStaticBooleanField = SetStaticBooleanField;
//   myEnv->SetStaticByteField = SetStaticByteField;
//   myEnv->SetStaticCharField = SetStaticCharField;
//   myEnv->SetStaticShortField = SetStaticShortField;
//   myEnv->SetStaticIntField = SetStaticIntField;
//   myEnv->SetStaticLongField = SetStaticLongField;
//   myEnv->SetStaticFloatField = SetStaticFloatField;
//   myEnv->SetStaticDoubleField = SetStaticDoubleField;
//   myEnv->NewString = NewString;
//   myEnv->GetStringLength = GetStringLength;
//   myEnv->GetStringChars = GetStringChars;
//   myEnv->ReleaseStringChars = ReleaseStringChars;
//   myEnv->NewStringUTF = NewStringUTF;
//   myEnv->GetStringUTFLength = GetStringUTFLength;
//   myEnv->GetStringUTFChars = GetStringUTFChars;
//   myEnv->ReleaseStringUTFChars = ReleaseStringUTFChars;
//   myEnv->GetArrayLength = GetArrayLength;*/
//   myEnv->NewObjectArray = NewObjectArray;
//   myEnv->GetObjectArrayElement = GetObjectArrayElement;
//   myEnv->SetObjectArrayElement = SetObjectArrayElement;
//   myEnv->NewBooleanArray = NewBooleanArray;
//   myEnv->NewByteArray = NewByteArray;
//   myEnv->NewCharArray = NewCharArray;
//   myEnv->NewShortArray = NewShortArray;
//   myEnv->NewIntArray = NewIntArray;
//   myEnv->NewLongArray = NewLongArray;
//   myEnv->NewFloatArray = NewFloatArray;
//   myEnv->NewDoubleArray = NewDoubleArray;
//   myEnv->GetBooleanArrayElements = GetBooleanArrayElements;
//   myEnv->GetByteArrayElements = GetByteArrayElements;
//   myEnv->GetCharArrayElements = GetCharArrayElements;
//   myEnv->GetShortArrayElements = GetShortArrayElements;
//   myEnv->GetIntArrayElements = GetIntArrayElements;
//   myEnv->GetLongArrayElements = GetLongArrayElements;
//   myEnv->GetFloatArrayElements = GetFloatArrayElements;
//   myEnv->GetDoubleArrayElements = GetDoubleArrayElements;
//   myEnv->ReleaseBooleanArrayElements = ReleaseBooleanArrayElements;
//   myEnv->ReleaseByteArrayElements = ReleaseByteArrayElements;
//   myEnv->ReleaseCharArrayElements = ReleaseCharArrayElements;
//   myEnv->ReleaseShortArrayElements = ReleaseShortArrayElements;
//   myEnv->ReleaseIntArrayElements = ReleaseIntArrayElements;
//   myEnv->ReleaseLongArrayElements = ReleaseLongArrayElements;
//   myEnv->ReleaseFloatArrayElements = ReleaseFloatArrayElements;
//   myEnv->ReleaseDoubleArrayElements = ReleaseDoubleArrayElements;
//   myEnv->GetBooleanArrayRegion = GetBooleanArrayRegion;
//   myEnv->GetByteArrayRegion = GetByteArrayRegion;
//   myEnv->GetCharArrayRegion = GetCharArrayRegion;
//   myEnv->GetShortArrayRegion = GetShortArrayRegion;
//   myEnv->GetIntArrayRegion = GetIntArrayRegion;
//   myEnv->GetLongArrayRegion = GetLongArrayRegion;
//   myEnv->GetFloatArrayRegion = GetFloatArrayRegion;
//   myEnv->GetDoubleArrayRegion = GetDoubleArrayRegion;
//   myEnv->SetBooleanArrayRegion = SetBooleanArrayRegion;
//   myEnv->SetByteArrayRegion = SetByteArrayRegion;
//   myEnv->SetCharArrayRegion = SetCharArrayRegion;
//   myEnv->SetShortArrayRegion = SetShortArrayRegion;
//   myEnv->SetIntArrayRegion = SetIntArrayRegion;
//   myEnv->SetLongArrayRegion = SetLongArrayRegion;
//   myEnv->SetFloatArrayRegion = SetFloatArrayRegion;
//   myEnv->SetDoubleArrayRegion = SetDoubleArrayRegion;
//   myEnv->RegisterNatives = RegisterNatives;
//   myEnv->UnregisterNatives = UnregisterNatives;
//   myEnv->MonitorEnter = MonitorEnter;
//   myEnv->MonitorExit = MonitorExit;
//   myEnv->GetJavaVM = GetJavaVM;
//   myEnv->GetStringRegion = GetStringRegion;
//   myEnv->GetStringUTFRegion = GetStringUTFRegion;
//   myEnv->GetPrimitiveArrayCritical = GetPrimitiveArrayCritical;
//   myEnv->ReleasePrimitiveArrayCritical = ReleasePrimitiveArrayCritical;
//   myEnv->GetStringCritical = GetStringCritical;
//   myEnv->ReleaseStringCritical = ReleaseStringCritical;
//   myEnv->NewWeakGlobalRef = NewWeakGlobalRef;
//   myEnv->DeleteWeakGlobalRef = DeleteWeakGlobalRef;
//   myEnv->ExceptionCheck = ExceptionCheck;
//   myEnv->NewDirectByteBuffer = NewDirectByteBuffer;
//   myEnv->GetDirectBufferAddress = GetDirectBufferAddress;
//   myEnv->GetDirectBufferCapacity = GetDirectBufferCapacity;
//   myEnv->GetObjectRefType = GetObjectRefType;
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

//export DefineClass
func DefineClass(env *C.JNIEnv, name *C.char, loader C.jobject, buf *C.jbyte, len C.jsize) C.jclass {panic("JNI:DefineClass()!")}
//export FindClass
func FindClass(env *C.JNIEnv, name *C.char) C.jclass {panic("JNI:FindClass()!")}
//export FromReflectedMethod
func FromReflectedMethod(env *C.JNIEnv, method C.jobject) C.jmethodID {panic("JNI:FromReflectedMethod()!")}
//export FromReflectedField
func FromReflectedField(env *C.JNIEnv, field C.jobject) C.jfieldID {panic("JNI:FromReflectedField()!")}
//export ToReflectedMethod
func ToReflectedMethod(env *C.JNIEnv, cls C.jclass, methodID C.jmethodID, isStatic C.jboolean) C.jobject {panic("JNI:ToReflectedMethod()!")}
//export GetSuperclass
func GetSuperclass(env *C.JNIEnv, sub C.jclass) C.jclass {panic("JNI:GetSuperclass()!")}
//export IsAssignableFrom
func IsAssignableFrom(env *C.JNIEnv, sub C.jclass, sup C.jclass) C.jboolean {panic("JNI:IsAssignableFrom()!")}
//export ToReflectedField
func ToReflectedField(env *C.JNIEnv, cls C.jclass, fieldID C.jfieldID, isStatic C.jboolean) C.jobject {panic("JNI:ToReflectedField()!")}
//export Throw
func Throw(env *C.JNIEnv, obj C.jthrowable) C.jint {panic("JNI:Throw()!")}
//export ThrowNew
func ThrowNew(env *C.JNIEnv, clazz C.jclass, msg *C.char) C.jint {panic("JNI:ThrowNew()!")}
//export ExceptionOccurred
func ExceptionOccurred(env *C.JNIEnv) C.jthrowable {panic("JNI:ExceptionOccurred()!")}
//export ExceptionDescribe
func ExceptionDescribe(env *C.JNIEnv) {panic("JNI:ExceptionDescribe()!")}
//export ExceptionClear
func ExceptionClear(env *C.JNIEnv) {panic("JNI:ExceptionClear()!")}
//export FatalError
func FatalError(env *C.JNIEnv, msg *C.char) {panic("JNI:FatalError()!")}
//export PushLocalFrame
func PushLocalFrame(env *C.JNIEnv, capacity C.jint) C.jint {panic("JNI:PushLocalFrame()!")}
//export PopLocalFrame
func PopLocalFrame(env *C.JNIEnv, result C.jobject) C.jobject {panic("JNI:PopLocalFrame()!")}
//export NewGlobalRef
func NewGlobalRef(env *C.JNIEnv, lobj C.jobject) C.jobject {panic("JNI:NewGlobalRef()!")}
//export DeleteGlobalRef
func DeleteGlobalRef(env *C.JNIEnv, gref C.jobject) {panic("JNI:DeleteGlobalRef()!")}
//export DeleteLocalRef
func DeleteLocalRef(env *C.JNIEnv, obj C.jobject) {panic("JNI:DeleteLocalRef()!")}
//export IsSameObject
func IsSameObject(env *C.JNIEnv, obj1 C.jobject, obj2 C.jobject) C.jboolean {panic("JNI:IsSameObject()!")}
//export NewLocalRef
func NewLocalRef(env *C.JNIEnv, ref C.jobject) C.jobject {panic("JNI:NewLocalRef()!")}
//export EnsureLocalCapacity
func EnsureLocalCapacity(env *C.JNIEnv, capacity C.jint) C.jint {panic("JNI:EnsureLocalCapacity()!")}
//export AllocObject
func AllocObject(env *C.JNIEnv, clazz C.jclass) C.jobject {panic("JNI:AllocObject()!")}
////export NewObject
//func NewObject(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID, ...) C.jobject {panic("JNI:NewObject()!")}
////export NewObjectV
//func NewObjectV(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID, args C.va_list) C.jobject {panic("JNI:NewObjectV()!")}
//export NewObjectA
func NewObjectA(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID, args *C.jvalue) C.jobject {panic("JNI:NewObjectA()!")}
//export GetObjectClass
func GetObjectClass(env *C.JNIEnv, obj C.jobject) C.jclass {panic("JNI:GetObjectClass()!")}
//export IsInstanceOf
func IsInstanceOf(env *C.JNIEnv, obj C.jobject, clazz C.jclass) C.jboolean {panic("JNI:IsInstanceOf()!")}
//export GetMethodID
func GetMethodID(env *C.JNIEnv, clazz C.jclass, name *C.char, sig *C.char) C.jmethodID {panic("JNI:GetMethodID()!")}
////export CallObjectMethod
//func CallObjectMethod(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, ...) C.jobject {panic("JNI:CallObjectMethod()!")}
////export CallObjectMethodV
//func CallObjectMethodV(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, args C.va_list) C.jobject {panic("JNI:CallObjectMethodV()!")}
//export CallObjectMethodA
func CallObjectMethodA(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, args *C.jvalue) C.jobject {panic("JNI:CallObjectMethodA()!")}
////export CallBooleanMethod
//func CallBooleanMethod(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, ...) C.jboolean {panic("JNI:CallBooleanMethod()!")}
////export CallBooleanMethodV
//func CallBooleanMethodV(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, args C.va_list) C.jboolean {panic("JNI:CallBooleanMethodV()!")}
//export CallBooleanMethodA
func CallBooleanMethodA(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, args *C.jvalue) C.jboolean {panic("JNI:CallBooleanMethodA()!")}
////export CallByteMethod
//func CallByteMethod(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, ...) C.jbyte {panic("JNI:CallByteMethod()!")}
////export CallByteMethodV
//func CallByteMethodV(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, args C.va_list) C.jbyte {panic("JNI:CallByteMethodV()!")}
//export CallByteMethodA
func CallByteMethodA(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, args *C.jvalue) C.jbyte {panic("JNI:CallByteMethodA()!")}
////export CallCharMethod
//func CallCharMethod(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, ...) C.jchar {panic("JNI:CallCharMethod()!")}
////export CallCharMethodV
//func CallCharMethodV(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, args C.va_list) C.jchar {panic("JNI:CallCharMethodV()!")}
//export CallCharMethodA
func CallCharMethodA(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, args *C.jvalue) C.jchar {panic("JNI:CallCharMethodA()!")}
////export CallShortMethod
//func CallShortMethod(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, ...) C.jshort {panic("JNI:CallShortMethod()!")}
////export CallShortMethodV
//func CallShortMethodV(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, args C.va_list) C.jshort {panic("JNI:CallShortMethodV()!")}
//export CallShortMethodA
func CallShortMethodA(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, args *C.jvalue) C.jshort {panic("JNI:CallShortMethodA()!")}
////export CallIntMethod
//func CallIntMethod(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, ...) C.jint {panic("JNI:CallIntMethod()!")}
////export CallIntMethodV
//func CallIntMethodV(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, args C.va_list) C.jint {panic("JNI:CallIntMethodV()!")}
//export CallIntMethodA
func CallIntMethodA(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, args *C.jvalue) C.jint {panic("JNI:CallIntMethodA()!")}
////export CallLongMethod
//func CallLongMethod(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, ...) C.jlong {panic("JNI:CallLongMethod()!")}
////export CallLongMethodV
//func CallLongMethodV(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, args C.va_list) C.jlong {panic("JNI:CallLongMethodV()!")}
//export CallLongMethodA
func CallLongMethodA(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, args *C.jvalue) C.jlong {panic("JNI:CallLongMethodA()!")}
////export CallFloatMethod
//func CallFloatMethod(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, ...) C.jfloat {panic("JNI:CallFloatMethod()!")}
////export CallFloatMethodV
//func CallFloatMethodV(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, args C.va_list) C.jfloat {panic("JNI:CallFloatMethodV()!")}
//export CallFloatMethodA
func CallFloatMethodA(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, args *C.jvalue) C.jfloat {panic("JNI:CallFloatMethodA()!")}
////export CallDoubleMethod
//func CallDoubleMethod(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, ...) C.jdouble {panic("JNI:CallDoubleMethod()!")}
////export CallDoubleMethodV
//func CallDoubleMethodV(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, args C.va_list) C.jdouble {panic("JNI:CallDoubleMethodV()!")}
//export CallDoubleMethodA
func CallDoubleMethodA(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, args *C.jvalue) C.jdouble {panic("JNI:CallDoubleMethodA()!")}
////export CallVoidMethod
//func CallVoidMethod(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, ...) {panic("JNI:CallVoidMethod()!")}
////export CallVoidMethodV
//func CallVoidMethodV(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, args C.va_list) {panic("JNI:CallVoidMethodV()!")}
//export CallVoidMethodA
func CallVoidMethodA(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, args *C.jvalue) {panic("JNI:CallVoidMethodA()!")}
////export CallNonvirtualObjectMethod
//func CallNonvirtualObjectMethod(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID, ...) C.jobject {panic("JNI:CallNonvirtualObjectMethod()!")}
////export CallNonvirtualObjectMethodV
//func CallNonvirtualObjectMethodV(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID, args C.va_list) C.jobject {panic("JNI:CallNonvirtualObjectMethodV()!")}
//export CallNonvirtualObjectMethodA
func CallNonvirtualObjectMethodA(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID, args *C.jvalue) C.jobject {panic("JNI:CallNonvirtualObjectMethodA()!")}
////export CallNonvirtualBooleanMethod
//func CallNonvirtualBooleanMethod(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID, ...) C.jboolean {panic("JNI:CallNonvirtualBooleanMethod()!")}
////export CallNonvirtualBooleanMethodV
//func CallNonvirtualBooleanMethodV(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID, args C.va_list) C.jboolean {panic("JNI:CallNonvirtualBooleanMethodV()!")}
//export CallNonvirtualBooleanMethodA
func CallNonvirtualBooleanMethodA(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID, args *C.jvalue) C.jboolean {panic("JNI:CallNonvirtualBooleanMethodA()!")}
////export CallNonvirtualByteMethod
//func CallNonvirtualByteMethod(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID, ...) C.jbyte {panic("JNI:CallNonvirtualByteMethod()!")}
////export CallNonvirtualByteMethodV
//func CallNonvirtualByteMethodV(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID, args C.va_list) C.jbyte {panic("JNI:CallNonvirtualByteMethodV()!")}
//export CallNonvirtualByteMethodA
func CallNonvirtualByteMethodA(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID, args *C.jvalue) C.jbyte {panic("JNI:CallNonvirtualByteMethodA()!")}
////export CallNonvirtualCharMethod
//func CallNonvirtualCharMethod(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID, ...) C.jchar {panic("JNI:CallNonvirtualCharMethod()!")}
////export CallNonvirtualCharMethodV
//func CallNonvirtualCharMethodV(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID, args C.va_list) C.jchar {panic("JNI:CallNonvirtualCharMethodV()!")}
//export CallNonvirtualCharMethodA
func CallNonvirtualCharMethodA(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID, args *C.jvalue) C.jchar {panic("JNI:CallNonvirtualCharMethodA()!")}
////export CallNonvirtualShortMethod
//func CallNonvirtualShortMethod(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID, ...) C.jshort {panic("JNI:CallNonvirtualShortMethod()!")}
////export CallNonvirtualShortMethodV
//func CallNonvirtualShortMethodV(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID, args C.va_list) C.jshort {panic("JNI:CallNonvirtualShortMethodV()!")}
//export CallNonvirtualShortMethodA
func CallNonvirtualShortMethodA(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID, args *C.jvalue) C.jshort {panic("JNI:CallNonvirtualShortMethodA()!")}
////export CallNonvirtualIntMethod
//func CallNonvirtualIntMethod(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID, ...) C.jint {panic("JNI:CallNonvirtualIntMethod()!")}
////export CallNonvirtualIntMethodV
//func CallNonvirtualIntMethodV(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID, args C.va_list) C.jint {panic("JNI:CallNonvirtualIntMethodV()!")}
//export CallNonvirtualIntMethodA
func CallNonvirtualIntMethodA(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID, args *C.jvalue) C.jint {panic("JNI:CallNonvirtualIntMethodA()!")}
////export CallNonvirtualLongMethod
//func CallNonvirtualLongMethod(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID, ...) C.jlong {panic("JNI:CallNonvirtualLongMethod()!")}
////export CallNonvirtualLongMethodV
//func CallNonvirtualLongMethodV(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID, args C.va_list) C.jlong {panic("JNI:CallNonvirtualLongMethodV()!")}
//export CallNonvirtualLongMethodA
func CallNonvirtualLongMethodA(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID, args *C.jvalue) C.jlong {panic("JNI:CallNonvirtualLongMethodA()!")}
////export CallNonvirtualFloatMethod
//func CallNonvirtualFloatMethod(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID, ...) C.jfloat {panic("JNI:CallNonvirtualFloatMethod()!")}
////export CallNonvirtualFloatMethodV
//func CallNonvirtualFloatMethodV(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID, args C.va_list) C.jfloat {panic("JNI:CallNonvirtualFloatMethodV()!")}
//export CallNonvirtualFloatMethodA
func CallNonvirtualFloatMethodA(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID, args *C.jvalue) C.jfloat {panic("JNI:CallNonvirtualFloatMethodA()!")}
////export CallNonvirtualDoubleMethod
//func CallNonvirtualDoubleMethod(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID, ...) C.jdouble {panic("JNI:CallNonvirtualDoubleMethod()!")}
////export CallNonvirtualDoubleMethodV
//func CallNonvirtualDoubleMethodV(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID, args C.va_list) C.jdouble {panic("JNI:CallNonvirtualDoubleMethodV()!")}
//export CallNonvirtualDoubleMethodA
func CallNonvirtualDoubleMethodA(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID, args *C.jvalue) C.jdouble {panic("JNI:CallNonvirtualDoubleMethodA()!")}
////export CallNonvirtualVoidMethod
//func CallNonvirtualVoidMethod(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID, ...) {panic("JNI:CallNonvirtualVoidMethod()!")}
////export CallNonvirtualVoidMethodV
//func CallNonvirtualVoidMethodV(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID, args C.va_list) {panic("JNI:CallNonvirtualVoidMethodV()!")}
//export CallNonvirtualVoidMethodA
func CallNonvirtualVoidMethodA(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID, args *C.jvalue) {panic("JNI:CallNonvirtualVoidMethodA()!")}
//export GetFieldID
func GetFieldID(env *C.JNIEnv, clazz C.jclass, name *C.char, sig *C.char) C.jfieldID {panic("JNI:GetFieldID()!")}
//export GetObjectField
func GetObjectField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID) C.jobject {panic("JNI:GetObjectField()!")}
//export GetBooleanField
func GetBooleanField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID) C.jboolean {panic("JNI:GetBooleanField()!")}
//export GetByteField
func GetByteField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID) C.jbyte {panic("JNI:GetByteField()!")}
//export GetCharField
func GetCharField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID) C.jchar {panic("JNI:GetCharField()!")}
//export GetShortField
func GetShortField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID) C.jshort {panic("JNI:GetShortField()!")}
//export GetIntField
func GetIntField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID) C.jint {panic("JNI:GetIntField()!")}
//export GetLongField
func GetLongField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID) C.jlong {panic("JNI:GetLongField()!")}
//export GetFloatField
func GetFloatField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID) C.jfloat {panic("JNI:GetFloatField()!")}
//export GetDoubleField
func GetDoubleField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID) C.jdouble {panic("JNI:GetDoubleField()!")}
//export SetObjectField
func SetObjectField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID, val C.jobject) {panic("JNI:SetObjectField()!")}
//export SetBooleanField
func SetBooleanField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID, val C.jboolean) {panic("JNI:SetBooleanField()!")}
//export SetByteField
func SetByteField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID, val C.jbyte) {panic("JNI:SetByteField()!")}
//export SetCharField
func SetCharField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID, val C.jchar) {panic("JNI:SetCharField()!")}
//export SetShortField
func SetShortField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID, val C.jshort) {panic("JNI:SetShortField()!")}
//export SetIntField
func SetIntField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID, val C.jint) {panic("JNI:SetIntField()!")}
//export SetLongField
func SetLongField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID, val C.jlong) {panic("JNI:SetLongField()!")}
//export SetFloatField
func SetFloatField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID, val C.jfloat) {panic("JNI:SetFloatField()!")}
//export SetDoubleField
func SetDoubleField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID, val C.jdouble) {panic("JNI:SetDoubleField()!")}
//export GetStaticMethodID
func GetStaticMethodID(env *C.JNIEnv, clazz C.jclass, name *C.char, sig *C.char) C.jmethodID {panic("JNI:GetStaticMethodID()!")}
////export CallStaticObjectMethod
//func CallStaticObjectMethod(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID, ...) C.jobject {panic("JNI:CallStaticObjectMethod()!")}
////export CallStaticObjectMethodV
//func CallStaticObjectMethodV(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID, args C.va_list) C.jobject {panic("JNI:CallStaticObjectMethodV()!")}
//export CallStaticObjectMethodA
func CallStaticObjectMethodA(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID, args *C.jvalue) C.jobject {panic("JNI:CallStaticObjectMethodA()!")}
////export CallStaticBooleanMethod
//func CallStaticBooleanMethod(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID, ...) C.jboolean {panic("JNI:CallStaticBooleanMethod()!")}
////export CallStaticBooleanMethodV
//func CallStaticBooleanMethodV(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID, args C.va_list) C.jboolean {panic("JNI:CallStaticBooleanMethodV()!")}
//export CallStaticBooleanMethodA
func CallStaticBooleanMethodA(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID, args *C.jvalue) C.jboolean {panic("JNI:CallStaticBooleanMethodA()!")}
////export CallStaticByteMethod
//func CallStaticByteMethod(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID, ...) C.jbyte {panic("JNI:CallStaticByteMethod()!")}
////export CallStaticByteMethodV
//func CallStaticByteMethodV(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID, args C.va_list) C.jbyte {panic("JNI:CallStaticByteMethodV()!")}
//export CallStaticByteMethodA
func CallStaticByteMethodA(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID, args *C.jvalue) C.jbyte {panic("JNI:CallStaticByteMethodA()!")}
////export CallStaticCharMethod
//func CallStaticCharMethod(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID, ...) C.jchar {panic("JNI:CallStaticCharMethod()!")}
////export CallStaticCharMethodV
//func CallStaticCharMethodV(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID, args C.va_list) C.jchar {panic("JNI:CallStaticCharMethodV()!")}
//export CallStaticCharMethodA
func CallStaticCharMethodA(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID, args *C.jvalue) C.jchar {panic("JNI:CallStaticCharMethodA()!")}
////export CallStaticShortMethod
//func CallStaticShortMethod(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID, ...) C.jshort {panic("JNI:CallStaticShortMethod()!")}
////export CallStaticShortMethodV
//func CallStaticShortMethodV(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID, args C.va_list) C.jshort {panic("JNI:CallStaticShortMethodV()!")}
//export CallStaticShortMethodA
func CallStaticShortMethodA(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID, args *C.jvalue) C.jshort {panic("JNI:CallStaticShortMethodA()!")}
////export CallStaticIntMethod
//func CallStaticIntMethod(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID, ...) C.jint {panic("JNI:CallStaticIntMethod()!")}
////export CallStaticIntMethodV
//func CallStaticIntMethodV(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID, args C.va_list) C.jint {panic("JNI:CallStaticIntMethodV()!")}
//export CallStaticIntMethodA
func CallStaticIntMethodA(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID, args *C.jvalue) C.jint {panic("JNI:CallStaticIntMethodA()!")}
////export CallStaticLongMethod
//func CallStaticLongMethod(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID, ...) C.jlong {panic("JNI:CallStaticLongMethod()!")}
////export CallStaticLongMethodV
//func CallStaticLongMethodV(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID, args C.va_list) C.jlong {panic("JNI:CallStaticLongMethodV()!")}
//export CallStaticLongMethodA
func CallStaticLongMethodA(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID, args *C.jvalue) C.jlong {panic("JNI:CallStaticLongMethodA()!")}
////export CallStaticFloatMethod
//func CallStaticFloatMethod(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID, ...) C.jfloat {panic("JNI:CallStaticFloatMethod()!")}
////export CallStaticFloatMethodV
//func CallStaticFloatMethodV(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID, args C.va_list) C.jfloat {panic("JNI:CallStaticFloatMethodV()!")}
//export CallStaticFloatMethodA
func CallStaticFloatMethodA(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID, args *C.jvalue) C.jfloat {panic("JNI:CallStaticFloatMethodA()!")}
////export CallStaticDoubleMethod
//func CallStaticDoubleMethod(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID, ...) C.jdouble {panic("JNI:CallStaticDoubleMethod()!")}
////export CallStaticDoubleMethodV
//func CallStaticDoubleMethodV(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID, args C.va_list) C.jdouble {panic("JNI:CallStaticDoubleMethodV()!")}
//export CallStaticDoubleMethodA
func CallStaticDoubleMethodA(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID, args *C.jvalue) C.jdouble {panic("JNI:CallStaticDoubleMethodA()!")}
////export CallStaticVoidMethod
//func CallStaticVoidMethod(env *C.JNIEnv, cls C.jclass, methodID C.jmethodID, ...) {panic("JNI:CallStaticVoidMethod()!")}
////export CallStaticVoidMethodV
//func CallStaticVoidMethodV(env *C.JNIEnv, cls C.jclass, methodID C.jmethodID, args C.va_list) {panic("JNI:CallStaticVoidMethodV()!")}
//export CallStaticVoidMethodA
func CallStaticVoidMethodA(env *C.JNIEnv, cls C.jclass, methodID C.jmethodID, args *C.jvalue) {panic("JNI:CallStaticVoidMethodA()!")}
//export GetStaticFieldID
func GetStaticFieldID(env *C.JNIEnv, clazz C.jclass, name *C.char, sig *C.char) C.jfieldID {panic("JNI:GetStaticFieldID()!")}
//export GetStaticObjectField
func GetStaticObjectField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID) C.jobject {panic("JNI:GetStaticObjectField()!")}
//export GetStaticBooleanField
func GetStaticBooleanField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID) C.jboolean {panic("JNI:GetStaticBooleanField()!")}
//export GetStaticByteField
func GetStaticByteField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID) C.jbyte {panic("JNI:GetStaticByteField()!")}
//export GetStaticCharField
func GetStaticCharField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID) C.jchar {panic("JNI:GetStaticCharField()!")}
//export GetStaticShortField
func GetStaticShortField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID) C.jshort {panic("JNI:GetStaticShortField()!")}
//export GetStaticIntField
func GetStaticIntField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID) C.jint {panic("JNI:GetStaticIntField()!")}
//export GetStaticLongField
func GetStaticLongField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID) C.jlong {panic("JNI:GetStaticLongField()!")}
//export GetStaticFloatField
func GetStaticFloatField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID) C.jfloat {panic("JNI:GetStaticFloatField()!")}
//export GetStaticDoubleField
func GetStaticDoubleField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID) C.jdouble {panic("JNI:GetStaticDoubleField()!")}
//export SetStaticObjectField
func SetStaticObjectField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID, value C.jobject) {panic("JNI:SetStaticObjectField()!")}
//export SetStaticBooleanField
func SetStaticBooleanField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID, value C.jboolean) {panic("JNI:SetStaticBooleanField()!")}
//export SetStaticByteField
func SetStaticByteField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID, value C.jbyte) {panic("JNI:SetStaticByteField()!")}
//export SetStaticCharField
func SetStaticCharField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID, value C.jchar) {panic("JNI:SetStaticCharField()!")}
//export SetStaticShortField
func SetStaticShortField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID, value C.jshort) {panic("JNI:SetStaticShortField()!")}
//export SetStaticIntField
func SetStaticIntField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID, value C.jint) {panic("JNI:SetStaticIntField()!")}
//export SetStaticLongField
func SetStaticLongField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID, value C.jlong) {panic("JNI:SetStaticLongField()!")}
//export SetStaticFloatField
func SetStaticFloatField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID, value C.jfloat) {panic("JNI:SetStaticFloatField()!")}
//export SetStaticDoubleField
func SetStaticDoubleField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID, value C.jdouble) {panic("JNI:SetStaticDoubleField()!")}
//export NewString
func NewString(env *C.JNIEnv, unicode *C.jchar, len C.jsize) C.jstring {panic("JNI:NewString()!")}
//export GetStringLength
func GetStringLength(env *C.JNIEnv, str C.jstring) C.jsize {panic("JNI:GetStringLength()!")}
//export GetStringChars
func GetStringChars(env *C.JNIEnv, str C.jstring, isCopy *C.jboolean) *C.jchar {panic("JNI:GetStringChars()!")}
//export ReleaseStringChars
func ReleaseStringChars(env *C.JNIEnv, str C.jstring, chars *C.jchar) {panic("JNI:ReleaseStringChars()!")}
//export NewStringUTF
func NewStringUTF(env *C.JNIEnv, utf *C.char) C.jstring {panic("JNI:NewStringUTF()!")}
//export GetStringUTFLength
func GetStringUTFLength(env *C.JNIEnv, str C.jstring) C.jsize {panic("JNI:GetStringUTFLength()!")}
//export GetStringUTFChars
func GetStringUTFChars(env *C.JNIEnv, str C.jstring, isCopy *C.jboolean) *C.char {panic("JNI:GetStringUTFChars()!")}
//export ReleaseStringUTFChars
func ReleaseStringUTFChars(env *C.JNIEnv, str C.jstring, chars *C.char) {panic("JNI:ReleaseStringUTFChars()!")}
//export GetArrayLength
func GetArrayLength(env *C.JNIEnv, array C.jarray) C.jsize {panic("JNI:GetArrayLength()!")}
//export NewObjectArray
func NewObjectArray(env *C.JNIEnv, len C.jsize, clazz C.jclass, init C.jobject) C.jobjectArray {panic("JNI:NewObjectArray()!")}
//export GetObjectArrayElement
func GetObjectArrayElement(env *C.JNIEnv, array C.jobjectArray, index C.jsize) C.jobject {panic("JNI:GetObjectArrayElement()!")}
//export SetObjectArrayElement
func SetObjectArrayElement(env *C.JNIEnv, array C.jobjectArray, index C.jsize, val C.jobject) {panic("JNI:SetObjectArrayElement()!")}
//export NewBooleanArray
func NewBooleanArray(env *C.JNIEnv, len C.jsize) C.jbooleanArray {panic("JNI:NewBooleanArray()!")}
//export NewByteArray
func NewByteArray(env *C.JNIEnv, len C.jsize) C.jbyteArray {panic("JNI:NewByteArray()!")}
//export NewCharArray
func NewCharArray(env *C.JNIEnv, len C.jsize) C.jcharArray {panic("JNI:NewCharArray()!")}
//export NewShortArray
func NewShortArray(env *C.JNIEnv, len C.jsize) C.jshortArray {panic("JNI:NewShortArray()!")}
//export NewIntArray
func NewIntArray(env *C.JNIEnv, len C.jsize) C.jintArray {panic("JNI:NewIntArray()!")}
//export NewLongArray
func NewLongArray(env *C.JNIEnv, len C.jsize) C.jlongArray {panic("JNI:NewLongArray()!")}
//export NewFloatArray
func NewFloatArray(env *C.JNIEnv, len C.jsize) C.jfloatArray {panic("JNI:NewFloatArray()!")}
//export NewDoubleArray
func NewDoubleArray(env *C.JNIEnv, len C.jsize) C.jdoubleArray {panic("JNI:NewDoubleArray()!")}
//export GetBooleanArrayElements
func GetBooleanArrayElements(env *C.JNIEnv, array C.jbooleanArray, isCopy *C.jboolean) *C.jboolean {panic("JNI:GetBooleanArrayElements()!")}
//export GetByteArrayElements
func GetByteArrayElements(env *C.JNIEnv, array C.jbyteArray, isCopy *C.jboolean) *C.jbyte {panic("JNI:GetByteArrayElements()!")}
//export GetCharArrayElements
func GetCharArrayElements(env *C.JNIEnv, array C.jcharArray, isCopy *C.jboolean) *C.jchar {panic("JNI:GetCharArrayElements()!")}
//export GetShortArrayElements
func GetShortArrayElements(env *C.JNIEnv, array C.jshortArray, isCopy *C.jboolean) *C.jshort {panic("JNI:GetShortArrayElements()!")}
//export GetIntArrayElements
func GetIntArrayElements(env *C.JNIEnv, array C.jintArray, isCopy *C.jboolean) *C.jint {panic("JNI:GetIntArrayElements()!")}
//export GetLongArrayElements
func GetLongArrayElements(env *C.JNIEnv, array C.jlongArray, isCopy *C.jboolean) *C.jlong {panic("JNI:GetLongArrayElements()!")}
//export GetFloatArrayElements
func GetFloatArrayElements(env *C.JNIEnv, array C.jfloatArray, isCopy *C.jboolean) *C.jfloat {panic("JNI:GetFloatArrayElements()!")}
//export GetDoubleArrayElements
func GetDoubleArrayElements(env *C.JNIEnv, array C.jdoubleArray, isCopy *C.jboolean) *C.jdouble {panic("JNI:GetDoubleArrayElements()!")}
//export ReleaseBooleanArrayElements
func ReleaseBooleanArrayElements(env *C.JNIEnv, array C.jbooleanArray, elems *C.jboolean, mode C.jint) {panic("JNI:ReleaseBooleanArrayElements()!")}
//export ReleaseByteArrayElements
func ReleaseByteArrayElements(env *C.JNIEnv, array C.jbyteArray, elems *C.jbyte, mode C.jint) {panic("JNI:ReleaseByteArrayElements()!")}
//export ReleaseCharArrayElements
func ReleaseCharArrayElements(env *C.JNIEnv, array C.jcharArray, elems *C.jchar, mode C.jint) {panic("JNI:ReleaseCharArrayElements()!")}
//export ReleaseShortArrayElements
func ReleaseShortArrayElements(env *C.JNIEnv, array C.jshortArray, elems *C.jshort, mode C.jint) {panic("JNI:ReleaseShortArrayElements()!")}
//export ReleaseIntArrayElements
func ReleaseIntArrayElements(env *C.JNIEnv, array C.jintArray, elems *C.jint, mode C.jint) {panic("JNI:ReleaseIntArrayElements()!")}
//export ReleaseLongArrayElements
func ReleaseLongArrayElements(env *C.JNIEnv, array C.jlongArray, elems *C.jlong, mode C.jint) {panic("JNI:ReleaseLongArrayElements()!")}
//export ReleaseFloatArrayElements
func ReleaseFloatArrayElements(env *C.JNIEnv, array C.jfloatArray, elems *C.jfloat, mode C.jint) {panic("JNI:ReleaseFloatArrayElements()!")}
//export ReleaseDoubleArrayElements
func ReleaseDoubleArrayElements(env *C.JNIEnv, array C.jdoubleArray, elems *C.jdouble, mode C.jint) {panic("JNI:ReleaseDoubleArrayElements()!")}
//export GetBooleanArrayRegion
func GetBooleanArrayRegion(env *C.JNIEnv, array C.jbooleanArray, start C.jsize, l C.jsize, buf *C.jboolean) {panic("JNI:GetBooleanArrayRegion()!")}
//export GetByteArrayRegion
func GetByteArrayRegion(env *C.JNIEnv, array C.jbyteArray, start C.jsize, len C.jsize, buf *C.jbyte) {panic("JNI:GetByteArrayRegion()!")}
//export GetCharArrayRegion
func GetCharArrayRegion(env *C.JNIEnv, array C.jcharArray, start C.jsize, len C.jsize, buf *C.jchar) {panic("JNI:GetCharArrayRegion()!")}
//export GetShortArrayRegion
func GetShortArrayRegion(env *C.JNIEnv, array C.jshortArray, start C.jsize, len C.jsize, buf *C.jshort) {panic("JNI:GetShortArrayRegion()!")}
//export GetIntArrayRegion
func GetIntArrayRegion(env *C.JNIEnv, array C.jintArray, start C.jsize, len C.jsize, buf *C.jint) {panic("JNI:GetIntArrayRegion()!")}
//export GetLongArrayRegion
func GetLongArrayRegion(env *C.JNIEnv, array C.jlongArray, start C.jsize, len C.jsize, buf *C.jlong) {panic("JNI:GetLongArrayRegion()!")}
//export GetFloatArrayRegion
func GetFloatArrayRegion(env *C.JNIEnv, array C.jfloatArray, start C.jsize, len C.jsize, buf *C.jfloat) {panic("JNI:GetFloatArrayRegion()!")}
//export GetDoubleArrayRegion
func GetDoubleArrayRegion(env *C.JNIEnv, array C.jdoubleArray, start C.jsize, len C.jsize, buf *C.jdouble) {panic("JNI:GetDoubleArrayRegion()!")}
//export SetBooleanArrayRegion
func SetBooleanArrayRegion(env *C.JNIEnv, array C.jbooleanArray, start C.jsize, l C.jsize, buf *C.jboolean) {panic("JNI:SetBooleanArrayRegion()!")}
//export SetByteArrayRegion
func SetByteArrayRegion(env *C.JNIEnv, array C.jbyteArray, start C.jsize, len C.jsize, buf *C.jbyte) {panic("JNI:SetByteArrayRegion()!")}
//export SetCharArrayRegion
func SetCharArrayRegion(env *C.JNIEnv, array C.jcharArray, start C.jsize, len C.jsize, buf *C.jchar) {panic("JNI:SetCharArrayRegion()!")}
//export SetShortArrayRegion
func SetShortArrayRegion(env *C.JNIEnv, array C.jshortArray, start C.jsize, len C.jsize, buf *C.jshort) {panic("JNI:SetShortArrayRegion()!")}
//export SetIntArrayRegion
func SetIntArrayRegion(env *C.JNIEnv, array C.jintArray, start C.jsize, len C.jsize, buf *C.jint) {panic("JNI:SetIntArrayRegion()!")}
//export SetLongArrayRegion
func SetLongArrayRegion(env *C.JNIEnv, array C.jlongArray, start C.jsize, len C.jsize, buf *C.jlong) {panic("JNI:SetLongArrayRegion()!")}
//export SetFloatArrayRegion
func SetFloatArrayRegion(env *C.JNIEnv, array C.jfloatArray, start C.jsize, len C.jsize, buf *C.jfloat) {panic("JNI:SetFloatArrayRegion()!")}
//export SetDoubleArrayRegion
func SetDoubleArrayRegion(env *C.JNIEnv, array C.jdoubleArray, start C.jsize, len C.jsize, buf *C.jdouble) {panic("JNI:SetDoubleArrayRegion()!")}
//export RegisterNatives
func RegisterNatives(env *C.JNIEnv, clazz C.jclass, methods *C.JNINativeMethod, nMethods C.jint) C.jint {panic("JNI:RegisterNatives()!")}
//export UnregisterNatives
func UnregisterNatives(env *C.JNIEnv, clazz C.jclass) C.jint {panic("JNI:UnregisterNatives()!")}
//export MonitorEnter
func MonitorEnter(env *C.JNIEnv, obj C.jobject) C.jint {panic("JNI:MonitorEnter()!")}
//export MonitorExit
func MonitorExit(env *C.JNIEnv, obj C.jobject) C.jint {panic("JNI:MonitorExit()!")}
//export GetJavaVM
func GetJavaVM(env *C.JNIEnv, vm **C.JavaVM) C.jint {panic("JNI:GetJavaVM()!")}
//export GetStringRegion
func GetStringRegion(env *C.JNIEnv, str C.jstring, start C.jsize, len C.jsize, buf *C.jchar) {panic("JNI:GetStringRegion()!")}
//export GetStringUTFRegion
func GetStringUTFRegion(env *C.JNIEnv, str C.jstring, start C.jsize, len C.jsize, buf *C.char) {panic("JNI:GetStringUTFRegion()!")}
//export GetPrimitiveArrayCritical
func GetPrimitiveArrayCritical(env *C.JNIEnv, array C.jarray, isCopy *C.jboolean) unsafe.Pointer {panic("JNI:GetPrimitiveArrayCritical()!")}
//export ReleasePrimitiveArrayCritical
func ReleasePrimitiveArrayCritical(env *C.JNIEnv, array C.jarray, carray unsafe.Pointer, mode C.jint) {panic("JNI:ReleasePrimitiveArrayCritical()!")}
//export GetStringCritical
func GetStringCritical(env *C.JNIEnv, string C.jstring, isCopy *C.jboolean) *C.jchar {panic("JNI:GetStringCritical()!")}
//export ReleaseStringCritical
func ReleaseStringCritical(env *C.JNIEnv, string C.jstring, cstring *C.jchar) {panic("JNI:ReleaseStringCritical()!")}
//export NewWeakGlobalRef
func NewWeakGlobalRef(env *C.JNIEnv, obj C.jobject) C.jweak {panic("JNI:NewWeakGlobalRef()!")}
//export DeleteWeakGlobalRef
func DeleteWeakGlobalRef(env *C.JNIEnv, ref C.jweak) {panic("JNI:DeleteWeakGlobalRef()!")}
//export ExceptionCheck
func ExceptionCheck(env *C.JNIEnv) C.jboolean {panic("JNI:ExceptionCheck()!")}
//export NewDirectByteBuffer
func NewDirectByteBuffer(env *C.JNIEnv, address unsafe.Pointer, capacity C.jlong) C.jobject {panic("JNI:NewDirectByteBuffer()!")}
//export GetDirectBufferAddress
func GetDirectBufferAddress(env *C.JNIEnv, buf C.jobject) unsafe.Pointer {panic("JNI:GetDirectBufferAddress()!")}
//export GetDirectBufferCapacity
func GetDirectBufferCapacity(env *C.JNIEnv, buf C.jobject) C.jlong {panic("JNI:GetDirectBufferCapacity()!")}
//export GetObjectRefType
func GetObjectRefType(env *C.JNIEnv, obj C.jobject) C.jobjectRefType {panic("JNI:GetObjectRefType()!")}

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
