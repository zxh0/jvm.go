// #include <stdio.h>
#include <stdlib.h>
#include "wrapper.h"

void GoPanic(char *err);
int GoGetParamCount(jclass clazz, jmethodID methodID);

jint          GoGetVersion             (JNIEnv *env);
jclass        GoFindClass              (JNIEnv *env, const char *name);
jobject       GoNewGlobalRef           (JNIEnv *env, jobject lobj);
void          GoDeleteLocalRef         (JNIEnv *env, jobject obj);
jstring       GoNewString              (JNIEnv *env, const jchar *unicode, jsize len);
jstring       GoNewStringUTF           (JNIEnv *env, const char *utf);
jsize         GoGetStringLength        (JNIEnv *env, jstring str);
const jchar * GoGetStringCritical      (JNIEnv *env, jstring string, jboolean *isCopy);
void          GoReleaseStringCritical  (JNIEnv *env, jstring string, const jchar *cstring);
jmethodID     GoGetMethodID            (JNIEnv *env, jclass clazz, const char *name, const char *sig);
jmethodID     GoGetStaticMethodID      (JNIEnv *env, jclass clazz, const char *name, const char *sig);
jobject       GoCallStaticObjectMethodA(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);
jboolean      GoExceptionCheck         (JNIEnv *env);


void copyVaListToArray(va_list args, int argCount, jvalue argArr[]) {
    for (int i = 0; i < argCount; i++) {
        argArr[i] = va_arg(args, jvalue);
    }
}

jint GetVersion(JNIEnv *env) {
    GoGetVersion(env);
    GoPanic("JNICALL GetVersion()!");
    return 0;
}

jclass DefineClass(JNIEnv *env, const char *name, jobject loader, const jbyte *buf, jsize len) {
    GoPanic("JNICALL DefineClass()!");
    return 0;
}
jclass FindClass(JNIEnv *env, const char *name) {
    //GoPanic("JNICALL FindClass()!");
    return GoFindClass(env, name);
}

jmethodID FromReflectedMethod(JNIEnv *env, jobject method) {
    GoPanic("JNICALL FromReflectedMethod()!");
    return 0;
}
jfieldID FromReflectedField(JNIEnv *env, jobject field) {
    GoPanic("JNICALL FromReflectedField()!");
    return 0;
}
jobject ToReflectedMethod(JNIEnv *env, jclass cls, jmethodID methodID, jboolean isStatic) {
    GoPanic("JNICALL ToReflectedMethod()!");
    return 0;
}
jclass GetSuperclass(JNIEnv *env, jclass sub) {
    GoPanic("JNICALL GetSuperclass()!");
    return 0;
}
jboolean IsAssignableFrom(JNIEnv *env, jclass sub, jclass sup) {
    GoPanic("JNICALL IsAssignableFrom()!");
    return 0;
}
jobject ToReflectedField(JNIEnv *env, jclass cls, jfieldID fieldID, jboolean isStatic) {
    GoPanic("JNICALL ToReflectedField()!");
    return 0;
}
jint Throw(JNIEnv *env, jthrowable obj) {
    GoPanic("JNICALL Throw()!");
    return 0;
}
jint ThrowNew(JNIEnv *env, jclass clazz, const char *msg) {
    GoPanic("JNICALL ThrowNew()!");
    return 0;
}
jthrowable ExceptionOccurred(JNIEnv *env) {
    GoPanic("JNICALL ExceptionOccurred()!");
    return 0;
}
void ExceptionDescribe(JNIEnv *env) {
    GoPanic("JNICALL ExceptionDescribe()!");
}
void ExceptionClear(JNIEnv *env) {
    GoPanic("JNICALL ExceptionClear()!");
}
void FatalError(JNIEnv *env, const char *msg) {
    GoPanic("JNICALL FatalError()!");
}
jint PushLocalFrame(JNIEnv *env, jint capacity) {
    GoPanic("JNICALL PushLocalFrame()!");
    return 0;
}
jobject PopLocalFrame(JNIEnv *env, jobject result) {
    GoPanic("JNICALL PopLocalFrame()!");
    return 0;
}

jobject NewGlobalRef(JNIEnv *env, jobject lobj) {
    //GoPanic("JNICALL NewGlobalRef()!");
    return GoNewGlobalRef(env, lobj);
}
void DeleteGlobalRef(JNIEnv *env, jobject gref) {
    GoPanic("JNICALL DeleteGlobalRef()!");
}
void DeleteLocalRef(JNIEnv *env, jobject obj) {
    //GoPanic("JNICALL DeleteLocalRef()!");
    GoDeleteLocalRef(env, obj);
}
jboolean IsSameObject(JNIEnv *env, jobject obj1, jobject obj2) {
    GoPanic("JNICALL IsSameObject()!");
    return 0;
}
jobject NewLocalRef(JNIEnv *env, jobject ref) {
    GoPanic("JNICALL NewLocalRef()!");
    return 0;
}
jint EnsureLocalCapacity(JNIEnv *env, jint capacity) {
    //GoPanic("JNICALL EnsureLocalCapacity()!");
    return 0;
}

jobject AllocObject(JNIEnv *env, jclass clazz) {
    GoPanic("JNICALL AllocObject()!");
    return 0;
}
jobject NewObject(JNIEnv *env, jclass clazz, jmethodID methodID, ...) {
    GoPanic("JNICALL NewObject()!");
    return 0;
}
jobject NewObjectV(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args) {
    GoPanic("JNICALL NewObjectV()!");
    return 0;
}
jobject NewObjectA(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args) {
    GoPanic("JNICALL NewObjectA()!");
    return 0;
}
jclass GetObjectClass(JNIEnv *env, jobject obj) {
    GoPanic("JNICALL GetObjectClass()!");
    return 0;
}
jboolean IsInstanceOf(JNIEnv *env, jobject obj, jclass clazz) {
    GoPanic("JNICALL IsInstanceOf()!");
    return 0;
}

