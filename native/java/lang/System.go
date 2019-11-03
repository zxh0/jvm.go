package lang

import (
	"runtime"
	"sort"
	"time"
	"unsafe"

	"github.com/zxh0/jvm.go/native"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
	"github.com/zxh0/jvm.go/vm"
)

func init() {
	_system(arraycopy, "arraycopy", "(Ljava/lang/Object;ILjava/lang/Object;II)V")
	_system(currentTimeMillis, "currentTimeMillis", "()J")
	_system(identityHashCode, "identityHashCode", "(Ljava/lang/Object;)I")
	_system(initProperties, "initProperties", "(Ljava/util/Properties;)Ljava/util/Properties;")
	_system(nanoTime, "nanoTime", "()J")
	_system(setErr0, "setErr0", "(Ljava/io/PrintStream;)V")
	_system(setIn0, "setIn0", "(Ljava/io/InputStream;)V")
	_system(setOut0, "setOut0", "(Ljava/io/PrintStream;)V")
}

func _system(method native.Method, name, desc string) {
	native.Register("java/lang/System", name, desc, method)
}

// public static native void arraycopy(Object src, int srcPos, Object dest, int destPos, int length)
// (Ljava/lang/Object;ILjava/lang/Object;II)V
func arraycopy(frame *rtda.Frame) {
	src := frame.GetRefVar(0)
	srcPos := frame.GetIntVar(1)
	dest := frame.GetRefVar(2)
	destPos := frame.GetIntVar(3)
	length := frame.GetIntVar(4)

	// NullPointerException
	if src == nil || dest == nil {
		panic("NPE") // todo
	}
	// ArrayStoreException
	if !checkArrayCopy(src, dest) {
		panic("ArrayStoreException")
	}
	// IndexOutOfBoundsException
	if srcPos < 0 || destPos < 0 || length < 0 ||
		srcPos+length > src.ArrayLength() ||
		destPos+length > dest.ArrayLength() {

		panic("IndexOutOfBoundsException") // todo
	}

	heap.ArrayCopy(src, dest, srcPos, destPos, length)
}

func checkArrayCopy(src, dest *heap.Object) bool {
	srcClass := src.Class
	destClass := dest.Class

	if !srcClass.IsArray() || !destClass.IsArray() {
		return false
	}
	if srcClass.IsPrimitiveArray() || destClass.IsPrimitiveArray() {
		return srcClass == destClass
	}
	return true
}

// public static native long currentTimeMillis();
// ()J
func currentTimeMillis(frame *rtda.Frame) {
	millis := time.Now().UnixNano() / int64(time.Millisecond)
	frame.PushLong(millis)
}

// public static native int identityHashCode(Object x);
// (Ljava/lang/Object;)I
func identityHashCode(frame *rtda.Frame) {
	ref := frame.GetRefVar(0)

	// todo
	hashCode := int32(uintptr(unsafe.Pointer(ref)))
	frame.PushInt(hashCode)
}

// private static native Properties initProperties(Properties props);
// (Ljava/util/Properties;)Ljava/util/Properties;
func initProperties(frame *rtda.Frame) {
	props := frame.GetRefVar(0)

	frame.PushRef(props)

	// public synchronized Object setProperty(String key, String value)
	setPropMethod := props.Class.GetInstanceMethod("setProperty", "(Ljava/lang/String;Ljava/lang/String;)Ljava/lang/Object;")
	thread := frame.Thread

	sysPropMap := _getSysProps(thread.VMOptions)
	sysPropKeys := _getSysPropKeys(sysPropMap)

	for _, key := range sysPropKeys {
		val := sysPropMap[key]
		jKey := frame.GetRuntime().JSFromGoStr(key)
		jVal := frame.GetRuntime().JSFromGoStr(val)
		args := []heap.Slot{heap.NewRefSlot(props), heap.NewRefSlot(jKey), heap.NewRefSlot(jVal)}
		thread.InvokeMethodWithShim(setPropMethod, args)
	}
}

func _getSysPropKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for key, _ := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func _getSysProps(opts *vm.Options) map[string]string {
	return map[string]string{
		"java.version":         "1.8.0",
		"java.vendor":          "jvm.go",
		"java.vendor.url":      "https://github.com/zxh0/jvm.go",
		"java.home":            opts.AbsJavaHome,
		"java.class.version":   "52.0",
		"java.class.path":      opts.ClassPath, // TODO
		"java.awt.graphicsenv": "sun.awt.CGraphicsEnvironment",
		"os.name":              runtime.GOOS,   // todo
		"os.arch":              runtime.GOARCH, // todo
		"os.version":           "",             // todo
		"file.separator":       "/",            // todo os.PathSeparator
		"path.separator":       ":",            // todo os.PathListSeparator
		"line.separator":       "\n",           // todo
		"user.name":            "",             // todo
		"user.home":            "",             // todo
		"user.dir":             ".",            // todo
		"user.country":         "CN",           // todo
		"file.encoding":        "UTF-8",
		"sun.stdout.encoding":  "UTF-8",
		"sun.stderr.encoding":  "UTF-8",
	}
}

// public static native long nanoTime();
// ()J
func nanoTime(frame *rtda.Frame) {
	nanoTime := time.Now().UnixNano()
	frame.PushLong(nanoTime)
}

// private static native void setErr0(PrintStream err);
// (Ljava/io/PrintStream;)V
func setErr0(frame *rtda.Frame) {
	err := frame.GetRefVar(0) // TODO

	sysClass := frame.GetClass()
	sysClass.SetStaticValue("err", "Ljava/io/PrintStream;", heap.NewRefSlot(err))
}

// private static native void setIn0(InputStream in);
// (Ljava/io/InputStream;)V
func setIn0(frame *rtda.Frame) {
	in := frame.GetRefVar(0) // TODO

	sysClass := frame.GetClass()
	sysClass.SetStaticValue("in", "Ljava/io/InputStream;", heap.NewRefSlot(in))
}

// private static native void setOut0(PrintStream out);
// (Ljava/io/PrintStream;)V
func setOut0(frame *rtda.Frame) {
	out := frame.GetRefVar(0) // TODO

	sysClass := frame.GetClass()
	sysClass.SetStaticValue("out", "Ljava/io/PrintStream;", heap.NewRefSlot(out))
}
