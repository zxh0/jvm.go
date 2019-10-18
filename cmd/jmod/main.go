package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io/ioutil"
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
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	zr, err := zip.NewReader(bytes.NewReader(data[4:]), int64(len(data)-4))
	if err != nil {
		panic(err)
	}

	for _, f := range zr.File {
		fmt.Println(f.Name)
	}
}
