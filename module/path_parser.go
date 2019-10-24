package module

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/zxh0/jvm.go/vm"
	"github.com/zxh0/jvm.go/vmutils"
)

// http://openjdk.java.net/jeps/261#Module-paths
func ParseModulePath(options *vm.Options) Path {
	sysPath := parsePath(filepath.Join(options.AbsJavaHome, "jmods")) // system module path
	if options.ModulePath == "" {
		return sysPath
	}

	userPath := parsePath(options.ModulePath)
	return append(sysPath, userPath...)
}

func parsePath(path string) []Module {
	list := make([]Module, 0, 4)
	for _, elem := range strings.Split(path, string(os.PathListSeparator)) {
		list = append(list, parseModules(elem, true)...)
	}
	return list
}

func parseModules(path string, goIntoDir bool) []Module {
	switch {
	case vmutils.IsJModFile(path):
		return []Module{NewJModModule(path)}
	case vmutils.IsJarFile(path):
		return []Module{newJarModule(path)}
	case isExplodedModule(path):
		return []Module{NewExplodedModule(path)}
	case goIntoDir:
		return listDir(path)
	default:
		return []Module{}
	}
}

func isExplodedModule(dir string) bool {
	return vmutils.IsExists(filepath.Join(dir, "module-info.class"))
}

func listDir(dir string) []Module {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err) // TODO
	}

	list := make([]Module, 0, 4)
	for _, file := range files {
		filename := filepath.Join(dir, file.Name())
		list = append(list, parseModules(filename, false)...)
	}

	Path(list).Sort()
	return list
}

func newJarModule(filename string) Module {
	zipFile, err := vmutils.OpenZipFile(filename)
	if err != nil {
		panic(err) // TODO
	}
	defer zipFile.Close()

	if zipFile.HasFile("module-info.class") {
		return NewModularJAR(filename)
	} else {
		return NewAutomaticModule(filename)
	}
}
