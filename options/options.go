package options

import (
	"os"
	"path/filepath"
	"strings"
)

var (
	VerboseClass    bool
	ThreadStackSize uint
	AbsJavaHome     string // /path/to/jre
	AbsJreLib       string // /path/to/jre/lib
)

func InitOptions(verboseClass bool, xss int, useJavaHome bool) {
	VerboseClass = verboseClass
	ThreadStackSize = uint(xss)
	initJavaHome(useJavaHome)
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
		if strings.Contains(absJh, "jre") {
			AbsJavaHome = absJh
			AbsJreLib = filepath.Join(absJh, "lib")
		} else {
			AbsJavaHome = filepath.Join(absJh, "jre")
			AbsJreLib = filepath.Join(absJh, "jre", "lib")
		}
	} else {
		panic(err)
	}
}
