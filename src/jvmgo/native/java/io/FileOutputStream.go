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
    vars := frame.LocalVars()
    fosObj := vars.GetRef(0) // this
    byteArrObj := vars.GetRef(1) // b
    offset := vars.GetInt(2) // off
    length := vars.GetInt(3) // len
    //vars.GetBoolean(4) // append

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
    jBytes = jBytes[offset: offset+length]
    goBytes := util.CastInt8sToUint8s(jBytes)
    goFile.Write(goBytes)
}