jmethodID GetMethodID(JNIEnv *env, jclass clazz, const char *name, const char *sig) {
    //GoPanic("JNICALL GetMethodID()!");
    return GoGetMethodID(env, clazz, name, sig);
}
jobject CallObjectMethod(JNIEnv *env, jobject obj, jmethodID methodID, ...) {
    GoPanic("JNICALL CallObjectMethod()!");
    return 0;
}
jobject CallObjectMethodV(JNIEnv *env, jobject obj, jmethodID methodID, va_list args) {
    GoPanic("JNICALL CallObjectMethodV()!");
    return 0;
}
jobject CallObjectMethodA(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue * args) {
    GoPanic("JNICALL CallObjectMethodA()!");
    return 0;
}
jboolean CallBooleanMethod(JNIEnv *env, jobject obj, jmethodID methodID, ...) {
    GoPanic("JNICALL CallBooleanMethod()!");
    return 0;
}
jboolean CallBooleanMethodV(JNIEnv *env, jobject obj, jmethodID methodID, va_list args) {
    GoPanic("JNICALL CallBooleanMethodV()!");
    return 0;
}
jboolean CallBooleanMethodA(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue * args) {
    GoPanic("JNICALL CallBooleanMethodA()!");
    return 0;
}
jbyte CallByteMethod(JNIEnv *env, jobject obj, jmethodID methodID, ...) {
    GoPanic("JNICALL CallByteMethod()!");
    return 0;
}
jbyte CallByteMethodV(JNIEnv *env, jobject obj, jmethodID methodID, va_list args) {
    GoPanic("JNICALL CallByteMethodV()!");
    return 0;
}
jbyte CallByteMethodA(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args) {
    GoPanic("JNICALL CallByteMethodA()!");
    return 0;
}
jchar CallCharMethod(JNIEnv *env, jobject obj, jmethodID methodID, ...) {
    GoPanic("JNICALL CallCharMethod()!");
    return 0;
}
jchar CallCharMethodV(JNIEnv *env, jobject obj, jmethodID methodID, va_list args) {
    GoPanic("JNICALL CallCharMethodV()!");
    return 0;
}
jchar CallCharMethodA(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args) {
    GoPanic("JNICALL CallCharMethodA()!");
    return 0;
}
jshort CallShortMethod(JNIEnv *env, jobject obj, jmethodID methodID, ...) {
    GoPanic("JNICALL CallShortMethod()!");
    return 0;
}
jshort CallShortMethodV(JNIEnv *env, jobject obj, jmethodID methodID, va_list args) {
    GoPanic("JNICALL CallShortMethodV()!");
    return 0;
}
jshort CallShortMethodA(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args) {
    GoPanic("JNICALL CallShortMethodA()!");
    return 0;
}
jint CallIntMethod(JNIEnv *env, jobject obj, jmethodID methodID, ...) {
    GoPanic("JNICALL CallIntMethod()!");
    return 0;
}
jint CallIntMethodV(JNIEnv *env, jobject obj, jmethodID methodID, va_list args) {
    GoPanic("JNICALL CallIntMethodV()!");
    return 0;
}
jint CallIntMethodA(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args) {
    GoPanic("JNICALL CallIntMethodA()!");
    return 0;
}
jlong CallLongMethod(JNIEnv *env, jobject obj, jmethodID methodID, ...) {
    GoPanic("JNICALL CallLongMethod()!");
    return 0;
}
jlong CallLongMethodV(JNIEnv *env, jobject obj, jmethodID methodID, va_list args) {
    GoPanic("JNICALL CallLongMethodV()!");
    return 0;
}
jlong CallLongMethodA(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args) {
    GoPanic("JNICALL CallLongMethodA()!");
    return 0;
}
jfloat CallFloatMethod(JNIEnv *env, jobject obj, jmethodID methodID, ...) {
    GoPanic("JNICALL CallFloatMethod()!");
    return 0;
}
jfloat CallFloatMethodV(JNIEnv *env, jobject obj, jmethodID methodID, va_list args) {
    GoPanic("JNICALL CallFloatMethodV()!");
    return 0;
}
jfloat CallFloatMethodA(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args) {
    GoPanic("JNICALL CallFloatMethodA()!");
    return 0;
}
jdouble CallDoubleMethod(JNIEnv *env, jobject obj, jmethodID methodID, ...) {
    GoPanic("JNICALL CallDoubleMethod()!");
    return 0;
}
jdouble CallDoubleMethodV(JNIEnv *env, jobject obj, jmethodID methodID, va_list args) {
    GoPanic("JNICALL CallDoubleMethodV()!");
    return 0;
}
jdouble CallDoubleMethodA(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args) {
    GoPanic("JNICALL CallDoubleMethodA()!");
    return 0;
}
void CallVoidMethod(JNIEnv *env, jobject obj, jmethodID methodID, ...) {
    GoPanic("JNICALL CallVoidMethod()!");
}
void CallVoidMethodV(JNIEnv *env, jobject obj, jmethodID methodID, va_list args) {
    GoPanic("JNICALL CallVoidMethodV()!");
}
void CallVoidMethodA(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue * args) {
    GoPanic("JNICALL CallVoidMethodA()!");
}
jobject CallNonvirtualObjectMethod(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...) {
    GoPanic("JNICALL CallNonvirtualObjectMethod()!");
    return 0;
}
jobject CallNonvirtualObjectMethodV(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args) {
    GoPanic("JNICALL CallNonvirtualObjectMethodV()!");
    return 0;
}
jobject CallNonvirtualObjectMethodA(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue * args) {
    GoPanic("JNICALL CallNonvirtualObjectMethodA()!");
    return 0;
}
jboolean CallNonvirtualBooleanMethod(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...) {
    GoPanic("JNICALL CallNonvirtualBooleanMethod()!");
    return 0;
}
jboolean CallNonvirtualBooleanMethodV(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args) {
    GoPanic("JNICALL CallNonvirtualBooleanMethodV()!");
    return 0;
}
jboolean CallNonvirtualBooleanMethodA(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue * args) {
    GoPanic("JNICALL CallNonvirtualBooleanMethodA()!");
    return 0;
}
jbyte CallNonvirtualByteMethod(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...) {
    GoPanic("JNICALL CallNonvirtualByteMethod()!");
    return 0;
}
jbyte CallNonvirtualByteMethodV(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args) {
    GoPanic("JNICALL CallNonvirtualByteMethodV()!");
    return 0;
}
jbyte CallNonvirtualByteMethodA(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue *args) {
    GoPanic("JNICALL CallNonvirtualByteMethodA()!");
    return 0;
}
jchar CallNonvirtualCharMethod(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...) {
    GoPanic("JNICALL CallNonvirtualCharMethod()!");
    return 0;
}
jchar CallNonvirtualCharMethodV(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args) {
    GoPanic("JNICALL CallNonvirtualCharMethodV()!");
    return 0;
}
jchar CallNonvirtualCharMethodA(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue *args) {
    GoPanic("JNICALL CallNonvirtualCharMethodA()!");
    return 0;
}
jshort CallNonvirtualShortMethod(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...) {
    GoPanic("JNICALL CallNonvirtualShortMethod()!");
    return 0;
}
jshort CallNonvirtualShortMethodV(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args) {
    GoPanic("JNICALL CallNonvirtualShortMethodV()!");
    return 0;
}
jshort CallNonvirtualShortMethodA(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue *args) {
    GoPanic("JNICALL CallNonvirtualShortMethodA()!");
    return 0;
}
jint CallNonvirtualIntMethod(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...) {
    GoPanic("JNICALL CallNonvirtualIntMethod()!");
    return 0;
}
jint CallNonvirtualIntMethodV(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args) {
    GoPanic("JNICALL CallNonvirtualIntMethodV()!");
    return 0;
}
jint CallNonvirtualIntMethodA(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue *args) {
    GoPanic("JNICALL CallNonvirtualIntMethodA()!");
    return 0;
}
jlong CallNonvirtualLongMethod(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...) {
    GoPanic("JNICALL CallNonvirtualLongMethod()!");
    return 0;
}
jlong CallNonvirtualLongMethodV(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args) {
    GoPanic("JNICALL CallNonvirtualLongMethodV()!");
    return 0;
}
jlong CallNonvirtualLongMethodA(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue *args) {
    GoPanic("JNICALL CallNonvirtualLongMethodA()!");
    return 0;
}
jfloat CallNonvirtualFloatMethod(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...) {
    GoPanic("JNICALL CallNonvirtualFloatMethod()!");
    return 0;
}
jfloat CallNonvirtualFloatMethodV(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args) {
    GoPanic("JNICALL CallNonvirtualFloatMethodV()!");
    return 0;
}
jfloat CallNonvirtualFloatMethodA(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue *args) {
    GoPanic("JNICALL CallNonvirtualFloatMethodA()!");
    return 0;
}
jdouble CallNonvirtualDoubleMethod(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...) {
    GoPanic("JNICALL CallNonvirtualDoubleMethod()!");
    return 0;
}
jdouble CallNonvirtualDoubleMethodV(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args) {
    GoPanic("JNICALL CallNonvirtualDoubleMethodV()!");
    return 0;
}
jdouble CallNonvirtualDoubleMethodA(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue *args) {
    GoPanic("JNICALL CallNonvirtualDoubleMethodA()!");
    return 0;
}
void CallNonvirtualVoidMethod(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...) {
    GoPanic("JNICALL CallNonvirtualVoidMethod()!");
}
void CallNonvirtualVoidMethodV(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args) {
    GoPanic("JNICALL CallNonvirtualVoidMethodV()!");
}
void CallNonvirtualVoidMethodA(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue * args) {
    GoPanic("JNICALL CallNonvirtualVoidMethodA()!");
}
jfieldID GetFieldID(JNIEnv *env, jclass clazz, const char *name, const char *sig) {
    GoPanic("JNICALL GetFieldID()!");
    return 0;
}
jobject GetObjectField(JNIEnv *env, jobject obj, jfieldID fieldID) {
    GoPanic("JNICALL GetObjectField()!");
    return 0;
}
jboolean GetBooleanField(JNIEnv *env, jobject obj, jfieldID fieldID) {
    GoPanic("JNICALL GetBooleanField()!");
    return 0;
}
jbyte GetByteField(JNIEnv *env, jobject obj, jfieldID fieldID) {
    GoPanic("JNICALL GetByteField()!");
    return 0;
}
jchar GetCharField(JNIEnv *env, jobject obj, jfieldID fieldID) {
    GoPanic("JNICALL GetCharField()!");
    return 0;
}
jshort GetShortField(JNIEnv *env, jobject obj, jfieldID fieldID) {
    GoPanic("JNICALL GetShortField()!");
    return 0;
}
jint GetIntField(JNIEnv *env, jobject obj, jfieldID fieldID) {
    GoPanic("JNICALL GetIntField()!");
    return 0;
}
jlong GetLongField(JNIEnv *env, jobject obj, jfieldID fieldID) {
    GoPanic("JNICALL GetLongField()!");
    return 0;
}
jfloat GetFloatField(JNIEnv *env, jobject obj, jfieldID fieldID) {
    GoPanic("JNICALL GetFloatField()!");
    return 0;
}
jdouble GetDoubleField(JNIEnv *env, jobject obj, jfieldID fieldID) {
    GoPanic("JNICALL GetDoubleField()!");
    return 0;
}
void SetObjectField(JNIEnv *env, jobject obj, jfieldID fieldID, jobject val) {
    GoPanic("JNICALL SetObjectField()!");
}
void SetBooleanField(JNIEnv *env, jobject obj, jfieldID fieldID, jboolean val) {
    GoPanic("JNICALL SetBooleanField()!");
}
void SetByteField(JNIEnv *env, jobject obj, jfieldID fieldID, jbyte val) {
    GoPanic("JNICALL SetByteField()!");
}
void SetCharField(JNIEnv *env, jobject obj, jfieldID fieldID, jchar val) {
    GoPanic("JNICALL SetCharField()!");
}
void SetShortField(JNIEnv *env, jobject obj, jfieldID fieldID, jshort val) {
    GoPanic("JNICALL SetShortField()!");
}
void SetIntField(JNIEnv *env, jobject obj, jfieldID fieldID, jint val) {
    GoPanic("JNICALL SetIntField()!");
}
void SetLongField(JNIEnv *env, jobject obj, jfieldID fieldID, jlong val) {
    GoPanic("JNICALL SetLongField()!");
}
void SetFloatField(JNIEnv *env, jobject obj, jfieldID fieldID, jfloat val) {
    GoPanic("JNICALL SetFloatField()!");
}
void SetDoubleField(JNIEnv *env, jobject obj, jfieldID fieldID, jdouble val) {
    GoPanic("JNICALL SetDoubleField()!");
}

