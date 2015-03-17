package jvm

import (
	. "github.com/zxh0/jvm.go/jvmgo/any"
	"github.com/zxh0/jvm.go/jvmgo/cmdline"
	"github.com/zxh0/jvm.go/jvmgo/jvm/interpreter"
	"github.com/zxh0/jvm.go/jvmgo/jvm/keepalive"
	"github.com/zxh0/jvm.go/jvmgo/jvm/options"
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
	_ "github.com/zxh0/jvm.go/jvmgo/native"
)

func Startup(cmd *cmdline.Command) {
	initOptions(cmd.Options())
	rtc.InitBootLoader(cmd.Options().Classpath())

	mainThread := createMainThread(cmd.Class(), cmd.Args())
	interpreter.Loop(mainThread)
	keepalive.KeepAlive()
}

func initOptions(_options *cmdline.Options) {
	options.VerboseClass = _options.VerboseClass()
	options.ThreadStackSize = uint(_options.Xss())
}

func createMainThread(className string, args []string) *rtda.Thread {
	mainThread := rtda.NewThread(nil)
	bootMethod := rtc.BootstrapMethod()
	bootArgs := []Any{className, args}
	mainThread.InvokeMethodWithShim(bootMethod, bootArgs)
	return mainThread
}
