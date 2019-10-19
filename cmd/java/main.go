package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"strings"

	"github.com/zxh0/jvm.go/classpath"
	"github.com/zxh0/jvm.go/cpu"
	_ "github.com/zxh0/jvm.go/native"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
	"github.com/zxh0/jvm.go/vm"
)

func main() {
	opts, args := parseOptions()
	if opts.HelpFlag || len(args) == 0 {
		printUsage()
	} else if opts.VersionFlag {
		printVersion()
	} else {
		startJVM(opts, args[0], args[1:])
	}
}

func parseOptions() (vm.Options, []string) {
	options := vm.Options{}
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

	options.Init()
	return options, flag.Args()
}

func printUsage() {
	fmt.Printf("usage: %s [-options] class [args...]\n", os.Args[0])
	flag.PrintDefaults()
}

func printVersion() {
	fmt.Println("jvm.go 0.1.8.0")
}

func startJVM(opts vm.Options, mainClass string, args []string) {
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
	cpu.Loop(mainThread)
	cpu.KeepAlive()
}

func createMainThread(opts vm.Options, className string, args []string) *rtda.Thread {
	mainThread := rtda.NewThread(nil, opts)
	bootMethod := rtda.ShimBootstrapMethod
	bootArgs := []heap.Slot{heap.NewHackSlot(className), heap.NewHackSlot(args)}
	mainThread.InvokeMethodWithShim(bootMethod, bootArgs)
	return mainThread
}
