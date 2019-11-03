package io

import (
	"os"
	"path/filepath"

	"github.com/zxh0/jvm.go/native"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
	"github.com/zxh0/jvm.go/vmutils"
)

func init() {
	_ufs(canonicalize0, "canonicalize0", "(Ljava/lang/String;)Ljava/lang/String;")
	_ufs(getBooleanAttributes0, "getBooleanAttributes0", "(Ljava/io/File;)I")
	_ufs(getLastModifiedTime, "getLastModifiedTime", "(Ljava/io/File;)J")
}

func _ufs(method native.Method, name, desc string) {
	native.Register("java/io/UnixFileSystem", name, desc, method)
}

// private native String canonicalize0(String path) throws IOException;
// (Ljava/lang/String;)Ljava/lang/String;
func canonicalize0(frame *rtda.Frame) {
	pathStr := frame.GetRefVar(1)

	// todo
	path := pathStr.JSToGoStr()
	path2 := filepath.Clean(path)
	if path2 != path {
		pathStr = frame.GetRuntime().JSFromGoStr(path2)
	}

	frame.PushRef(pathStr)
}

// public native int getBooleanAttributes0(File f);
// (Ljava/io/File;)I
func getBooleanAttributes0(frame *rtda.Frame) {
	fileObj := frame.GetRefVar(1)
	path := _getPath(fileObj)

	// todo
	attributes0 := 0
	if vmutils.IsExists(path) {
		attributes0 |= 0x01
	}
	if vmutils.IsDir(path) {
		attributes0 |= 0x04
	}

	frame.PushInt(int32(attributes0))
}

func _getPath(fileObj *heap.Object) string {
	pathStr := fileObj.GetFieldValue("path", "Ljava/lang/String;").Ref
	return pathStr.JSToGoStr()
}

// public native long getLastModifiedTime(File f);
//
func getLastModifiedTime(frame *rtda.Frame) {
	fileObj := frame.GetRefVar(1)

	path := _getPath(fileObj)
	t := _getLastModifiedTime(path)

	frame.PushLong(t)
}

func _getLastModifiedTime(path string) int64 {
	fileInfo, err := os.Stat(path)
	if err == nil {
		return fileInfo.ModTime().UnixNano() / 1000
	}
	return 0
}
