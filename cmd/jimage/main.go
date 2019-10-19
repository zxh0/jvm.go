package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/docopt/docopt-go"
	"github.com/zxh0/jvm.go/jimage"
)

const (
	version = "jimage.go 0.0.1"
	usage   = `jimage.

Usage:
  jimage info <file>
  jimage list [--verbose] <file>
  jimage -h | --help
  jimage --version

Options:
  -h --help     Print this help message
  --version     Print version information
  --verbose     Listing prints entry size and offset attributes`
)

func main() {
	args := os.Args[1:]
	if opts, err := docopt.ParseArgs(usage, args, version); err != nil {
		fmt.Println(usage)
	} else if opts["info"].(bool) {
		info(opts["<file>"].(string))
	} else if opts["list"].(bool) {
		list(
			opts["<file>"].(string),
			opts["--verbose"].(bool),
		)
	}
}

func info(filename string) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err) // TODO
	}

	header := jimage.ReadHeader(bytes)
	fmt.Printf(" Major Version:  %d\n", header.MajorVersion)
	fmt.Printf(" Minor Version:  %d\n", header.MinorVersion)
	fmt.Printf(" Flags:          %d\n", header.Flags)
	fmt.Printf(" Resource Count: %d\n", header.ResourceCount)
	fmt.Printf(" Table Length:   %d\n", header.TableLength)
	fmt.Printf(" Offsets Size:   %d\n", header.GetOffsetsSize())
	fmt.Printf(" Redirects Size: %d\n", header.GetRedirectSize())
	fmt.Printf(" Locations Size: %d\n", header.LocationsSize)
	fmt.Printf(" Strings Size:   %d\n", header.StringsSize)
	fmt.Printf(" Index Size:     %d\n", header.GetIndexSize())
}

func list(filename string, verbose bool) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err) // TODO
	}

	absPath, _ := filepath.Abs(filename)
	fmt.Printf("jimage: %s\n", absPath)
	image := jimage.ReadImage(bytes)
	listEntryNames(image, verbose)
}

func listEntryNames(image jimage.Image, verbose bool) {
	oldModule := ""
	for _, name := range image.GetEntryNames() {
		if !jimage.IsTreeInfoResource(name) {
			if module := getModuleName(name); module != oldModule {
				printModule(module, verbose)
				oldModule = module
			}

			printEntryName(image, name, verbose)
		}
	}
}

func printModule(module string, verbose bool) {
	fmt.Println("\nModule: " + module)
	if verbose {
		fmt.Println("Offset       Size       Compressed Entry")
	}
}

func printEntryName(image jimage.Image, name string, verbose bool) {
	if verbose {
		location := image.FindLocation(name)
		fmt.Printf("%12d %10d %10d %s\n",
			location.GetContentOffset(),
			location.GetUncompressedSize(),
			location.GetCompressedSize(),
			trimModule(name))
	} else {
		fmt.Println("    " + trimModule(name))
	}
}

func getModuleName(name string) string {
	if slashIdx := strings.IndexByte(name[1:], '/') + 1; slashIdx >= 0 {
		return name[1:slashIdx]
	}
	return "<unknown>"
}

func trimModule(name string) string {
	slashIdx := strings.IndexByte(name[1:], '/') + 1
	if slashIdx >= 0 && slashIdx+1 < len(name) {
		return name[slashIdx+1:]
	}
	return name
}
