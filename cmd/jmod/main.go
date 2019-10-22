package main

import (
	"fmt"
	"github.com/zxh0/jvm.go/vmutils"
	"os"

	"github.com/docopt/docopt-go"
)

const (
	version = "jmod.go 0.0.1"
	usage   = `jmod.

Usage:
  jmod list <file>
  jmod -h | --help
  jmod --version

Commands:
  list <file>   Prints the names of all the entries.

Options:
  -h --help     Print this help message
  --version     Print version information`
)

func main() {
	args := os.Args[1:]
	if opts, err := docopt.ParseArgs(usage, args, version); err != nil {
		fmt.Println(usage)
	} else if opts["list"].(bool) {
		list(opts["<file>"].(string))
	}
}

func list(filename string) {
	if r, err := vmutils.OpenJModReader(filename); err != nil {
		panic(err)
	} else {
		for _, f := range r.File {
			fmt.Println(f.Name)
		}
	}
}
