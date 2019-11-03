package zip

import (
	gozip "archive/zip"
	"fmt"

	"github.com/zxh0/jvm.go/native"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/vmutils"
)

const (
	JZENTRY_NAME    = 0
	JZENTRY_EXTRA   = 1
	JZENTRY_COMMENT = 2
)

func init() {
	_zf(zf_initIDs, "initIDs", "()V")
	_zf(freeEntry, "freeEntry", "(JJ)V")
	_zf(getEntry, "getEntry", "(J[BZ)J")
	_zf(getEntryBytes, "getEntryBytes", "(JI)[B")
	_zf(getEntryCrc, "getEntryCrc", "(J)J")
	_zf(getEntryCSize, "getEntryCSize", "(J)J")
	_zf(getEntryFlag, "getEntryFlag", "(J)I")
	_zf(getEntryMethod, "getEntryMethod", "(J)I")
	_zf(getEntrySize, "getEntrySize", "(J)J")
	_zf(getEntryTime, "getEntryTime", "(J)J")
	_zf(getNextEntry, "getNextEntry", "(JI)J")
	_zf(getTotal, "getTotal", "(J)I")
	_zf(open, "open", "(Ljava/lang/String;IJZ)J")
	_zf(read, "read", "(JJJ[BII)I")
	_zf(startsWithLOC, "startsWithLOC", "(J)Z")
}

func _zf(method native.Method, name, desc string) {
	native.Register("java/util/zip/ZipFile", name, desc, method)
}

// private static native void initIDs();
// ()V
func zf_initIDs(frame *rtda.Frame) {
	// todo
}

// private static native long open(String name, int mode, long lastModified,
//                                 boolean usemmap) throws IOException;
// (Ljava/lang/String;IJZ)J
func open(frame *rtda.Frame) {
	nameObj := frame.GetRefVar(0)

	name := nameObj.JSToGoStr()
	jzfile, err := openZip(name)
	if err != nil {
		// todo
		panic("IOException!" + err.Error())
	}

	frame.PushLong(jzfile)
}

// private static native boolean startsWithLOC(long jzfile);
// (J)Z
func startsWithLOC(frame *rtda.Frame) {
	// todo
	frame.PushBoolean(true)
}

// private static native int getTotal(long jzfile);
// (J)I
func getTotal(frame *rtda.Frame) {
	jzfile := frame.GetLongVar(0)

	total := getEntryCount(jzfile)

	frame.PushInt(total)
}

// private static native long getNextEntry(long jzfile, int i);
// (JI)J
func getNextEntry(frame *rtda.Frame) {
	jzfile := frame.GetLongVar(0)
	i := frame.GetIntVar(2)

	jzentry := getJzentry(jzfile, i)

	frame.PushLong(jzentry)
}

// private static native void freeEntry(long jzfile, long jzentry);
// (JJ)V
func freeEntry(frame *rtda.Frame) {
	// todo
}

// private static native long getEntry(long jzfile, byte[] name, boolean addSlash);
// (J[BZ)J
func getEntry(frame *rtda.Frame) {
	jzfile := frame.GetLongVar(0)
	nameObj := frame.GetRefVar(2)
	//addSlash := frame.GetBooleanVar(3)

	// todo
	name := nameObj.GetGoBytes()
	jzentry := getJzentry2(jzfile, name)

	frame.PushLong(jzentry)
}

// private static native byte[] getEntryBytes(long jzentry, int type);
// (JI)[B
func getEntryBytes(frame *rtda.Frame) {
	jzentry := frame.GetLongVar(0)
	_type := frame.GetIntVar(2)

	goBytes := _getEntryBytes(jzentry, _type)
	jBytes := vmutils.CastBytesToInt8s(goBytes)
	byteArr := frame.GetRuntime().NewByteArray(jBytes)

	frame.PushRef(byteArr)
}

func _getEntryBytes(jzentry int64, _type int32) []byte {
	entry := getEntryFile(jzentry)
	switch _type {
	case JZENTRY_NAME:
		return []byte(entry.Name)
	case JZENTRY_EXTRA:
		return entry.Extra
	case JZENTRY_COMMENT:
		return []byte(entry.Comment)
	}
	panic(fmt.Errorf("invalid type: %v", _type))
}

// private static native int getEntryFlag(long jzentry);
// (J)I
func getEntryFlag(frame *rtda.Frame) {
	entry := _getEntryPop(frame)
	flag := int32(entry.Flags)

	frame.PushInt(flag)
}

// private static native long getEntryTime(long jzentry);
// (J)J
func getEntryTime(frame *rtda.Frame) {
	entry := _getEntryPop(frame)
	modDate := entry.ModifiedDate
	modTime := entry.ModifiedTime
	time := int64(modDate)<<16 | int64(modTime)

	frame.PushLong(time)
}

// private static native long getEntryCrc(long jzentry);
// (J)J
func getEntryCrc(frame *rtda.Frame) {
	entry := _getEntryPop(frame)
	crc := int64(entry.CRC32)

	frame.PushLong(crc)
}

// private static native long getEntrySize(long jzentry);
// (J)J
func getEntrySize(frame *rtda.Frame) {
	entry := _getEntryPop(frame)
	size := int64(entry.UncompressedSize64)

	frame.PushLong(size)
}

// private static native long getEntryCSize(long jzentry);
// (J)J
func getEntryCSize(frame *rtda.Frame) {
	// entry := _getEntryPop(frame)
	// size := int64(entry.CompressedSize64)

	// stack := frame.
	// frame.PushLong(size)

	// todo
	getEntrySize(frame)
}

// private static native int getEntryMethod(long jzentry);
// (J)I
func getEntryMethod(frame *rtda.Frame) {
	// entry := _getEntryPop(frame)
	// method := int32(entry.Method)

	// todo
	frame.PushInt(0)
}

func _getEntryPop(frame *rtda.Frame) *gozip.File {
	jzentry := frame.GetLongVar(0)

	entry := getEntryFile(jzentry)
	return entry
}

// private static native int read(long jzfile, long jzentry,
//                                long pos, byte[] b, int off, int len);
// (JJJ[BII)I
func read(frame *rtda.Frame) {
	//jzfile := frame.GetLongVar(0)
	jzentry := frame.GetLongVar(2)
	pos := frame.GetLongVar(4)
	byteArr := frame.GetRefVar(6)
	off := frame.GetIntVar(7)
	_len := frame.GetIntVar(8)

	goBytes := byteArr.GetGoBytes()
	goBytes = goBytes[off : off+_len]
	n := readEntry(jzentry, pos, goBytes)

	frame.PushInt(int32(n))
}
