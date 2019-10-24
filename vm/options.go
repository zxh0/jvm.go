package vm

import (
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	_1k = 1024
	_1m = _1k * _1k
	_1g = _1k * _1m
)

type Options struct {
	MainModule      string
	MainClass       string
	ClassPath       string
	ModulePath      string
	VerboseClass    bool
	VerboseModule   bool
	VerboseJNI      bool
	Xss             string
	Xjre            string
	XUseJavaHome    bool
	XDebugInstr     bool
	XCPUProfile     string
	AbsJavaHome     string // /path/to/jre
	AbsJreLib       string // /path/to/jre/lib
	ThreadStackSize int
}

func (options *Options) Init() {
	if options.ModulePath != "" {
		options.AbsJavaHome = getJavaHome13(options.Xjre)
	} else {
		options.AbsJavaHome = getJavaHome8(options.Xjre, options.XUseJavaHome)
		options.AbsJreLib = filepath.Join(options.AbsJavaHome, "lib")
	}
	options.ThreadStackSize = parseXss(options.Xss)
}

func getJavaHome13(jreDir string) string {
	if absJH, err := filepath.Abs(jreDir); err != nil {
		panic(err) // TODO
	} else {
		return absJH
	}
}

func getJavaHome8(jreDir string, useOsEnv bool) string {
	jh := "./jre"
	if jreDir != "" {
		jh = jreDir
	} else if useOsEnv {
		if jh = os.Getenv("JAVA_HOME"); jh == "" {
			panic("$JAVA_HOME not set!")
		}
	}

	if absJH, err := filepath.Abs(jh); err == nil {
		if strings.Contains(absJH, "jre") {
			return absJH
		} else {
			return filepath.Join(absJH, "jre")
		}
	} else {
		panic(err) // TODO
	}
}

// -Xss<size>[g|G|m|M|k|K]
func parseXss(size string) int {
	if size == "" {
		return 16 * _1k
	}
	switch size[len(size)-1] {
	case 'g', 'G':
		return parseSS(size[:len(size)-1], _1g)
	case 'm', 'M':
		return parseSS(size[:len(size)-1], _1m)
	case 'k', 'K':
		return parseSS(size[:len(size)-1], _1k)
	default:
		return parseSS(size, 1)
	}
}

func parseSS(size string, unit int) int {
	if i, err := strconv.Atoi(size); err != nil {
		panic(errors.New("invalid thread stack size: " + size))
	} else {
		return i * unit
	}
}
