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
	classPath := cmd.Options().Classpath()
	classLoader := rtc.NewClassLoader(classPath)
	mainThread := createMainThread(classLoader, cmd.Class(), cmd.Args())
	interpreter.Loop(mainThread)
	keepalive.KeepAlive()
}

func initOptions(_options *cmdline.Options) {
	options.VerboseClass = _options.VerboseClass()
	options.ThreadStackSize = uint(_options.Xss())
}

func createMainThread(classLoader *rtc.ClassLoader, className string, args []string) *rtda.Thread {
	mainThread := rtda.NewThread(nil)
	bootMethod := rtc.NewBootstrapMethod([]byte{0xff, 0xb1}, classLoader)
	bootFrame := mainThread.NewFrame(bootMethod)
	mainThread.PushFrame(bootFrame)

	stack := bootFrame.OperandStack()
	stack.Push(args)
	stack.Push(className)
	stack.Push(classLoader)
	//stack.PushInt(0)

	return mainThread
}
