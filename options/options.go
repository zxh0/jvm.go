package options

import (
	"errors"
	"flag"
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
	Classpath       string
	HelpFlag        bool
	VerboseClass    bool
	VerboseInstr    bool
	VerboseJNI      bool
	VersionFlag     bool
	Xss             string
	XUseJavaHome    bool
	XDebugInstr     bool
	XCPUProfile     string
	AbsJavaHome     string // /path/to/jre
	AbsJreLib       string // /path/to/jre/lib
	ThreadStackSize int
}

func Parse() (Options, []string) {
	options := Options{}
	flag.StringVar(&options.Classpath, "classpath", "", "Specifies a list of directories, JAR files, and ZIP archives to search for class files.")
	flag.StringVar(&options.Classpath, "cp", "", "Specifies a list of directories, JAR files, and ZIP archives to search for class files.")
	flag.BoolVar(&options.HelpFlag, "help", false, "Displays usage information and exit.")
	flag.BoolVar(&options.HelpFlag, "h", false, "Displays usage information and exit.")
	flag.BoolVar(&options.HelpFlag, "?", false, "Displays usage information and exit.")
	flag.BoolVar(&options.VerboseClass, "verbose:class", false, "Displays information about each class loaded.")
	flag.BoolVar(&options.VerboseInstr, "verbose:instr", false, "Displays information about each instruction executed.")
	flag.BoolVar(&options.VerboseJNI, "verbose:jni", false, "Displays information about the use of native methods and other Java Native Interface (JNI) activity.")
	flag.BoolVar(&options.VersionFlag, "version", false, "Displays version information and exit.")
	flag.StringVar(&options.Xss, "Xss", "", "Sets the thread stack size.")
	flag.BoolVar(&options.XUseJavaHome, "XuseJavaHome", false, "Uses JAVA_HOME")
	flag.BoolVar(&options.XDebugInstr, "Xdebug:instr", false, "Displays executed instructions")
	flag.StringVar(&options.XCPUProfile, "Xprofile:cpu", "", "")
	flag.Parse()

	options.AbsJavaHome = getJavaHome(options.XUseJavaHome)
	options.AbsJreLib = filepath.Join(options.AbsJavaHome, "lib")
	options.ThreadStackSize = parseXss(options.Xss)
	return options, flag.Args()
}

func PrintDefaults() {
	flag.PrintDefaults()
}

func getJavaHome(useOsEnv bool) string {
	jh := "./jre"
	if useOsEnv {
		jh = os.Getenv("JAVA_HOME")
		if jh == "" {
			panic("$JAVA_HOME not set!")
		}
	}

	if absJh, err := filepath.Abs(jh); err == nil {
		if strings.Contains(absJh, "jre") {
			return absJh
		} else {
			return filepath.Join(absJh, "jre")
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
