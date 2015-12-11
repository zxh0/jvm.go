package io

import (
	"io"
	"os"

	"github.com/zxh0/jvm.go/jvmgo/rtda"
	"github.com/zxh0/jvm.go/jvmgo/rtda/heap"
)

func init() {
	_fis(available, "available", "()I")
	_fis(close0, "close0", "()V")
	_fis(readBytes, "readBytes", "([BII)I")
	_fis(open, "open0", "(Ljava/lang/String;)V")
}

func _fis(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("java/io/FileInputStream", name, desc, method)
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
