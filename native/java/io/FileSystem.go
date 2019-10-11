package io

import (
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
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
	thread := frame.Thread
	unixFsClass := frame.GetClassLoader().LoadClass("java/io/UnixFileSystem")
	if unixFsClass.InitializationNotStarted() {
		frame.NextPC = thread.PC() // undo getFileSystem
		thread.InitClass(unixFsClass)
		return
	}

	unixFsObj := unixFsClass.NewObj()
	frame.PushRef(unixFsObj)

	// call <init>
	frame.PushRef(unixFsObj) // this
	constructor := unixFsClass.GetDefaultConstructor()
	thread.InvokeMethod(constructor)
}
