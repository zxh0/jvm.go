package io

import (
    . "jvmgo/any"
    "jvmgo/jvm/rtda"
    rtc "jvmgo/jvm/rtda/class"
)

func init() {
    _fos(fos_initIDs,   "initIDs",      "()V")
    _fos(writeBytes,    "writeBytes",   "([BIIZ)V")
}

func _fos(method Any, name, desc string) {
    rtc.RegisterNativeMethod("java/io/FileOutputStream", name, desc, method)
}

// private static native void initIDs();
// ()V
func fos_initIDs(frame *rtda.Frame) {
    // todo
}

// private native void writeBytes(byte b[], int off, int len, boolean append) throws IOException;
// ([BIIZ)V
func writeBytes(frame *rtda.Frame) {
    stack := frame.OperandStack()
    stack.PopBoolean() // append
    stack.PopInt() // len
    stack.PopInt() // off
    stack.PopRef() // b
    stack.PopRef() // this
    println("writeBytes!!!")
}
