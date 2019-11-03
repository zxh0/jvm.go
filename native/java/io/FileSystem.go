package io

import (
	"github.com/zxh0/jvm.go/native"
	"github.com/zxh0/jvm.go/rtda"
)

func init() {
	_fs(getFileSystem, "getFileSystem", "()Ljava/io/FileSystem;")
}

func _fs(method native.Method, name, desc string) {
	native.Register("java/io/FileSystem", name, desc, method)
}

// public static native FileSystem getFileSystem()
// ()Ljava/io/FileSystem;
func getFileSystem(frame *rtda.Frame) {
	thread := frame.Thread
	unixFsClass := frame.GetClassLoader().LoadClass("java/io/UnixFileSystem")
	if unixFsClass.InitializationNotStarted() {
		frame.NextPC = thread.PC // undo getFileSystem
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
