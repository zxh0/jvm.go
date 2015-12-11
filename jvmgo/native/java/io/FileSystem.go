package io

import (
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	"github.com/zxh0/jvm.go/jvmgo/rtda/heap"
)

func init() {
	_fs(getFileSystem, "getFileSystem", "()Ljava/io/FileSystem;")
}

func _fs(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("java/io/FileSystem", name, desc, method)
}

// public static native FileSystem getFileSystem()
// ()Ljava/io/FileSystem;
func getFileSystem(frame *rtda.Frame) {
	thread := frame.Thread()
	unixFsClass := frame.ClassLoader().LoadClass("java/io/UnixFileSystem")
	if unixFsClass.InitializationNotStarted() {
		frame.SetNextPC(thread.PC()) // undo getFileSystem
		thread.InitClass(unixFsClass)
		return
	}

	stack := frame.OperandStack()
	unixFsObj := unixFsClass.NewObj()
	stack.PushRef(unixFsObj)

	// call <init>
	stack.PushRef(unixFsObj) // this
	constructor := unixFsClass.GetDefaultConstructor()
	thread.InvokeMethod(constructor)
}