jmethodID GetStaticMethodID(JNIEnv *env, jclass clazz, const char *name, const char *sig) {
    //GoPanic("JNICALL GetStaticMethodID()!");
    return GoGetStaticMethodID(env, clazz, name, sig);
}
jobject CallStaticObjectMethod(JNIEnv *env, jclass clazz, jmethodID methodID, ...) {
    GoPanic("JNICALL CallStaticObjectMethod()!");
    return 0;
}
jobject CallStaticObjectMethodV(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args) {
    //GoPanic("JNICALL CallStaticObjectMethodV()!");
    int argCount = GoGetParamCount(clazz, methodID);
    jvalue argArr[argCount];
    copyVaListToArray(args, argCount, argArr);
    return GoCallStaticObjectMethodA(env, clazz, methodID, argArr);
}
jobject CallStaticObjectMethodA(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args) {
    //GoPanic("JNICALL CallStaticObjectMethodA()!");
    return GoCallStaticObjectMethodA(env, clazz, methodID, args);
}
jboolean CallStaticBooleanMethod(JNIEnv *env, jclass clazz, jmethodID methodID, ...) {
    GoPanic("JNICALL CallStaticBooleanMethod()!");
    return 0;
}
jboolean CallStaticBooleanMethodV(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args) {
    GoPanic("JNICALL CallStaticBooleanMethodV()!");
    return 0;
}
jboolean CallStaticBooleanMethodA(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args) {
    GoPanic("JNICALL CallStaticBooleanMethodA()!");
    return 0;
}
jbyte CallStaticByteMethod(JNIEnv *env, jclass clazz, jmethodID methodID, ...) {
    GoPanic("JNICALL CallStaticByteMethod()!");
    return 0;
}
jbyte CallStaticByteMethodV(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args) {
    GoPanic("JNICALL CallStaticByteMethodV()!");
    return 0;
}
jbyte CallStaticByteMethodA(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args) {
    GoPanic("JNICALL CallStaticByteMethodA()!");
    return 0;
}
jchar CallStaticCharMethod(JNIEnv *env, jclass clazz, jmethodID methodID, ...) {
    GoPanic("JNICALL CallStaticCharMethod()!");
    return 0;
}
jchar CallStaticCharMethodV(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args) {
    GoPanic("JNICALL CallStaticCharMethodV()!");
    return 0;
}
jchar CallStaticCharMethodA(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args) {
    GoPanic("JNICALL CallStaticCharMethodA()!");
    return 0;
}
jshort CallStaticShortMethod(JNIEnv *env, jclass clazz, jmethodID methodID, ...) {
    GoPanic("JNICALL CallStaticShortMethod()!");
    return 0;
}
jshort CallStaticShortMethodV(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args) {
    GoPanic("JNICALL CallStaticShortMethodV()!");
    return 0;
}
jshort CallStaticShortMethodA(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args) {
    GoPanic("JNICALL CallStaticShortMethodA()!");
    return 0;
}
jint CallStaticIntMethod(JNIEnv *env, jclass clazz, jmethodID methodID, ...) {
    GoPanic("JNICALL CallStaticIntMethod()!");
    return 0;
}
jint CallStaticIntMethodV(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args) {
    GoPanic("JNICALL CallStaticIntMethodV()!");
    return 0;
}
jint CallStaticIntMethodA(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args) {
    GoPanic("JNICALL CallStaticIntMethodA()!");
    return 0;
}
jlong CallStaticLongMethod(JNIEnv *env, jclass clazz, jmethodID methodID, ...) {
    GoPanic("JNICALL CallStaticLongMethod()!");
    return 0;
}
jlong CallStaticLongMethodV(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args) {
    GoPanic("JNICALL CallStaticLongMethodV()!");
    return 0;
}
jlong CallStaticLongMethodA(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args) {
    GoPanic("JNICALL CallStaticLongMethodA()!");
    return 0;
}
jfloat CallStaticFloatMethod(JNIEnv *env, jclass clazz, jmethodID methodID, ...) {
    GoPanic("JNICALL CallStaticFloatMethod()!");
    return 0;
}
jfloat CallStaticFloatMethodV(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args) {
    GoPanic("JNICALL CallStaticFloatMethodV()!");
    return 0;
}
jfloat CallStaticFloatMethodA(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args) {
    GoPanic("JNICALL CallStaticFloatMethodA()!");
    return 0;
}
jdouble CallStaticDoubleMethod(JNIEnv *env, jclass clazz, jmethodID methodID, ...) {
    GoPanic("JNICALL CallStaticDoubleMethod()!");
    return 0;
}
jdouble CallStaticDoubleMethodV(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args) {
    GoPanic("JNICALL CallStaticDoubleMethodV()!");
    return 0;
}
jdouble CallStaticDoubleMethodA(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args) {
    GoPanic("JNICALL CallStaticDoubleMethodA()!");
    return 0;
}
void CallStaticVoidMethod(JNIEnv *env, jclass cls, jmethodID methodID, ...) {
    GoPanic("JNICALL CallStaticVoidMethod()!");
}
void CallStaticVoidMethodV(JNIEnv *env, jclass cls, jmethodID methodID, va_list args) {
    GoPanic("JNICALL CallStaticVoidMethodV()!");
}
void CallStaticVoidMethodA(JNIEnv *env, jclass cls, jmethodID methodID, const jvalue * args) {
    GoPanic("JNICALL CallStaticVoidMethodA()!");
}

