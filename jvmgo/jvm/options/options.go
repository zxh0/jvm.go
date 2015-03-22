package options

import (
	"github.com/zxh0/jvm.go/jvmgo/cmdline"
	"path/filepath"
)

// todo
const JavaHome = "./jre/"

var (
	AbsJavaHome     string
	VerboseClass    bool
	ThreadStackSize uint
)

func init() {
	if absJavaHome, err := filepath.Abs(JavaHome); err == nil {
		AbsJavaHome = absJavaHome
	} else {
		panic(err)
	}
}

func InitOptions(cmdOptions *cmdline.Options) {
	VerboseClass = cmdOptions.VerboseClass()
	ThreadStackSize = uint(cmdOptions.Xss())
}
