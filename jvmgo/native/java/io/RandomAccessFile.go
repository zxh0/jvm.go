package io

import (
	"encoding/binary"
	. "github.com/zxh0/jvm.go/jvmgo/any"
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
	"github.com/zxh0/jvm.go/jvmgo/util"
	"os"
	"syscall"
)

func init() {
	_raf(raf_initIDs, "initIDs", "()V")
	_raf(raf_open, "open", "(Ljava/lang/String;I)V")
	_raf(raf_close0, "close0", "()V")
	_raf(raf_write0, "write0", "(I)V")
	_raf(raf_writeBytes, "writeBytes", "([BII)V")
}

func _raf(method Any, name, desc string) {
	rtc.RegisterNativeMethod("java/io/RandomAccessFile", name, desc, method)
}

// private static native void initIDs();
// ()V
func raf_initIDs(frame *rtda.Frame) {
	// todo
}

// private native void open(String name, int mode) throws FileNotFoundException;
// (Ljava/lang/String;)V
func raf_open(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()
	name := vars.GetRef(1)
	mode := vars.GetInt(2) //flag
	flag := 0

	if mode&1 > 0 {
		flag |= syscall.O_RDONLY
	}

	//write
	if mode&2 > 0 {
		flag |= syscall.O_RDWR | syscall.O_CREAT
	}

	if mode&4 > 0 {
		flag |= syscall.O_SYNC | syscall.O_CREAT
	}

	if mode&8 > 0 {
		flag |= syscall.O_DSYNC | syscall.O_CREAT
	}

	goName := rtda.GoString(name)
	goFile, err := os.OpenFile(goName, flag, 0660)
	if err != nil {
		frame.Thread().ThrowFileNotFoundException(goName)
		return
	}

	this.SetExtra(goFile)
}

// private native void close0() throws IOException;
// ()V
func raf_close0(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()

	goFile := this.Extra().(*os.File)
	err := goFile.Close()
	if err != nil {
		// todo
		panic("IOException")
	}
}

// private native void writeBytes(byte b[], int off, int len) throws IOException;
// ([BIIZ)V
func raf_writeBytes(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()       // this
	byteArrObj := vars.GetRef(1) // b
	offset := vars.GetInt(2)     // off
	length := vars.GetInt(3)     // len

	goFile := this.Extra().(*os.File)

	jBytes := byteArrObj.Fields().([]int8)
	jBytes = jBytes[offset : offset+length]
	goBytes := util.CastInt8sToUint8s(jBytes)
	goFile.Write(goBytes)
}

// private native void write0(int b) throws IOException;
// (I)V
func raf_write0(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()
	intObj := vars.GetInt(1) // b

	goFile := this.Extra().(*os.File)
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, uint16(intObj))
	_, err := goFile.Write(b)

	if err != nil {
		panic("IOException!" + err.Error())
	}
}