jfieldID GetStaticFieldID(JNIEnv *env, jclass clazz, const char *name, const char *sig) {
    GoPanic("JNICALL GetStaticFieldID()!");
    return 0;
}
jobject GetStaticObjectField(JNIEnv *env, jclass clazz, jfieldID fieldID) {
    GoPanic("JNICALL GetStaticObjectField()!");
    return 0;
}
jboolean GetStaticBooleanField(JNIEnv *env, jclass clazz, jfieldID fieldID) {
    GoPanic("JNICALL GetStaticBooleanField()!");
    return 0;
}
jbyte GetStaticByteField(JNIEnv *env, jclass clazz, jfieldID fieldID) {
    GoPanic("JNICALL GetStaticByteField()!");
    return 0;
}
jchar GetStaticCharField(JNIEnv *env, jclass clazz, jfieldID fieldID) {
    GoPanic("JNICALL GetStaticCharField()!");
    return 0;
}
jshort GetStaticShortField(JNIEnv *env, jclass clazz, jfieldID fieldID) {
    GoPanic("JNICALL GetStaticShortField()!");
    return 0;
}
jint GetStaticIntField(JNIEnv *env, jclass clazz, jfieldID fieldID) {
    GoPanic("JNICALL GetStaticIntField()!");
    return 0;
}
jlong GetStaticLongField(JNIEnv *env, jclass clazz, jfieldID fieldID) {
    GoPanic("JNICALL GetStaticLongField()!");
    return 0;
}
jfloat GetStaticFloatField(JNIEnv *env, jclass clazz, jfieldID fieldID) {
    GoPanic("JNICALL GetStaticFloatField()!");
    return 0;
}
jdouble GetStaticDoubleField(JNIEnv *env, jclass clazz, jfieldID fieldID) {
    GoPanic("JNICALL GetStaticDoubleField()!");
    return 0;
}
void SetStaticObjectField(JNIEnv *env, jclass clazz, jfieldID fieldID, jobject value) {
    GoPanic("JNICALL SetStaticObjectField()!");
}
void SetStaticBooleanField(JNIEnv *env, jclass clazz, jfieldID fieldID, jboolean value) {
    GoPanic("JNICALL SetStaticBooleanField()!");
}
void SetStaticByteField(JNIEnv *env, jclass clazz, jfieldID fieldID, jbyte value) {
    GoPanic("JNICALL SetStaticByteField()!");
}
void SetStaticCharField(JNIEnv *env, jclass clazz, jfieldID fieldID, jchar value) {
    GoPanic("JNICALL SetStaticCharField()!");
}
void SetStaticShortField(JNIEnv *env, jclass clazz, jfieldID fieldID, jshort value) {
    GoPanic("JNICALL SetStaticShortField()!");
}
void SetStaticIntField(JNIEnv *env, jclass clazz, jfieldID fieldID, jint value) {
    GoPanic("JNICALL SetStaticIntField()!");
}
void SetStaticLongField(JNIEnv *env, jclass clazz, jfieldID fieldID, jlong value) {
    GoPanic("JNICALL SetStaticLongField()!");
}
void SetStaticFloatField(JNIEnv *env, jclass clazz, jfieldID fieldID, jfloat value) {
    GoPanic("JNICALL SetStaticFloatField()!");
}
void SetStaticDoubleField(JNIEnv *env, jclass clazz, jfieldID fieldID, jdouble value) {
    GoPanic("JNICALL SetStaticDoubleField()!");
}

jstring NewString(JNIEnv *env, const jchar *unicode, jsize len) {
    //GoPanic("JNICALL NewString()!");
    return GoNewString(env, unicode, len);
}
jsize GetStringLength(JNIEnv *env, jstring str) {
    //GoPanic("JNICALL GetStringLength()!");
    return GoGetStringLength(env, str);
}
const jchar *GetStringChars(JNIEnv *env, jstring str, jboolean *isCopy) {
    GoPanic("JNICALL GetStringChars()!");
    return 0;
}
void ReleaseStringChars(JNIEnv *env, jstring str, const jchar *chars) {
    GoPanic("JNICALL ReleaseStringChars()!");
}

