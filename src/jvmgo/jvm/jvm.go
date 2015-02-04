package jvm

import (
    . "jvmgo/any"
    "jvmgo/cmdline"
    _ "jvmgo/native"
    "jvmgo/jvm/rtda"
    rtc "jvmgo/jvm/rtda/class"
)

func Startup(cmd *cmdline.Command) {
    classPath := cmd.Options().Classpath()
    classLoader := rtc.NewClassLoader(classPath)
    mainThread := createMainThread(classLoader, cmd.Class(), cmd.Args())
    loop(mainThread)
}

func createMainThread(classLoader Any, className string, args []string) (*rtda.Thread) {
    fakeMethod := rtc.NewStartupMethod([]byte{0xff, 0xb1}, classLoader)
    mainThread := rtda.NewThread(128, nil)
    mainFrame := mainThread.NewFrame(fakeMethod)
    mainThread.PushFrame(mainFrame)
    
    stack := mainFrame.OperandStack()
    stack.Push(args)
    stack.Push(className)
    stack.Push(classLoader)

    return mainThread
}
