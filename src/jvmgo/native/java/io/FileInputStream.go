package io

import (
	. "jvmgo/any"
	"jvmgo/jvm/rtda"
	rtc "jvmgo/jvm/rtda/class"
	"jvmgo/util"
	"os"
)

func init() {
	_fis(fis_initIDs, "initIDs", "()V")
	_fis(close0, "close0", "()V")
	_fis(readBytes, "readBytes", "([BII)I")
	_fis(open, "open", "(Ljava/lang/String;)V")
}

func _fis(method Any, name, desc string) {
	rtc.RegisterNativeMethod("java/io/FileInputStream", name, desc, method)
}

// private static native void initIDs();
// ()V
func fis_initIDs(frame *rtda.Frame) {
	// todo
}

// private native void close0() throws IOException;
// ()V
func close0(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()

	goFile := this.Extra().(*os.File)
	err := goFile.Close()
	if err != nil {
		// todo
		panic("IOException")
	}
}

// private native void open(String name) throws FileNotFoundException;
// (Ljava/lang/String;)V
func open(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()
	name := vars.GetRef(1)

	goName := rtda.GoString(name)
	goFile, err := os.Open(goName)
	if err != nil {
		// todo
		panic(err.Error())
	}

	this.SetExtra(goFile)
}

// private native int readBytes(byte b[], int off, int len) throws IOException;
// ([BII)I
func readBytes(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()
	buf := vars.GetRef(1)
	off := vars.GetInt(2)
	_len := vars.GetInt(3)

	goFile := this.Extra().(*os.File)
	goBuf := util.CastInt8sToUint8s(buf.Fields().([]int8))
	goBuf = goBuf[:_len] // limit the maximum number of bytes read

	// func (f *File) ReadAt(b []byte, off int64) (n int, err error)
	n, err := goFile.ReadAt(goBuf, int64(off))
	if err != nil {
		frame.OperandStack().PushInt(int32(n))
	} else {
		// todo
		panic("IOException")
	}
}