jstring NewStringUTF(JNIEnv *env, const char *utf) {
    //GoPanic("JNICALL NewStringUTF()!");
    return GoNewStringUTF(env, utf);
}
jsize GetStringUTFLength(JNIEnv *env, jstring str) {
    GoPanic("JNICALL GetStringUTFLength()!");
    return 0;
}
const char* GetStringUTFChars(JNIEnv *env, jstring str, jboolean *isCopy) {
    GoPanic("JNICALL GetStringUTFChars()!");
    return 0;
}
void ReleaseStringUTFChars(JNIEnv *env, jstring str, const char* chars) {
    GoPanic("JNICALL ReleaseStringUTFChars()!");
}

jsize GetArrayLength(JNIEnv *env, jarray array) {
    GoPanic("JNICALL GetArrayLength()!");
    return 0;
}
jobjectArray NewObjectArray(JNIEnv *env, jsize len, jclass clazz, jobject init) {
    GoPanic("JNICALL NewObjectArray()!");
    return 0;
}
jobject GetObjectArrayElement(JNIEnv *env, jobjectArray array, jsize index) {
    GoPanic("JNICALL GetObjectArrayElement()!");
    return 0;
}
void SetObjectArrayElement(JNIEnv *env, jobjectArray array, jsize index, jobject val) {
    GoPanic("JNICALL SetObjectArrayElement()!");
}
jbooleanArray NewBooleanArray(JNIEnv *env, jsize len) {
    GoPanic("JNICALL NewBooleanArray()!");
    return 0;
}
jbyteArray NewByteArray(JNIEnv *env, jsize len) {
    GoPanic("JNICALL NewByteArray()!");
    return 0;
}
jcharArray NewCharArray(JNIEnv *env, jsize len) {
    GoPanic("JNICALL NewCharArray()!");
    return 0;
}
jshortArray NewShortArray(JNIEnv *env, jsize len) {
    GoPanic("JNICALL NewShortArray()!");
    return 0;
}
jintArray NewIntArray(JNIEnv *env, jsize len) {
    GoPanic("JNICALL NewIntArray()!");
    return 0;
}
jlongArray NewLongArray(JNIEnv *env, jsize len) {
    GoPanic("JNICALL NewLongArray()!");
    return 0;
}
jfloatArray NewFloatArray(JNIEnv *env, jsize len) {
    GoPanic("JNICALL NewFloatArray()!");
    return 0;
}
jdoubleArray NewDoubleArray(JNIEnv *env, jsize len) {
    GoPanic("JNICALL NewDoubleArray()!");
    return 0;
}
jboolean * GetBooleanArrayElements(JNIEnv *env, jbooleanArray array, jboolean *isCopy) {
    GoPanic("JNICALL GetBooleanArrayElements()!");
    return 0;
}
jbyte * GetByteArrayElements(JNIEnv *env, jbyteArray array, jboolean *isCopy) {
    GoPanic("JNICALL GetByteArrayElements()!");
    return 0;
}
jchar * GetCharArrayElements(JNIEnv *env, jcharArray array, jboolean *isCopy) {
    GoPanic("JNICALL GetCharArrayElements()!");
    return 0;
}
jshort * GetShortArrayElements(JNIEnv *env, jshortArray array, jboolean *isCopy) {
    GoPanic("JNICALL GetShortArrayElements()!");
    return 0;
}
jint * GetIntArrayElements(JNIEnv *env, jintArray array, jboolean *isCopy) {
    GoPanic("JNICALL GetIntArrayElements()!");
    return 0;
}
jlong * GetLongArrayElements(JNIEnv *env, jlongArray array, jboolean *isCopy) {
    GoPanic("JNICALL GetLongArrayElements()!");
    return 0;
}
jfloat * GetFloatArrayElements(JNIEnv *env, jfloatArray array, jboolean *isCopy) {
    GoPanic("JNICALL GetFloatArrayElements()!");
    return 0;
}
jdouble * GetDoubleArrayElements(JNIEnv *env, jdoubleArray array, jboolean *isCopy) {
    GoPanic("JNICALL GetDoubleArrayElements()!");
    return 0;
}
void ReleaseBooleanArrayElements(JNIEnv *env, jbooleanArray array, jboolean *elems, jint mode) {
    GoPanic("JNICALL ReleaseBooleanArrayElements()!");
}
void ReleaseByteArrayElements(JNIEnv *env, jbyteArray array, jbyte *elems, jint mode) {
    GoPanic("JNICALL ReleaseByteArrayElements()!");
}
void ReleaseCharArrayElements(JNIEnv *env, jcharArray array, jchar *elems, jint mode) {
    GoPanic("JNICALL ReleaseCharArrayElements()!");
}
void ReleaseShortArrayElements(JNIEnv *env, jshortArray array, jshort *elems, jint mode) {
    GoPanic("JNICALL ReleaseShortArrayElements()!");
}
void ReleaseIntArrayElements(JNIEnv *env, jintArray array, jint *elems, jint mode) {
    GoPanic("JNICALL ReleaseIntArrayElements()!");
}
void ReleaseLongArrayElements(JNIEnv *env, jlongArray array, jlong *elems, jint mode) {
    GoPanic("JNICALL ReleaseLongArrayElements()!");
}
void ReleaseFloatArrayElements(JNIEnv *env, jfloatArray array, jfloat *elems, jint mode) {
    GoPanic("JNICALL ReleaseFloatArrayElements()!");
}
void ReleaseDoubleArrayElements(JNIEnv *env, jdoubleArray array, jdouble *elems, jint mode) {
    GoPanic("JNICALL ReleaseDoubleArrayElements()!");
}
void GetBooleanArrayRegion(JNIEnv *env, jbooleanArray array, jsize start, jsize l, jboolean *buf) {
    GoPanic("JNICALL GetBooleanArrayRegion()!");
}
void GetByteArrayRegion(JNIEnv *env, jbyteArray array, jsize start, jsize len, jbyte *buf) {
    GoPanic("JNICALL GetByteArrayRegion()!");
}
void GetCharArrayRegion(JNIEnv *env, jcharArray array, jsize start, jsize len, jchar *buf) {
    GoPanic("JNICALL GetCharArrayRegion()!");
}
void GetShortArrayRegion(JNIEnv *env, jshortArray array, jsize start, jsize len, jshort *buf) {
    GoPanic("JNICALL GetShortArrayRegion()!");
}
void GetIntArrayRegion(JNIEnv *env, jintArray array, jsize start, jsize len, jint *buf) {
    GoPanic("JNICALL GetIntArrayRegion()!");
}
void GetLongArrayRegion(JNIEnv *env, jlongArray array, jsize start, jsize len, jlong *buf) {
    GoPanic("JNICALL GetLongArrayRegion()!");
}
void GetFloatArrayRegion(JNIEnv *env, jfloatArray array, jsize start, jsize len, jfloat *buf) {
    GoPanic("JNICALL GetFloatArrayRegion()!");
}
void GetDoubleArrayRegion(JNIEnv *env, jdoubleArray array, jsize start, jsize len, jdouble *buf) {
    GoPanic("JNICALL GetDoubleArrayRegion()!");
}
void SetBooleanArrayRegion(JNIEnv *env, jbooleanArray array, jsize start, jsize l, const jboolean *buf) {
    GoPanic("JNICALL SetBooleanArrayRegion()!");
}
void SetByteArrayRegion(JNIEnv *env, jbyteArray array, jsize start, jsize len, const jbyte *buf) {
    GoPanic("JNICALL SetByteArrayRegion()!");
}
void SetCharArrayRegion(JNIEnv *env, jcharArray array, jsize start, jsize len, const jchar *buf) {
    GoPanic("JNICALL SetCharArrayRegion()!");
}
void SetShortArrayRegion(JNIEnv *env, jshortArray array, jsize start, jsize len, const jshort *buf) {
    GoPanic("JNICALL SetShortArrayRegion()!");
}
void SetIntArrayRegion(JNIEnv *env, jintArray array, jsize start, jsize len, const jint *buf) {
    GoPanic("JNICALL SetIntArrayRegion()!");
}
void SetLongArrayRegion(JNIEnv *env, jlongArray array, jsize start, jsize len, const jlong *buf) {
    GoPanic("JNICALL SetLongArrayRegion()!");
}
void SetFloatArrayRegion(JNIEnv *env, jfloatArray array, jsize start, jsize len, const jfloat *buf) {
    GoPanic("JNICALL SetFloatArrayRegion()!");
}
void SetDoubleArrayRegion(JNIEnv *env, jdoubleArray array, jsize start, jsize len, const jdouble *buf) {
    GoPanic("JNICALL SetDoubleArrayRegion()!");
}
jint RegisterNatives(JNIEnv *env, jclass clazz, const JNINativeMethod *methods, jint nMethods) {
    GoPanic("JNICALL RegisterNatives()!");
    return 0;
}
jint UnregisterNatives(JNIEnv *env, jclass clazz) {
    GoPanic("JNICALL UnregisterNatives()!");
    return 0;
}
jint MonitorEnter(JNIEnv *env, jobject obj) {
    GoPanic("JNICALL MonitorEnter()!");
    return 0;
}
jint MonitorExit(JNIEnv *env, jobject obj) {
    GoPanic("JNICALL MonitorExit()!");
    return 0;
}
jint GetJavaVM(JNIEnv *env, JavaVM **vm) {
    GoPanic("JNICALL GetJavaVM()!");
    return 0;
}

