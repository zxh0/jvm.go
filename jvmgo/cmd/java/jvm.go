package main

import (
	"os"
	"runtime/pprof"

	"github.com/zxh0/jvm.go/jvmgo/classpath"
	"github.com/zxh0/jvm.go/jvmgo/interpreter"
	"github.com/zxh0/jvm.go/jvmgo/jutil"
	_ "github.com/zxh0/jvm.go/jvmgo/native"
	"github.com/zxh0/jvm.go/jvmgo/options"
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	"github.com/zxh0/jvm.go/jvmgo/rtda/heap"
)

func startJVM(cmd *Command) {
	Xcpuprofile := cmd.Options().Xcpuprofile
	if Xcpuprofile != "" {
		f, err := os.Create(Xcpuprofile)
		if err != nil {
			panic(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	options.InitOptions(cmd.Options().VerboseClass(), cmd.Options().Xss, cmd.Options().XuseJavaHome)

	cp := classpath.Parse(cmd.Options().Classpath())
	heap.InitBootLoader(cp)

	mainClassName := jutil.ReplaceAll(cmd.Class(), ".", "/")
	mainThread := createMainThread(mainClassName, cmd.Args())
	interpreter.Loop(mainThread)
	interpreter.KeepAlive()
}

func createMainThread(className string, args []string) *rtda.Thread {
	mainThread := rtda.NewThread(nil)
	bootMethod := heap.BootstrapMethod()
	bootArgs := []interface{}{className, args}
	mainThread.InvokeMethodWithShim(bootMethod, bootArgs)
	return mainThread
}
