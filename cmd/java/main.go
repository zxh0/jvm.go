package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"strings"

	"github.com/zxh0/jvm.go/classpath"
	"github.com/zxh0/jvm.go/interpreter"
	_ "github.com/zxh0/jvm.go/native"
	"github.com/zxh0/jvm.go/options"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

func main() {
	opts, args := options.Parse()
	if opts.HelpFlag || len(args) == 0 {
		printUsage()
	} else if opts.VersionFlag {
		printVersion()
	} else {
		startJVM(opts, args[0], args[1:])
	}
}

func printUsage() {
	fmt.Printf("usage: %s [-options] class [args...]\n", os.Args[0])
	options.PrintDefaults()
}

func printVersion() {
	fmt.Println("jvm.go 0.1.8.0")
}

func startJVM(opts options.Options, mainClass string, args []string) {
	if opts.XCPUProfile != "" {
		f, err := os.Create(opts.XCPUProfile)
		if err != nil {
			panic(err)
		}
		_ = pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	cp := classpath.Parse(opts)
	heap.InitBootLoader(cp, opts.VerboseClass)

	mainClass = strings.ReplaceAll(mainClass, ".", "/")
	mainThread := createMainThread(opts, mainClass, args)
	interpreter.Loop(mainThread)
	interpreter.KeepAlive()
}

func createMainThread(opts options.Options, className string, args []string) *rtda.Thread {
	mainThread := rtda.NewThread(nil, opts)
	bootMethod := heap.BootstrapMethod()
	bootArgs := []heap.Slot{heap.NewHackSlot(className), heap.NewHackSlot(args)}
	mainThread.InvokeMethodWithShim(bootMethod, bootArgs)
	return mainThread
}
