package options

import (
	"os"
	"path/filepath"

	"github.com/zxh0/jvm.go/jvmgo/cmdline"
)

var (
	VerboseClass    bool
	ThreadStackSize uint
	AbsJavaHome     string // /path/to/jre
	AbsJreLib       string // /path/to/jre/lib
)

func InitOptions(cmdOptions *cmdline.Options) {
	VerboseClass = cmdOptions.VerboseClass()
	ThreadStackSize = uint(cmdOptions.Xss)
	initJavaHome(cmdOptions.XuseJavaHome)
}

func initJavaHome(useOsEnv bool) {
	jh := "./jre"
	if useOsEnv {
		jh = os.Getenv("JAVA_HOME")
		if jh == "" {
			panic("$JAVA_HOME not set!")
		}
	}

	if absJh, err := filepath.Abs(jh); err == nil {
		AbsJavaHome = absJh
		AbsJreLib = filepath.Join(absJh, "lib")
	} else {
		panic(err)
	}
}
