package main

import (
	"archive/zip"
	"fmt"
	"github.com/zxh0/jvm.go/jvmgo/classfile"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		jarFileName := os.Args[1]
		handleJar(jarFileName)
	}
}

func handleJar(jarFileName string) {
	//fmt.Printf("jar: %v\n", jarFileName)

	// open jar
	r, err := zip.OpenReader(jarFileName) // func OpenReader(name string) (*ReadCloser, error)
	if err != nil {
		panic(err.Error())
	}
	defer r.Close()

	// find classes
	for _, f := range r.File {
		if strings.HasSuffix(f.Name, ".class") {
			if !skip(f.Name) {
				handleClass(f)
			}
		}
	}
}

func skip(className string) bool {
	return strings.HasPrefix(className, "apple") ||
		strings.HasPrefix(className, "com/apple") ||
		strings.HasPrefix(className, "com/sun/java/swing") ||
		strings.HasPrefix(className, "com/sun/media/sound") ||
		strings.HasPrefix(className, "sun/font") ||
		strings.HasPrefix(className, "sun/java2d") ||
		strings.HasPrefix(className, "sun/lwawt/macosx")
}

func handleClass(f *zip.File) {
	//fmt.Printf("%v\n", f.Name)

	// open classfile
	rc, err := f.Open() // func (f *File) Open() (rc io.ReadCloser, err error)
	if err != nil {
		panic(err.Error())
	}

	// read class data
	data, err := ioutil.ReadAll(rc) // func ReadAll(r io.Reader) ([]byte, error)
	rc.Close()
	if err != nil {
		panic(err.Error())
	}

	// parse classfile
	cf, err := classfile.Parse(data)
	if err != nil {
		panic(err.Error())
	}

	handleClassfile(cf)
}

func handleClassfile(cf *classfile.ClassFile) {
	for _, m := range cf.Methods() {
		if isNative(m) {
			if isStatic(m) {
				fmt.Printf("%v.%v%v\n", cf.ClassName(), m.Name(), m.Descriptor())
			} else {
				fmt.Printf("%v#%v%v\n", cf.ClassName(), m.Name(), m.Descriptor())
			}
		}
	}
}

func isNative(m *classfile.MethodInfo) bool {
	return m.AccessFlags()&0x0100 != 0
}
func isStatic(m *classfile.MethodInfo) bool {
	return m.AccessFlags()&0x0008 != 0
}
