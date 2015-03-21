package options

import (
	"os"
	"path/filepath"
)

// todo
const JavaHome = "."

var (
	AbsJavaHome     string
	VerboseClass    bool
	ThreadStackSize uint
)

func init() {
	jh := os.Getenv("JAVA_HOME")
	if jh == "" {
		jh = JavaHome
	}

	if absJavaHome, err := filepath.Abs(jh); err == nil {
		AbsJavaHome = absJavaHome
	} else {
		panic(err)
	}
}
