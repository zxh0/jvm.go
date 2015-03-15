package options

import (
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
