package io

import (
    "os"
    . "jvmgo/any"
    "jvmgo/jvm/rtda"
    "jvmgo/util"
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
    byteArrObj := stack.PopRef() // b
    fosObj := stack.PopRef() // this

    fdObj := fosObj.GetFieldValue("fd", "Ljava/io/FileDescriptor;").(*rtc.Obj)
    if fdObj.Extra() == nil {
        goFd := fdObj.GetFieldValue("fd", "I").(int32)
        switch goFd {
        case 0: fdObj.SetExtra(os.Stdin)
        case 1: fdObj.SetExtra(os.Stdout)
        case 2: fdObj.SetExtra(os.Stderr)
        }
    }
    goFile := fdObj.Extra().(*os.File)

    jBytes := byteArrObj.Fields().([]int8)
    goBytes := util.Int8ToUint8(jBytes)
    goFile.Write(goBytes)
}
