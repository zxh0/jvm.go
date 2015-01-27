package main

import (
    . "jvmgo/any"
    "jvmgo/cmdline"
    "jvmgo/native"
    "jvmgo/rtda"
    rtc "jvmgo/rtda/class"
)

// load native methods
func init() {
    var _ native.NativeMethod = nil
}

func startJVM(cmd *cmdline.Command) {
    classPath := cmd.Options().Classpath()
    classLoader := rtc.NewClassLoader(classPath)
    mainThread := createMainThread(classLoader, cmd.Class(), cmd.Args())
    loop(mainThread)
}

func createMainThread(classLoader Any, className string, args []string) (*rtda.Thread) {
    fakeFields := []Any{classLoader, className, args}
    fakeRef := rtc.NewObj(fakeFields)

    fakeMethod := rtc.NewStartupMethod([]byte{0xff, 0xb1}, classLoader)
    fakeFrame := rtda.NewFrame(fakeMethod)
    fakeFrame.OperandStack().PushRef(fakeRef)

    mainThread := rtda.NewThread(128)
    mainThread.PushFrame(fakeFrame)
    return mainThread
}
