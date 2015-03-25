package io

import (
	. "github.com/zxh0/jvm.go/jvmgo/any"
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
	"io"
	"os"
)

func init() {
	_fis(fis_initIDs, "initIDs", "()V")
	_fis(available, "available", "()I")
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

// public native int available() throws IOException;
// ()I
func available(frame *rtda.Frame) {
	// todo
	stack := frame.OperandStack()
	stack.PushInt(1)
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
		frame.Thread().ThrowFileNotFoundException(goName)
		return
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
	goBuf := buf.GoBytes()
	goBuf = goBuf[off : off+_len]

	// func (f *File) Read(b []byte) (n int, err error)
	n, err := goFile.Read(goBuf)
	if err == nil || n > 0 || err == io.EOF {
		frame.OperandStack().PushInt(int32(n))
	} else {
		// todo
		panic("IOException!" + err.Error())
	}
}
