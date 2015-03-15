package jvm

import (
	"jvmgo/cmdline"
	"jvmgo/jvm/interpreter"
	"jvmgo/jvm/keepalive"
	"jvmgo/jvm/options"
	"jvmgo/jvm/rtda"
	rtc "jvmgo/jvm/rtda/class"
	_ "jvmgo/native"
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
	bootFrame := mainThread.NewFrame(bootMethod)
	mainThread.PushFrame(bootFrame)

	stack := bootFrame.OperandStack()
	stack.Push(args)
	stack.Push(className)

	return mainThread
}
