package main

import (
    . "jvmgo/any"
    "jvmgo/cmdline"
    _ "jvmgo/native"
    "jvmgo/rtda"
    rtc "jvmgo/rtda/class"
)

type JVM struct {
    // todo
}

func (self JVM) startup(cmd *cmdline.Command) {
    classPath := cmd.Options().Classpath()
    classLoader := rtc.NewClassLoader(classPath)
    mainThread := createMainThread(classLoader, cmd.Class(), cmd.Args())
    loop(mainThread)
}

func createMainThread(classLoader Any, className string, args []string) (*rtda.Thread) {
    fakeFields := []Any{classLoader, className, args}
    fakeRef := rtc.NewObj(fakeFields)
    fakeMethod := rtc.NewStartupMethod([]byte{0xff, 0xb1}, classLoader)

    mainThread := rtda.NewThread(128, nil)
    mainFrame := mainThread.NewFrame(fakeMethod)
    mainFrame.OperandStack().PushRef(fakeRef)
    mainThread.PushFrame(mainFrame)

    return mainThread
}
