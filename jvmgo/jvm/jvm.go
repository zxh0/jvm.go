package jvm

import (
	. "github.com/zxh0/jvm.go/jvmgo/any"
	"github.com/zxh0/jvm.go/jvmgo/classpath"
	"github.com/zxh0/jvm.go/jvmgo/cmdline"
	"github.com/zxh0/jvm.go/jvmgo/jvm/interpreter"
	"github.com/zxh0/jvm.go/jvmgo/jvm/keepalive"
	"github.com/zxh0/jvm.go/jvmgo/jvm/options"
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
	_ "github.com/zxh0/jvm.go/jvmgo/native"
	"os"
	"runtime/pprof"
)

func Startup(cmd *cmdline.Command) {
	Xcpuprofile := cmd.Options().Xcpuprofile
	if Xcpuprofile != "" {
		f, err := os.Create(Xcpuprofile)
		if err != nil {
			panic(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	initOptions(cmd.Options())

	cp := classpath.Parse(cmd.Options().Classpath())
	rtc.InitBootLoader(cp)

	mainThread := createMainThread(cmd.Class(), cmd.Args())
	interpreter.Loop(mainThread)
	keepalive.KeepAlive()
}

func initOptions(cmdOptions *cmdline.Options) {
	options.VerboseClass = cmdOptions.VerboseClass()
	options.ThreadStackSize = uint(cmdOptions.Xss())
}

func createMainThread(className string, args []string) *rtda.Thread {
	mainThread := rtda.NewThread(nil)
	bootMethod := rtc.BootstrapMethod()
	bootArgs := []Any{className, args}
	mainThread.InvokeMethodWithShim(bootMethod, bootArgs)
	return mainThread
}
