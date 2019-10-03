package main

import (
	"fmt"
	"io/ioutil"

	"github.com/docopt/docopt-go"
	"github.com/zxh0/jvm.go/jimage"
)

const usage = `jimage.

Usage:
  jimage info <file>
  jimage list [--verbose] <file>
  jimage -h | --help
  jimage --version

Options:
  -h --help     Print this help message
  --version     Print version information
  --verbose     Listing prints entry size and offset attributes`

func main() {
	if opts, err := docopt.ParseDoc(usage); err != nil {
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
		panic(err)
	}

	header := jimage.ReadHeader(jimage.NewImageReader(bytes))
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
	println("list ..." + filename)
	println(verbose)
}