void GetStringRegion(JNIEnv *env, jstring str, jsize start, jsize len, jchar *buf) {
    GoPanic("JNICALL GetStringRegion()!");
}
void GetStringUTFRegion(JNIEnv *env, jstring str, jsize start, jsize len, char *buf) {
    GoPanic("JNICALL GetStringUTFRegion()!");
}

void * GetPrimitiveArrayCritical(JNIEnv *env, jarray array, jboolean *isCopy) {
    GoPanic("JNICALL GetPrimitiveArrayCritical()!");
    return 0;
}
void ReleasePrimitiveArrayCritical(JNIEnv *env, jarray array, void *carray, jint mode) {
    GoPanic("JNICALL ReleasePrimitiveArrayCritical()!");
}

const jchar * GetStringCritical(JNIEnv *env, jstring string, jboolean *isCopy) {
    //GoPanic("JNICALL GetStringCritical()!");
    return GoGetStringCritical(env, string, isCopy);
}
void ReleaseStringCritical(JNIEnv *env, jstring string, const jchar *cstring) {
    //GoPanic("JNICALL ReleaseStringCritical()!");
    GoReleaseStringCritical(env, string, cstring);
}

jweak NewWeakGlobalRef(JNIEnv *env, jobject obj) {
    GoPanic("JNICALL NewWeakGlobalRef()!");
    return 0;
}
void DeleteWeakGlobalRef(JNIEnv *env, jweak ref) {
    GoPanic("JNICALL DeleteWeakGlobalRef()!");
}

jboolean ExceptionCheck(JNIEnv *env) {
    //GoPanic("JNICALL ExceptionCheck()!");
    return GoExceptionCheck(env);
}

jobject NewDirectByteBuffer(JNIEnv* env, void* address, jlong capacity) {
    GoPanic("JNICALL NewDirectByteBuffer()!");
    return 0;
}
void* GetDirectBufferAddress(JNIEnv* env, jobject buf) {
    GoPanic("JNICALL GetDirectBufferAddress()!");
    return 0;
}
jlong GetDirectBufferCapacity(JNIEnv* env, jobject buf) {
    GoPanic("JNICALL GetDirectBufferCapacity()!");
    return 0;
}
jobjectRefType GetObjectRefType(JNIEnv* env, jobject obj) {
    GoPanic("JNICALL GetObjectRefType()!");
    return 0;
}

