package lang

import (
	"runtime"
	"time"
	"unsafe"

	"github.com/zxh0/jvm.go/jvmgo/options"
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	"github.com/zxh0/jvm.go/jvmgo/rtda/heap"
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

func _system(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("java/lang/System", name, desc, method)
}

// public static native void arraycopy(Object src, int srcPos, Object dest, int destPos, int length)
// (Ljava/lang/Object;ILjava/lang/Object;II)V
func arraycopy(frame *rtda.Frame) {
	vars := frame.LocalVars()
	src := vars.GetRef(0)
	srcPos := vars.GetInt(1)
	dest := vars.GetRef(2)
	destPos := vars.GetInt(3)
	length := vars.GetInt(4)

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
		srcPos+length > heap.ArrayLength(src) ||
		destPos+length > heap.ArrayLength(dest) {

		panic("IndexOutOfBoundsException") // todo
	}

	heap.ArrayCopy(src, dest, srcPos, destPos, length)
}

func checkArrayCopy(src, dest *heap.Object) bool {
	srcClass := src.Class()
	destClass := dest.Class()

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
	stack := frame.OperandStack()
	stack.PushLong(millis)
}

// public static native int identityHashCode(Object x);
// (Ljava/lang/Object;)I
func identityHashCode(frame *rtda.Frame) {
	vars := frame.LocalVars()
	ref := vars.GetRef(0)

	// todo
	hashCode := int32(uintptr(unsafe.Pointer(ref)))
	stack := frame.OperandStack()
	stack.PushInt(hashCode)
}

// private static native Properties initProperties(Properties props);
// (Ljava/util/Properties;)Ljava/util/Properties;
func initProperties(frame *rtda.Frame) {
	vars := frame.LocalVars()
	props := vars.GetRef(0)

	stack := frame.OperandStack()
	stack.PushRef(props)

	// public synchronized Object setProperty(String key, String value)
	setPropMethod := props.Class().GetInstanceMethod("setProperty", "(Ljava/lang/String;Ljava/lang/String;)Ljava/lang/Object;")
	thread := frame.Thread()
	for key, val := range _sysProps() {
		jKey := rtda.JString(key)
		jVal := rtda.JString(val)
		args := []interface{}{props, jKey, jVal}
		thread.InvokeMethodWithShim(setPropMethod, args)
	}
}

func _sysProps() map[string]string {
	return map[string]string{
		"java.version":         "1.8.0",
		"java.vendor":          "jvm.go",
		"java.vendor.url":      "https://github.com/zxh0/jvm.go",
		"java.home":            options.AbsJavaHome,
		"java.class.version":   "52.0",
		"java.class.path":      heap.BootLoader().ClassPath().String(),
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
	stack := frame.OperandStack()
	stack.PushLong(nanoTime)
}

// private static native void setErr0(PrintStream err);
// (Ljava/io/PrintStream;)V
func setErr0(frame *rtda.Frame) {
	vars := frame.LocalVars()
	err := vars.GetRef(0)

	sysClass := frame.Method().Class()
	sysClass.SetStaticValue("err", "Ljava/io/PrintStream;", err)
}

// private static native void setIn0(InputStream in);
// (Ljava/io/InputStream;)V
func setIn0(frame *rtda.Frame) {
	vars := frame.LocalVars()
	in := vars.GetRef(0)

	sysClass := frame.Method().Class()
	sysClass.SetStaticValue("in", "Ljava/io/InputStream;", in)
}

// private static native void setOut0(PrintStream out);
// (Ljava/io/PrintStream;)V
func setOut0(frame *rtda.Frame) {
	vars := frame.LocalVars()
	out := vars.GetRef(0)

	sysClass := frame.Method().Class()
	sysClass.SetStaticValue("out", "Ljava/io/PrintStream;", out)
}
