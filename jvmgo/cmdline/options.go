package cmdline

import (
	"strconv"
	"strings"
)

const (
	_1k = 1024
	_1m = _1k * _1k
	_1g = _1k * _1m
)

type Options struct {
	classpath    string
	verboseClass bool
	Xss          int
	Xcpuprofile  string
	XuseJavaHome bool
}

// getters
func (self *Options) Classpath() string {
	return self.classpath
}
func (self *Options) VerboseClass() bool {
	return self.verboseClass
}

func parseOptions(argReader *ArgReader) *Options {
	options := &Options{
		Xss: 16 * _1k,
	}

	for argReader.hasMoreOptions() {
		optionName := argReader.removeFirst()
		switch optionName {
		case "-cp", "-classpath":
			options.classpath = argReader.removeFirst()
		case "-verbose", "-verbose:class":
			options.verboseClass = true
		case "-Xcpuprofile":
			options.Xcpuprofile = argReader.removeFirst()
		case "-XuseJavaHome":
			options.XuseJavaHome = true
		default:
			if strings.HasPrefix(optionName, "-Xss") {
				options.Xss = parseXss(optionName)
			} else {
				panic("Unrecognized option: " + optionName)
			}
		}
	}

	return options
}

// -Xss<size>[g|G|m|M|k|K]
func parseXss(optionName string) int {
	size := optionName[4:]
	switch size[len(size)-1] {
	case 'g', 'G':
		return _1g * parseInt(size[:len(size)-1])
	case 'm', 'M':
		return _1m * parseInt(size[:len(size)-1])
	case 'k', 'K':
		return _1k * parseInt(size[:len(size)-1])
	default:
		return parseInt(size)
	}
}

func parseInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err.Error())
	}
	return i
}