JNIEnv NewJNIEnvWrapper() {
    struct JNINativeInterface_ * myEnv = malloc(sizeof(struct JNINativeInterface_));
    myEnv->GetVersion = GetVersion;
    myEnv->DefineClass = DefineClass;
    myEnv->FindClass = FindClass;
    myEnv->FromReflectedMethod = FromReflectedMethod;
    myEnv->FromReflectedField = FromReflectedField;
    myEnv->ToReflectedMethod = ToReflectedMethod;
    myEnv->GetSuperclass = GetSuperclass;
    myEnv->IsAssignableFrom = IsAssignableFrom;
    myEnv->ToReflectedField = ToReflectedField;
    myEnv->Throw = Throw;
    myEnv->ThrowNew = ThrowNew;
    myEnv->ExceptionOccurred = ExceptionOccurred;
    myEnv->ExceptionDescribe = ExceptionDescribe;
    myEnv->ExceptionClear = ExceptionClear;
    myEnv->FatalError = FatalError;
    myEnv->PushLocalFrame = PushLocalFrame;
    myEnv->PopLocalFrame = PopLocalFrame;
    myEnv->NewGlobalRef = NewGlobalRef;
    myEnv->DeleteGlobalRef = DeleteGlobalRef;
    myEnv->DeleteLocalRef = DeleteLocalRef;
    myEnv->IsSameObject = IsSameObject;
    myEnv->NewLocalRef = NewLocalRef;
    myEnv->EnsureLocalCapacity = EnsureLocalCapacity;
    myEnv->AllocObject = AllocObject;
    myEnv->NewObject = NewObject;
    myEnv->NewObjectV = NewObjectV;
    myEnv->NewObjectA = NewObjectA;
    myEnv->GetObjectClass = GetObjectClass;
    myEnv->IsInstanceOf = IsInstanceOf;
    myEnv->GetMethodID = GetMethodID;
    myEnv->CallObjectMethod = CallObjectMethod;
    myEnv->CallObjectMethodV = CallObjectMethodV;
    myEnv->CallObjectMethodA = CallObjectMethodA;
    myEnv->CallBooleanMethod = CallBooleanMethod;
    myEnv->CallBooleanMethodV = CallBooleanMethodV;
    myEnv->CallBooleanMethodA = CallBooleanMethodA;
    myEnv->CallByteMethod = CallByteMethod;
    myEnv->CallByteMethodV = CallByteMethodV;
    myEnv->CallByteMethodA = CallByteMethodA;
    myEnv->CallCharMethod = CallCharMethod;
    myEnv->CallCharMethodV = CallCharMethodV;
    myEnv->CallCharMethodA = CallCharMethodA;
    myEnv->CallShortMethod = CallShortMethod;
    myEnv->CallShortMethodV = CallShortMethodV;
    myEnv->CallShortMethodA = CallShortMethodA;
    myEnv->CallIntMethod = CallIntMethod;
    myEnv->CallIntMethodV = CallIntMethodV;
    myEnv->CallIntMethodA = CallIntMethodA;
    myEnv->CallLongMethod = CallLongMethod;
    myEnv->CallLongMethodV = CallLongMethodV;
    myEnv->CallLongMethodA = CallLongMethodA;
    myEnv->CallFloatMethod = CallFloatMethod;
    myEnv->CallFloatMethodV = CallFloatMethodV;
    myEnv->CallFloatMethodA = CallFloatMethodA;
    myEnv->CallDoubleMethod = CallDoubleMethod;
    myEnv->CallDoubleMethodV = CallDoubleMethodV;
    myEnv->CallDoubleMethodA = CallDoubleMethodA;
    myEnv->CallVoidMethod = CallVoidMethod;
    myEnv->CallVoidMethodV = CallVoidMethodV;
    myEnv->CallVoidMethodA = CallVoidMethodA;
    myEnv->CallNonvirtualObjectMethod = CallNonvirtualObjectMethod;
    myEnv->CallNonvirtualObjectMethodV = CallNonvirtualObjectMethodV;
    myEnv->CallNonvirtualObjectMethodA = CallNonvirtualObjectMethodA;
    myEnv->CallNonvirtualBooleanMethod = CallNonvirtualBooleanMethod;
    myEnv->CallNonvirtualBooleanMethodV = CallNonvirtualBooleanMethodV;
    myEnv->CallNonvirtualBooleanMethodA = CallNonvirtualBooleanMethodA;
    myEnv->CallNonvirtualByteMethod = CallNonvirtualByteMethod;
    myEnv->CallNonvirtualByteMethodV = CallNonvirtualByteMethodV;
    myEnv->CallNonvirtualByteMethodA = CallNonvirtualByteMethodA;
    myEnv->CallNonvirtualCharMethod = CallNonvirtualCharMethod;
    myEnv->CallNonvirtualCharMethodV = CallNonvirtualCharMethodV;
    myEnv->CallNonvirtualCharMethodA = CallNonvirtualCharMethodA;
    myEnv->CallNonvirtualShortMethod = CallNonvirtualShortMethod;
    myEnv->CallNonvirtualShortMethodV = CallNonvirtualShortMethodV;
    myEnv->CallNonvirtualShortMethodA = CallNonvirtualShortMethodA;
    myEnv->CallNonvirtualIntMethod = CallNonvirtualIntMethod;
    myEnv->CallNonvirtualIntMethodV = CallNonvirtualIntMethodV;
    myEnv->CallNonvirtualIntMethodA = CallNonvirtualIntMethodA;
    myEnv->CallNonvirtualLongMethod = CallNonvirtualLongMethod;
    myEnv->CallNonvirtualLongMethodV = CallNonvirtualLongMethodV;
    myEnv->CallNonvirtualLongMethodA = CallNonvirtualLongMethodA;
    myEnv->CallNonvirtualFloatMethod = CallNonvirtualFloatMethod;
    myEnv->CallNonvirtualFloatMethodV = CallNonvirtualFloatMethodV;
    myEnv->CallNonvirtualFloatMethodA = CallNonvirtualFloatMethodA;
    myEnv->CallNonvirtualDoubleMethod = CallNonvirtualDoubleMethod;
    myEnv->CallNonvirtualDoubleMethodV = CallNonvirtualDoubleMethodV;
    myEnv->CallNonvirtualDoubleMethodA = CallNonvirtualDoubleMethodA;
    myEnv->CallNonvirtualVoidMethod = CallNonvirtualVoidMethod;
    myEnv->CallNonvirtualVoidMethodV = CallNonvirtualVoidMethodV;
    myEnv->CallNonvirtualVoidMethodA = CallNonvirtualVoidMethodA;
    myEnv->GetFieldID = GetFieldID;
    myEnv->GetObjectField = GetObjectField;
    myEnv->GetBooleanField = GetBooleanField;
    myEnv->GetByteField = GetByteField;
    myEnv->GetCharField = GetCharField;
    myEnv->GetShortField = GetShortField;
    myEnv->GetIntField = GetIntField;
    myEnv->GetLongField = GetLongField;
    myEnv->GetFloatField = GetFloatField;
    myEnv->GetDoubleField = GetDoubleField;
    myEnv->SetObjectField = SetObjectField;
    myEnv->SetBooleanField = SetBooleanField;
    myEnv->SetByteField = SetByteField;
    myEnv->SetCharField = SetCharField;
    myEnv->SetShortField = SetShortField;
    myEnv->SetIntField = SetIntField;
    myEnv->SetLongField = SetLongField;
    myEnv->SetFloatField = SetFloatField;
    myEnv->SetDoubleField = SetDoubleField;
    myEnv->GetStaticMethodID = GetStaticMethodID;
    myEnv->CallStaticObjectMethod = CallStaticObjectMethod;
    myEnv->CallStaticObjectMethodV = CallStaticObjectMethodV;
    myEnv->CallStaticObjectMethodA = CallStaticObjectMethodA;
    myEnv->CallStaticBooleanMethod = CallStaticBooleanMethod;
    myEnv->CallStaticBooleanMethodV = CallStaticBooleanMethodV;
    myEnv->CallStaticBooleanMethodA = CallStaticBooleanMethodA;
    myEnv->CallStaticByteMethod = CallStaticByteMethod;
    myEnv->CallStaticByteMethodV = CallStaticByteMethodV;
    myEnv->CallStaticByteMethodA = CallStaticByteMethodA;
    myEnv->CallStaticCharMethod = CallStaticCharMethod;
    myEnv->CallStaticCharMethodV = CallStaticCharMethodV;
    myEnv->CallStaticCharMethodA = CallStaticCharMethodA;
    myEnv->CallStaticShortMethod = CallStaticShortMethod;
    myEnv->CallStaticShortMethodV = CallStaticShortMethodV;
    myEnv->CallStaticShortMethodA = CallStaticShortMethodA;
    myEnv->CallStaticIntMethod = CallStaticIntMethod;
    myEnv->CallStaticIntMethodV = CallStaticIntMethodV;
    myEnv->CallStaticIntMethodA = CallStaticIntMethodA;
    myEnv->CallStaticLongMethod = CallStaticLongMethod;
    myEnv->CallStaticLongMethodV = CallStaticLongMethodV;
    myEnv->CallStaticLongMethodA = CallStaticLongMethodA;
    myEnv->CallStaticFloatMethod = CallStaticFloatMethod;
    myEnv->CallStaticFloatMethodV = CallStaticFloatMethodV;
    myEnv->CallStaticFloatMethodA = CallStaticFloatMethodA;
    myEnv->CallStaticDoubleMethod = CallStaticDoubleMethod;
    myEnv->CallStaticDoubleMethodV = CallStaticDoubleMethodV;
    myEnv->CallStaticDoubleMethodA = CallStaticDoubleMethodA;
    myEnv->CallStaticVoidMethod = CallStaticVoidMethod;
    myEnv->CallStaticVoidMethodV = CallStaticVoidMethodV;
    myEnv->CallStaticVoidMethodA = CallStaticVoidMethodA;
    myEnv->GetStaticFieldID = GetStaticFieldID;
    myEnv->GetStaticObjectField = GetStaticObjectField;
    myEnv->GetStaticBooleanField = GetStaticBooleanField;
    myEnv->GetStaticByteField = GetStaticByteField;
    myEnv->GetStaticCharField = GetStaticCharField;
    myEnv->GetStaticShortField = GetStaticShortField;
    myEnv->GetStaticIntField = GetStaticIntField;
    myEnv->GetStaticLongField = GetStaticLongField;
    myEnv->GetStaticFloatField = GetStaticFloatField;
    myEnv->GetStaticDoubleField = GetStaticDoubleField;
    myEnv->SetStaticObjectField = SetStaticObjectField;
    myEnv->SetStaticBooleanField = SetStaticBooleanField;
    myEnv->SetStaticByteField = SetStaticByteField;
    myEnv->SetStaticCharField = SetStaticCharField;
    myEnv->SetStaticShortField = SetStaticShortField;
    myEnv->SetStaticIntField = SetStaticIntField;
    myEnv->SetStaticLongField = SetStaticLongField;
    myEnv->SetStaticFloatField = SetStaticFloatField;
    myEnv->SetStaticDoubleField = SetStaticDoubleField;
    myEnv->NewString = NewString;
    myEnv->GetStringLength = GetStringLength;
    myEnv->GetStringChars = GetStringChars;
    myEnv->ReleaseStringChars = ReleaseStringChars;
    myEnv->NewStringUTF = NewStringUTF;
    myEnv->GetStringUTFLength = GetStringUTFLength;
    myEnv->GetStringUTFChars = GetStringUTFChars;
    myEnv->ReleaseStringUTFChars = ReleaseStringUTFChars;
    myEnv->GetArrayLength = GetArrayLength;
    myEnv->NewObjectArray = NewObjectArray;
    myEnv->GetObjectArrayElement = GetObjectArrayElement;
    myEnv->SetObjectArrayElement = SetObjectArrayElement;
    myEnv->NewBooleanArray = NewBooleanArray;
    myEnv->NewByteArray = NewByteArray;
    myEnv->NewCharArray = NewCharArray;
    myEnv->NewShortArray = NewShortArray;
    myEnv->NewIntArray = NewIntArray;
    myEnv->NewLongArray = NewLongArray;
    myEnv->NewFloatArray = NewFloatArray;
    myEnv->NewDoubleArray = NewDoubleArray;
    myEnv->GetBooleanArrayElements = GetBooleanArrayElements;
    myEnv->GetByteArrayElements = GetByteArrayElements;
    myEnv->GetCharArrayElements = GetCharArrayElements;
    myEnv->GetShortArrayElements = GetShortArrayElements;
    myEnv->GetIntArrayElements = GetIntArrayElements;
    myEnv->GetLongArrayElements = GetLongArrayElements;
    myEnv->GetFloatArrayElements = GetFloatArrayElements;
    myEnv->GetDoubleArrayElements = GetDoubleArrayElements;
    myEnv->ReleaseBooleanArrayElements = ReleaseBooleanArrayElements;
    myEnv->ReleaseByteArrayElements = ReleaseByteArrayElements;
    myEnv->ReleaseCharArrayElements = ReleaseCharArrayElements;
    myEnv->ReleaseShortArrayElements = ReleaseShortArrayElements;
    myEnv->ReleaseIntArrayElements = ReleaseIntArrayElements;
    myEnv->ReleaseLongArrayElements = ReleaseLongArrayElements;
    myEnv->ReleaseFloatArrayElements = ReleaseFloatArrayElements;
    myEnv->ReleaseDoubleArrayElements = ReleaseDoubleArrayElements;
    myEnv->GetBooleanArrayRegion = GetBooleanArrayRegion;
    myEnv->GetByteArrayRegion = GetByteArrayRegion;
    myEnv->GetCharArrayRegion = GetCharArrayRegion;
    myEnv->GetShortArrayRegion = GetShortArrayRegion;
    myEnv->GetIntArrayRegion = GetIntArrayRegion;
    myEnv->GetLongArrayRegion = GetLongArrayRegion;
    myEnv->GetFloatArrayRegion = GetFloatArrayRegion;
    myEnv->GetDoubleArrayRegion = GetDoubleArrayRegion;
    myEnv->SetBooleanArrayRegion = SetBooleanArrayRegion;
    myEnv->SetByteArrayRegion = SetByteArrayRegion;
    myEnv->SetCharArrayRegion = SetCharArrayRegion;
    myEnv->SetShortArrayRegion = SetShortArrayRegion;
    myEnv->SetIntArrayRegion = SetIntArrayRegion;
    myEnv->SetLongArrayRegion = SetLongArrayRegion;
    myEnv->SetFloatArrayRegion = SetFloatArrayRegion;
    myEnv->SetDoubleArrayRegion = SetDoubleArrayRegion;
    myEnv->RegisterNatives = RegisterNatives;
    myEnv->UnregisterNatives = UnregisterNatives;
    myEnv->MonitorEnter = MonitorEnter;
    myEnv->MonitorExit = MonitorExit;
    myEnv->GetJavaVM = GetJavaVM;
    myEnv->GetStringRegion = GetStringRegion;
    myEnv->GetStringUTFRegion = GetStringUTFRegion;
    myEnv->GetPrimitiveArrayCritical = GetPrimitiveArrayCritical;
    myEnv->ReleasePrimitiveArrayCritical = ReleasePrimitiveArrayCritical;
    myEnv->GetStringCritical = GetStringCritical;
    myEnv->ReleaseStringCritical = ReleaseStringCritical;
    myEnv->NewWeakGlobalRef = NewWeakGlobalRef;
    myEnv->DeleteWeakGlobalRef = DeleteWeakGlobalRef;
    myEnv->ExceptionCheck = ExceptionCheck;
    myEnv->NewDirectByteBuffer = NewDirectByteBuffer;
    myEnv->GetDirectBufferAddress = GetDirectBufferAddress;
    myEnv->GetDirectBufferCapacity = GetDirectBufferCapacity;
    myEnv->GetObjectRefType = GetObjectRefType;
    return myEnv;
}
