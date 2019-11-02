package loader

import (
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

func init() {
	_bl(getSystemPackageNames, "getSystemPackageNames", "()[Ljava/lang/String;")
	_bl(getSystemPackageLocation, "getSystemPackageLocation", "(Ljava/lang/String;)Ljava/lang/String;")
	_bl(setBootLoaderUnnamedModule0, "setBootLoaderUnnamedModule0", "(Ljava/lang/Module;)V")
}

func _bl(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("jdk/internal/loader/BootLoader", name, desc, method)
}

/**
 * Returns an array of the binary name of the packages defined by
 * the boot loader, in VM internal form (forward slashes instead of dot).
 */
// private static native String[] getSystemPackageNames();
func getSystemPackageNames(frame *rtda.Frame) {
	panic("TODO")
}

/**
 * Returns the location of the package of the given name, if
 * defined by the boot loader; otherwise {@code null} is returned.
 *
 * The location may be a module from the runtime image or exploded image,
 * or from the boot class append path (i.e. -Xbootclasspath/a or
 * BOOT-CLASS-PATH attribute specified in java agent).
 */
// private static native String getSystemPackageLocation(String name);
func getSystemPackageLocation(frame *rtda.Frame) {
	jPkgName := frame.GetRefVar(0)
	goPkgName := jPkgName.JSToGoStr()

	location := frame.GetRuntime().GetSystemPackageLocation(goPkgName)

	// TODO
	panic("TODO:" + goPkgName + ",,," + frame.Thread.VMOptions.AbsJavaHome + ":::" + location)
}

// private static native void setBootLoaderUnnamedModule0(Module module);
func setBootLoaderUnnamedModule0(frame *rtda.Frame) {
	// TODO
}
