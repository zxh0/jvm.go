package main

import (
	"errors"
	"strconv"
	"strings"
)

const (
	_1k = 1024
	_1m = _1k * _1k
	_1g = _1k * _1m
)

type Options struct {
	Classpath    string
	VerboseClass bool
	Xss          int
	Xcpuprofile  string
	XuseJavaHome bool
}

// java [ options ] class [ arguments ]
type Command struct {
	Options Options
	Class   string
	Args    []string
}

func parseCommand(osArgs []string) (Command, error) {
	args := osArgs[1:]
	options, err := parseOptions(&args)
	if err != nil {
		return Command{}, err
	}

	class := removeFirst(&args)
	cmd := Command{
		Options: options,
		Class:   class,
		Args:    args,
	}
	return cmd, nil
}

func parseOptions(args *[]string) (Options, error) {
	options := Options{
		Xss: 16 * _1k,
	}

	for hasMoreOptions(*args) {
		optionName := removeFirst(args)
		switch optionName {
		case "-cp", "-classpath":
			options.Classpath = removeFirst(args)
		case "-verbose", "-verbose:class":
			options.VerboseClass = true
		case "-Xcpuprofile":
			options.Xcpuprofile = removeFirst(args)
		case "-XuseJavaHome":
			options.XuseJavaHome = true
		default:
			if !strings.HasPrefix(optionName, "-Xss") {
				return options, errors.New("Unrecognized option: " + optionName)
			}
			ss, err := parseXss(optionName)
			if err != nil {
				return options, err
			}
			options.Xss = ss
		}
	}

	return options, nil
}

func hasMoreOptions(args []string) bool {
	return len(args) > 0 && args[0][0] == '-'
}

func removeFirst(args *[]string) string {
	first := (*args)[0]
	*args = (*args)[1:]
	return first
}

// -Xss<size>[g|G|m|M|k|K]
func parseXss(optionName string) (int, error) {
	size := optionName[4:] // remove -Xss
	if len(size) > 0 {
		switch size[len(size)-1] {
		case 'g', 'G':
			return parseSS(size[:len(size)-1], _1g)
		case 'm', 'M':
			return parseSS(size[:len(size)-1], _1m)
		case 'k', 'K':
			return parseSS(size[:len(size)-1], _1k)
		}
	}
	return parseSS(size, 1)
}

func parseSS(size string, unit int) (int, error) {
	if i, err := strconv.Atoi(size); err == nil {
		return i * unit, nil
	}
	return 0, errors.New("Invalid thread stack size: -Xss")
}
